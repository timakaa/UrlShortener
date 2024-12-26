package utils

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/timakaa/test-go/internal/models"

	"gorm.io/gorm"
)

func generateShortURL() string {
	// Characters allowed in the short URL
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	// Create a byte slice to store the result
	result := make([]byte, length)
	
	// Generate random bytes
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	
	return string(result)
}

// GenerateUniqueShortURL generates a unique short URL and checks against the database
func GenerateUniqueShortURL(db *gorm.DB) (string, error) {
	maxAttempts := 10
	
	for i := 0; i < maxAttempts; i++ {
		shortURL := generateShortURL()
		
		// Check if exists using GORM
		var count int64
		if err := db.Model(&models.Url{}).Where("short_url = ?", shortURL).Count(&count).Error; err != nil {
			log.Printf("Error checking short URL existence: %v", err)
			continue
		}
		
		if count == 0 {
			return shortURL, nil
		}
	}
	
	return "", fmt.Errorf("failed to generate unique short URL after %d attempts", maxAttempts)
}
