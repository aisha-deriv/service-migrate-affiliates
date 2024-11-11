package model

import (
	"encoding/json"
	"net/http"
)

// Response is a generic response structure
type Response[T any] struct {
	Data       T      `json:"data,omitempty"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode"`
}

// WriteJSON writes the Response object as JSON to the ResponseWriter
func WriteJSON[T any](w http.ResponseWriter, response Response[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
