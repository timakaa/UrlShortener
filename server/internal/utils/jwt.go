package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims structure for JWT token
type Claims struct {
	UserID uint `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.RegisteredClaims
}

// Secret key for JWT signing (in production should be in env variables)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// ParseToken validates JWT token and returns claims
func ParseToken(tokenString string) (*Claims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Convert claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
