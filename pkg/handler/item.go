package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gznrf/todo-app"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	listId, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	var input todo.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	}); err != nil {
		WriteError(w, 500, err)
		return
	}

}

func (h *Handler) getAllItems(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {

}
