package handler

import (
	"blog/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", h.CreateArticle).Methods("POST")
	router.HandleFunc("/{name}", h.GetArticle).Methods("GET")
	return router
}
