package handler

import (
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	article := router.Group("/")
	{
		article.GET("/", h.GetAllArticles)
		article.GET("/articles/:id", h.GetArticleByID)
		middleware := article.Group("/articles", h.userIdentity)
		{
			middleware.POST("", h.CreateArticle)
			middleware.DELETE(":id", h.DeleteArticleByID)
		}
	}

	auth := router.Group("/")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//router.HandleFunc("/articles", h.CreateArticle).Methods("POST")
	//router.HandleFunc("/articles", h.GetArticleByID).Methods("GET")
	//router.HandleFunc("/", h.GetAllArticles).Methods("GET")
	//router.HandleFunc("/articles", h.DeleteArticleByID).Methods("DELETE")
	//
	//router.HandleFunc("/sign-up", h.signUp).Methods("POST")
	//router.HandleFunc("/sign-in", h.signIn).Methods("POST")

	return router
}
