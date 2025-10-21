package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse is the standard, uniform JSON response.
// This struct is identical to the previous example.
type APIResponse[T any] struct {
	Success bool    `json:"success"`
	Data    T       `json:"data"`
	Error   *string `json:"error"`
}

// write is an internal helper to set headers and write JSON
func write[T any](w http.ResponseWriter, statusCode int, payload APIResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// Success sends a standard 2xx success response.
func Success[T any](w http.ResponseWriter, statusCode int, data T) {
	write(w, statusCode, APIResponse[T]{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

// Error sends a standard 4xx/5xx error response.
func Error(w http.ResponseWriter, statusCode int, err error) {
	write(w, statusCode, APIResponse[any]{
		Success: false,
		Data:    nil,
		Error:   func() *string { s := err.Error(); return &s }(),
	})
}
