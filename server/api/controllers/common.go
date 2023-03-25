// Defines two functions for responding to HTTP requests with
// JSON data or errors.

package controllers

import (
	"encoding/json"
	"net/http"
)

// Responds with JSON data
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	// This function first serializes the payload into a JSON
	// format using the json.Marshal function
	response, err := json.Marshal(payload)
	if err != nil {
		// The function responds with an HTTP 500 Internal Server
		// Error status
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	//This tells the client that the response body will contain JSON data.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// Responds with JSON error
func respondError(w http.ResponseWriter, code int, message string) {
	// Map has key as "error" and value as message
	respondJSON(w, code, map[string]string{"error": message})
}
