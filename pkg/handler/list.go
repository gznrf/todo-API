package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gznrf/todo-app"
	"net/http"
	"strconv"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := Validate(input); err != nil {
		WriteError(w, 400, err)
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, map[string]interface{}{
		"id": id,
	}); err != nil {
		WriteError(w, 500, err)
		return
	}
}

type GetAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, GetAllListResponse{
		Data: lists,
	}); err != nil {
		WriteError(w, 500, err)
		return
	}
}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, list); err != nil {
		WriteError(w, 500, err)
		return
	}
}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	var input todo.UpdateListInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
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

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	if err := WriteJson(w, 200, statusResponse{
		Status: "ok",
	}); err != nil {
		WriteError(w, 500, err)
		return
	}
}
