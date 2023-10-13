package routes

import (
	"gateway_go/app/controllers/app"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 Demo 分组路由
func SetDemoGroupRoutes(router *gin.RouterGroup) {
	// demo测试路由
	router.GET("/ws", app.WebSocketHandler)
	router.GET("/upload", app.UploadHandler)
	router.GET("/download", app.DownloadHandler)

	router.GET("/concurrency", app.HandleConcurrencyRequest)
}
