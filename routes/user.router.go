package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 用户注册
	router.POST("/register", controllers.UserController.UserRegister)
	router.POST("/login", controllers.UserController.UserLogin)
	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/queryUserInfo", controllers.UserController.QueryUserInfo)
		authRouter.POST("/updatePassword", controllers.UserController.UserUpdatePassword)
		authRouter.POST("/avatar", controllers.UserController.UserInfoAvatar)
	}
}

