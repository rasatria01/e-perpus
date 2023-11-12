package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func JsonError(w http.ResponseWriter, statuscode int, message string) {
	JsonResponse(w, statuscode, map[string]string{"error": message})
}
