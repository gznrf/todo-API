package handler

import (
	"encoding/json"
	"github.com/gznrf/todo-app"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input todo.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := Validate(input); err != nil {
		WriteError(w, 400, err)
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, map[string]interface{}{
		"id": id,
	}); err != nil {
		WriteError(w, 500, err)
	}
}

type SignInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input SignInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := Validate(input); err != nil {
		WriteError(w, 400, err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, map[string]interface{}{
		"token": token,
	}); err != nil {
		WriteError(w, 500, err)
	}
}
