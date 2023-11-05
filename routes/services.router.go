package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetServicesGroupRoutes(router *gin.RouterGroup) {
	router.GET("/service_list", controllers.ServicesController.ServicesList)
	router.GET("/service_delete", controllers.ServicesController.ServicesDelete)

	// router.POST("/service_add_tcp", controllers.ServicesController.ServicesAddTcp)
	//router.POST("/service_update_tcp", controllers.ServicesController.ServicesUpdateTcp)
	//router.POST("/service_add_grpc", controllers.ServicesController.ServicesAddGrpc)
	//router.POST("/service_update_grpc", controllers.ServicesController.ServicesUpdateGrpc)
}
