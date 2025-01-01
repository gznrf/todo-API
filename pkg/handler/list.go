package handler

import (
	"encoding/json"
	"github.com/gznrf/todo-app"
	"net/http"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(w, r)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, 400, err)
		return
	}

	id, err = h.services.TodoList.Create(id, input)
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

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {

}
