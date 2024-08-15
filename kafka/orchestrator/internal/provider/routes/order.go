package routes

import (
	"orchestrator/internal/handler/order"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup) {
	rg.POST("/order", order.ActivatePackage)
}
