package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//// 前端项目静态资源
	//router.StaticFile("/", "./static/dist/index.html")
	//router.Static("/assets", "./static/dist/assets")
	//router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	//// 其他静态资源
	//router.Static("/public", "./static")
	//router.Static("/storage", "./storage/app/public")

	// user路由 /user
	userGroup := router.Group("/user")
	SetUserGroupRoutes(userGroup)

	// demo路由 /demo
	DemoGroup := router.Group("/demo")
	SetDemoGroupRoutes(DemoGroup)

	return router
}
