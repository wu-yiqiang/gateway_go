package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetFletGroupRoutes(router *gin.RouterGroup) {
	router.Group("").POST("/video", controllers.FletController.GetVideo)
	router.Group("").GET("/video/play", controllers.FletController.PlayVideo)
	router.Group("").GET("/video/banner", controllers.FletController.GetBanner)

	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/newsList", controllers.FletController.GetNewsLists)
		//authRouter.POST("/user", controllers.ServicesController.ServicesDelete)
		authRouter.POST("/weather", controllers.FletController.GetWeatherInfo)
		//authRouter.POST("/video", controllers.FletController.GetVideo)
	}
}
