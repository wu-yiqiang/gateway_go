package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 用户注册
	router.POST("/register", controllers.AdminController.AdminRegister)
	router.POST("/login", controllers.AdminController.AdminLogin)
	router.POST("/auth/changePassword", controllers.AdminController.AdminChangePassword)
	// token验证
	// authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	//{
	//	authRouter.POST("/auth/changePassword", controllers.AdminController.AdminChangePassword)
	//	authRouter.POST("/auth/logout", controllers.AdminController.AdminLogout)
	//}
}
