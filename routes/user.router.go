package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 用户注册
	router.POST("/register", controllers.AdminController.AdminRegister)
	router.POST("/login", controllers.AdminController.AdminLogin)
	router.POST("/changePassword", controllers.AdminController.AdminChangePassword)
	router.GET("/admin_info", controllers.AdminController.AdminInfo)
	// token验证
	// authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	//{
	//	authRouter.POST("/auth/changePassword", controllers.AdminController.AdminChangePassword)
	//	authRouter.POST("/auth/logout", controllers.AdminController.AdminLogout)
	//}
}

func SetAdminGroupRoutes(router *gin.RouterGroup) {
	router.GET("/logout", controllers.AdminController.AdminLogout)
	// token验证
	// authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	//{
	//	authRouter.POST("/auth/changePassword", controllers.AdminController.AdminChangePassword)
	//	authRouter.POST("/auth/logout", controllers.AdminController.AdminLogout)
	//}
}
