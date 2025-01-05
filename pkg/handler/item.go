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

	if err := Validate(input); err != nil {
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

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, items); err != nil {
		WriteError(w, 500, err)
	}
}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	itemId, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, item); err != nil {
		WriteError(w, 500, err)
	}
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	itemId, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	var input todo.UpdateItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := h.services.TodoItem.Update(userId, itemId, input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := WriteJson(w, 200, statusResponse{
		Status: "ok",
	}); err != nil {
		WriteError(w, 500, err)
		return
	}
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	itemId, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, statusResponse{
		Status: "ok",
	}); err != nil {
		WriteError(w, 500, err)
	}
}
