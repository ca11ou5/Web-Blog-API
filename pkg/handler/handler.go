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

	router.HandleFunc("/articles", h.CreateArticle).Methods("POST")
	router.HandleFunc("/articles", h.GetArticleByID).Methods("GET")
	router.HandleFunc("/", h.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles", h.DeleteArticleByID).Methods("DELETE")

	return router
}
