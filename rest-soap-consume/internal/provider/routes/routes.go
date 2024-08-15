package routes

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Server *gin.Engine
}

func NewRoutes() *Routes {
	routes := &Routes{
		Server: gin.New(),
	}

	api := routes.Server.Group("/api")
	consume := api.Group("/consume")
	routes.RestRoutes(consume)
	routes.SoapRoutes(consume)
	return routes
}
