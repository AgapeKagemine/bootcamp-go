package routes

import (
	"orchestrator/internal/handler"
	"orchestrator/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	route := gin.New()

	route.Use(middleware.LoggingMiddleware())

	api := route.Group("/api")
	api.GET("/token", handler.GenerateToken)

	api.Use(middleware.AuthMiddleware())

	RegisterOrderRoutes(api)

	return route
}
