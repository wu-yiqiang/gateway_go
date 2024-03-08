package routes

import (
	"gateway_go/common"
	"gateway_go/controllers"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoutes(router *gin.RouterGroup) {
	// 用户注册
	router.POST("/register", controllers.AdminController.AdminRegister)
	router.POST("/login", controllers.AdminController.AdminLogin)
	// 需要token验证的接口
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.POST("/changePassword", controllers.AdminController.AdminChangePassword)
		authRouter.GET("/admin_info", controllers.AdminController.AdminInfo)
		authRouter.POST("/avator", controllers.AdminController.AdminInfoAvator)
	}
}

func SetAdminGroupRoutes(router *gin.RouterGroup) {
	// token验证
	authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	{
		authRouter.GET("/logout", controllers.AdminController.AdminLogout)
	}
}
