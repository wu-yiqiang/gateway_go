package routes

import (
	"gateway_go/docs"
	"gateway_go/global"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	docs.SwaggerInfo.Title = global.App.Config.Swagger.Title
	docs.SwaggerInfo.Description = global.App.Config.Swagger.Desc
	docs.SwaggerInfo.Host = global.App.Config.Swagger.Host + ":" + global.App.Config.App.Port
	docs.SwaggerInfo.BasePath = global.App.Config.Swagger.BasePath
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()
	// 跨域中间件
	router.Use(middleware.CORS())
	// 日志
	//router.Use(middleware.CO)
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user路由
	userGroup := router.Group("/admin")
	{
		SetUserGroupRoutes(userGroup)
	}

	adminGroup := router.Group("/admin_login")
	{
		SetAdminGroupRoutes(adminGroup)
	}
	// services
	servicesGroup := router.Group("/service")
	{
		SetServicesGroupRoutes(servicesGroup)
	}

	// 租户管理
	tenementGroup := router.Group("/app")
	{
		SetTenementGroupRoutes(tenementGroup)
	}
	// 其他路由

	return router
}
