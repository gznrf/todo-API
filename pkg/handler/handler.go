package handler

import (
	"github.com/gorilla/mux"
	"github.com/gznrf/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.signIn).Methods("POST")
	}

	api := router.PathPrefix("/api").Subrouter()
	{
		api.Use(h.userIdentity)

		lists := api.PathPrefix("/lists").Subrouter()
		{
			lists.HandleFunc("/", h.createList).Methods("POST")
			lists.HandleFunc("/", h.getAllLists).Methods("GET")
			lists.HandleFunc("/{id}", h.getListById).Methods("GET")
			lists.HandleFunc("/{id}", h.updateList).Methods("PUT")
			lists.HandleFunc("/{id}", h.deleteList).Methods("DELETE")

			items := lists.PathPrefix("/{id}/items").Subrouter()
			{
				items.HandleFunc("/", h.createItem).Methods("POST")
				items.HandleFunc("/", h.getAllItems).Methods("GET")
			}
		}

		items := api.PathPrefix("/items").Subrouter()
		{
			items.HandleFunc("/{item_id}", h.getItemById).Methods("GET")
			items.HandleFunc("/{item_id}", h.updateItem).Methods("PUT")
			items.HandleFunc("/{item_id}", h.deleteItem).Methods("DELETE")
		}
	}

	return router
}
