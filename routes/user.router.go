package routes

import (
	"gateway_go/controllers"
	"gateway_go/middleware"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 用户注册
	router.POST("/register", controllers.AdminRegister)
	router.POST("/login", controllers.AdminLogin)
	// 需要验证
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/changepassword", controllers.AdminChangePassword)
		authRouter.POST("/auth/logout", controllers.AdminLogout)
	}
}
