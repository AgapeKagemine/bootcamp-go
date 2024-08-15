package handler

import (
	"net/http"
	"orchestrator/internal/helper/jwt"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	token, err := jwt.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
