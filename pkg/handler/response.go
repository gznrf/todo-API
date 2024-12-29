package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if err := WriteJson(w, status, map[string]string{"error": err.Error()}); err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}
}
