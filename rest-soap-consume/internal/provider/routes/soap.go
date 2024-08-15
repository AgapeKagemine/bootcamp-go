package routes

import (
	"consume-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func (r *Routes) SoapRoutes(rg *gin.RouterGroup) {
	soap := rg.Group("/soap")
	soap.GET("/", handler.ConsumeSOAP)
}
