package routes

import (
	"consume-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func (r *Routes) RestRoutes(rg *gin.RouterGroup) {
	rest := rg.Group("/rest")
	rest.GET("/", handler.ConsumeREST)
}
