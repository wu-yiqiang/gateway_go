package routes

import (
	_ "fmt"
	_ "gateway_go/app/common/request"
	"gateway_go/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetUserGroupRoutes 定义 user 分组路由
func SetUserGroupRoutes(router *gin.RouterGroup) {
	// user测试路由
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "user测试路由")
	})
	
	// 用户注册路由
	router.POST("/register", app.Register)
}
