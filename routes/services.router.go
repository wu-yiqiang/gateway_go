package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetServicesGroupRoutes(router *gin.RouterGroup) {

	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.GET("/service_list", controllers.ServicesController.ServicesList)
		authRouter.GET("/service_delete", controllers.ServicesController.ServicesDelete)

		// router.POST("/service_add_tcp", controllers.ServicesController.ServicesAddTcp)
		//router.POST("/service_update_tcp", controllers.ServicesController.ServicesUpdateTcp)
		authRouter.POST("/service_add_grpc", controllers.ServicesController.ServicesAddGrpc)
		//router.POST("/service_update_grpc", controllers.ServicesController.ServicesUpdateGrpc)
	}
}
