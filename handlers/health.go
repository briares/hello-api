// Package handlers Handlers for the app.
package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck translates a word given the language.
// This function contains a comment.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := map[string]string{"status": "up"} //
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
