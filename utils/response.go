package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	ResponseJSON(w, status, map[string]interface{}{"error": message})
}
