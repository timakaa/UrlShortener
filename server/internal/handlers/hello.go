package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/timakaa/test-go/pkg/response"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Request")
	
	// Create response
	response := response.Response{
		Message: "Hello from API!",
		Status:  true,
	}
	
	// Convert to JSON and send
	json.NewEncoder(w).Encode(response)
}