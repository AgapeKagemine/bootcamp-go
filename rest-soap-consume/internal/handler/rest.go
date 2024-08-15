package handler

import (
	"consume-api/internal/domain"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumeREST(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	var post []domain.Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, post)
}
