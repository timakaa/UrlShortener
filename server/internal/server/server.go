package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	globalMiddleware []func(http.HandlerFunc) http.HandlerFunc
}

func New() *Server {
	return &Server{
		globalMiddleware: make([]func(http.HandlerFunc) http.HandlerFunc, 0),
	}
}

// Use adds global middleware
func (s *Server) Use(middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	s.globalMiddleware = append(s.globalMiddleware, middleware...)
}

// Chain combines multiple middleware into one
func Chain(middlewares ...func(http.HandlerFunc) http.HandlerFunc) func(http.HandlerFunc) http.HandlerFunc {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(middlewares) - 1; i >= 0; i-- {
				last = middlewares[i](last)
			}
			last(w, r)
		}
	}
}

// Handle registers a new route with middleware
func (s *Server) Handle(pattern string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	// Combine global and route-specific middleware
	allMiddleware := append(s.globalMiddleware, middleware...)
	
	// Use the Chain helper to combine all middleware
	http.HandleFunc(pattern, Chain(allMiddleware...)(handler))
}

// GET registers a GET route
func (s *Server) GET(pattern string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	s.Handle(pattern, methodMiddleware("GET", handler), middleware...)
}

// POST registers a POST route
func (s *Server) POST(pattern string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	s.Handle(pattern, methodMiddleware("POST", handler), middleware...)
}

// DELETE registers a DELETE route
func (s *Server) DELETE(pattern string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	s.Handle(pattern, methodMiddleware("DELETE", handler), middleware...)
}

// PUT registers a PUT route
func (s *Server) PUT(pattern string, handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	s.Handle(pattern, methodMiddleware("PUT", handler), middleware...)
}

func methodMiddleware(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

// isPortAvailable checks if the port is available for use
func isPortAvailable(port string) bool {
	address := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	listener.Close()
	return true
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() error {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Add validation for port
	if _, err := strconv.Atoi(port); err != nil {
		log.Printf("Invalid port number: %s\n", port)
		port = "8080"
	}
	
	// Check if port is available
	if !isPortAvailable(port) {
		log.Fatalf("port %s is already in use", port)
		return errors.New("port is already in use")
	}
	
	address := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", address)
	
	server := &http.Server{
		Addr:    address,
		Handler: nil, // Use default ServeMux
	}
	
	return server.ListenAndServe()
}
