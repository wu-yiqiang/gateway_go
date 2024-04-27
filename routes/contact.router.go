package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetContactGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/add", controllers.ContactController.AddFriend)
	}
}
