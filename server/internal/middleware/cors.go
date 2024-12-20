package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/timakaa/test-go/internal/config"
)

// CorsMiddleware returns middleware with default settings
func CorsMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	var allowedOrigins []string = []string{"http://localhost:3000"}
	var allowedMethods []string = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	var allowedHeaders []string = []string{"Accept", "Content-Type", "Content-Length", "Authorization"}

	return CustomCorsMiddleware(config.CorsConfig{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		AllowCredentials: true,
		MaxAge:           300,
	})
}

// CustomCorsMiddleware returns middleware with custom settings
func CustomCorsMiddleware(config config.CorsConfig) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Set origin
			if len(config.AllowedOrigins) > 0 {
				if config.AllowedOrigins[0] == "*" {
					w.Header().Set("Access-Control-Allow-Origin", "*")
				} else {
					origin := r.Header.Get("Origin")
					for _, allowedOrigin := range config.AllowedOrigins {
						if origin == allowedOrigin {
							w.Header().Set("Access-Control-Allow-Origin", origin)
							break
						}
					}
				}
			}
			
			// Set other CORS headers
			if len(config.AllowedMethods) > 0 {
				w.Header().Set("Access-Control-Allow-Methods", 
					strings.Join(config.AllowedMethods, ", "))
			}
			
			if len(config.AllowedHeaders) > 0 {
				w.Header().Set("Access-Control-Allow-Headers", 
					strings.Join(config.AllowedHeaders, ", "))
			}
			
			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			
			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", 
					strconv.Itoa(config.MaxAge))
			}
			
			// Handle preflight
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		}
	}
}