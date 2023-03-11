package main

import (
  "github.com/gin-gonic/gin"
  "gateway_go/bootstrap"
  "gateway_go/global"
  "net/http"
)

func main() {
  // 初始化配置
  bootstrap.InitializeConfig()

  global.App.Log = bootstrap.InitializeLog()
  global.App.Log.Info("log init success!")

  r := gin.Default()

  // 测试路由
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  // 启动服务器
  r.Run(":" + global.App.Config.App.Port)
}