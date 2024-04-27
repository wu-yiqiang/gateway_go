package routes

import (
	"gateway_go/docs"
	"gateway_go/global"
	"gateway_go/middleware"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
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
	router.Use(RateLimitMiddleware(time.Second, 100, 100)) //初始100，每秒放出100
	// user路由
	userGroup := router.Group("/admin")
	{
		SetUserGroupRoutes(userGroup)
	}

	adminGroup := router.Group("/admin_login")
	{
		SetAdminGroupRoutes(adminGroup)
	}
	// 即时通讯模块
	contactGroup := router.Group("/contact")
	{
		SetContactGroupRoutes(contactGroup)
	}
	// flutter接口
	fletGroup := router.Group("/flet")
	{
		SetFletGroupRoutes(fletGroup)
	}
	// 文件管理
	fileGroup := router.Group("/file")
	{
		SetFileGroupRoutes(fileGroup)
	}
	return router
}

// 接口限流
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
