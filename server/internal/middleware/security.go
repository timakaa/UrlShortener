package middleware

import (
	"net/http"
	"strconv"

	"github.com/timakaa/test-go/internal/config"
)

// SecurityMiddleware returns middleware with default settings
func SecurityMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return CustomSecurityMiddleware(config.SecurityConfig{
		FrameOptions:   "DENY",
		XSSProtection: "1; mode=block",
		ContentPolicy:  "default-src 'self'",
		HSTPMaxAge:    31536000,
	})
}

// CustomSecurityMiddleware returns middleware with custom settings
func CustomSecurityMiddleware(config config.SecurityConfig) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Set security headers
			w.Header().Set("X-Frame-Options", config.FrameOptions)
			w.Header().Set("X-XSS-Protection", config.XSSProtection)
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
			w.Header().Set("Content-Security-Policy", config.ContentPolicy)
			w.Header().Set("Strict-Transport-Security", "max-age="+strconv.Itoa(config.HSTPMaxAge)+"; includeSubDomains")
			w.Header().Set("X-DNS-Prefetch-Control", "off")
			
			next.ServeHTTP(w, r)
		}
	}
}