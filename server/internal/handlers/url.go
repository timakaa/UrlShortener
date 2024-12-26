package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timakaa/test-go/internal/db"
	"github.com/timakaa/test-go/internal/models"
	"github.com/timakaa/test-go/internal/utils"
)

// CreateUrlHandler creates new short URL
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not authenticated",
		})
		return
	}

	var url models.Url
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := utils.GenerateUniqueShortURL(db.GetDB())
	if err != nil {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}
	url.ShortUrl = shortURL
	url.UserID = userID

	// Create record in database
	result := db.GetDB().Create(&url)
	if result.Error != nil {
		http.Error(w, "Failed to create URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

// GetUrlHandler retrieves URL by short code
func GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Query().Get("code")
	if shortURL == "" {
		http.Error(w, "Short URL code is required", http.StatusBadRequest)
		return
	}

	var url models.Url
	result := db.GetDB().Where("short_url = ?", shortURL).First(&url)
	if result.Error != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Increment visit counter
	db.GetDB().Model(&url).Update("visits", url.Visits+1)

	// Return JSON with original URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"originalUrl": url.OriginalUrl,
	})
}

// ErrorResponse represents error message structure
type ErrorResponse struct {
	Error string `json:"error"`
}

func GetUrlsHandler(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context
	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not authenticated",
		})
		return
	}

	var urls []models.Url
	if result := db.GetDB().Where("user_id = ?", userID).Find(&urls); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to fetch URLs",
		})
		return
	}

	json.NewEncoder(w).Encode(urls)
}

