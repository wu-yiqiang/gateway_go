package routes

import (
	"gateway_go/docs"
	"gateway_go/global"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetupRouter() *gin.Engine {
	docs.SwaggerInfo.Title = global.App.Config.Swagger.Title
	docs.SwaggerInfo.Description = global.App.Config.Swagger.Desc
	docs.SwaggerInfo.Host = global.App.Config.Swagger.Host + ":" + global.App.Config.App.Port
	docs.SwaggerInfo.BasePath = global.App.Config.Swagger.BasePath
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()
	// 静态资源
	router.Static("/assets", "./storage/images/")
	// router.Static("/flet/videos", "./storage/videos/")
	// 跨域中间件
	router.Use(middleware.CORS())
	// 日志
	//router.Use(middleware.CO)
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.RateLimitMiddleware(time.Second, 100, 100)) //初始100，每秒放出100
	// user模块
	userGroup := router.Group("/user")
	{
		SetUserGroupRoutes(userGroup)
	}
	// 即时通讯模块
	contactGroup := router.Group("")
	{
		SetContactGroupRoutes(contactGroup)
	}
	// 文件管理
	fileGroup := router.Group("/file")
	{
		SetFileGroupRoutes(fileGroup)
	}
	return router
}

