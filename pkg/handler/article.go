package handler

import (
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllArticles(c *gin.Context) {
	articles := h.service.Article.GetAllArticles()
	c.JSON(http.StatusOK, articles)
}

func (h *Handler) GetArticleByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Article.GetArticleByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) CreateArticle(c *gin.Context) {
	article := models.Article{}
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleID, err := h.service.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": articleID})
}

func (h *Handler) DeleteArticleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Article.DeleteArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
