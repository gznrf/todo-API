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

	if err := todo.ValidateUser(input); err != nil {
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

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {

}
