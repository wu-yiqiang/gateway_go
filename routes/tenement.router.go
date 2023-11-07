package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetTenementGroupRoutes(router *gin.RouterGroup) {

	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.GET("/app_list", controllers.TenementController.TenementList)
		authRouter.GET("/app_delete", controllers.TenementController.TenementDelete)
	}
}
