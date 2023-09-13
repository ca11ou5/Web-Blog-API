package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "empty auth header"})
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid auth header"})
		c.Abort()
		return
	}

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("userID", userID)
}
