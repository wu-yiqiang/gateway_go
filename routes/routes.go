package routes

import (
	"gateway_go/docs"
	"gateway_go/global"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
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

	// 其他路由

	return router
}
