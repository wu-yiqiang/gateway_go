package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetServicesGroupRoutes(router *gin.RouterGroup) {
	router.GET("/service_list", controllers.ServicesController.ServicesList)
	router.GET("/service_delete", controllers.ServicesController.ServicesDelete)
}
