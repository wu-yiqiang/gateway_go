package routes

import (
 _ "fmt"
  "github.com/gin-gonic/gin"
  "gateway_go/app/common/request"
  "net/http"
  "time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
  // API测试路由
  router.GET("/test", func(c *gin.Context) {
    time.Sleep(5*time.Second)
    c.String(http.StatusOK, "API测试路由")
  })
}
