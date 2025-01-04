package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

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

func Validate(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
