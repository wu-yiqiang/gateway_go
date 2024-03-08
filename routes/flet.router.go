package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetFletGroupRoutes(router *gin.RouterGroup) {

	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/newsList", controllers.FletController.GetNewsLists)
		//authRouter.POST("/user", controllers.ServicesController.ServicesDelete)
		authRouter.POST("/weather", controllers.FletController.GetWeatherInfo)
	}
}
