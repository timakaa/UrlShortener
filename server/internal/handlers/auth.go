package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/timakaa/test-go/internal/db"
	"github.com/timakaa/test-go/internal/email"
	"github.com/timakaa/test-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body
	var requestBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	email := requestBody.Email
	password := requestBody.Password

	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var user models.User
	result := db.GetDB().Where("email = ?", email).First(&user)
	if result.Error != nil || !user.Verified {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400, // 24 hours
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
	})
}

func GenerateVerificationCode() string {
	rand.NewSource(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}

type TempRegistrationData struct {
	Username string
	Email    string
	Password string
}

func RegisterInitHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var requestBody models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(requestBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", e.Field()))
			case "email":
				errorMessages = append(errorMessages, "Invalid email format")
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must not exceed %s characters", e.Field(), e.Param()))
			}
		}
		
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errorMessages,
		})
		return
	}

	// Check for existing username
	var existingUser models.User
	if err := db.GetDB().Where("username = ?", requestBody.Username).First(&existingUser).Error; err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	// Check for existing email
	if err := db.GetDB().Where("email = ?", requestBody.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Check for existing code in Redis
	exists, err := db.GetRedis().Exists(ctx, "code:"+requestBody.Email).Result()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if exists == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Verification code already sent. Please check your email",
		})
		return
	}

	code := GenerateVerificationCode()
	tempData := TempRegistrationData{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	// Serialize registration data
	jsonData, err := json.Marshal(tempData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Store both code and registration data
	err = db.GetRedis().Set(ctx, 
		"code:"+requestBody.Email, 
		code, 
		15*time.Minute,
	).Err()
	if err != nil {
		http.Error(w, "Failed to store verification code", http.StatusInternalServerError)
		return
	}

	// Store registration data with the same TTL
	err = db.GetRedis().Set(ctx, 
		"reg:"+requestBody.Email, 
		jsonData, 
		15*time.Minute,
	).Err()
	if err != nil {
		http.Error(w, "Failed to store registration data", http.StatusInternalServerError)
		return
	}

	go email.SendVerificationEmail(requestBody.Email, code)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Verification code sent to your email",
	})
}

func VerifyAndRegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var requestBody struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Get and verify code from Redis
	storedCode, err := db.GetRedis().Get(ctx, "code:"+requestBody.Email).Result()
	if err == redis.Nil {
		http.Error(w, "Verification code expired or invalid", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Failed to verify code", http.StatusInternalServerError)
		return
	}

	if storedCode != requestBody.Code {
		http.Error(w, "Invalid verification code", http.StatusBadRequest)
		return
	}

	// Get data from Redis
	jsonData, err := db.GetRedis().Get(ctx, "reg:"+requestBody.Email).Bytes()
	if err == redis.Nil {
		http.Error(w, "Registration data expired", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Failed to get registration data", http.StatusInternalServerError)
		return
	}

	var tempData TempRegistrationData
	if err := json.Unmarshal(jsonData, &tempData); err != nil {
		http.Error(w, "Invalid registration data", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tempData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: tempData.Username,
		Email:    tempData.Email,
		Password: string(hashedPassword),
		Verified: true,
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400, // 24 hours
	})

	// Clean up both verification code and registration data
	db.GetRedis().Del(ctx, 
		"code:"+requestBody.Email,
		"reg:"+requestBody.Email,
	)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":   tokenString,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
