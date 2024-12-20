package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timakaa/test-go/internal/db"
	"github.com/timakaa/test-go/internal/models"
)

// CreateUrlHandler creates new short URL
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	var url models.Url
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

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
	shortURL := r.URL.Query().Get("code") // Получаем короткий код из query параметров
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

	// Увеличиваем счетчик посещений
	db.GetDB().Model(&url).Update("visits", url.Visits+1)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}