package main

import (
 _ "github.com/gin-gonic/gin"
  "gateway_go/bootstrap"
  "gateway_go/global"
  _"net/http"
)

// @title go-api 框架
// @version 1.0
// @description gin-web框架
// @termsofservice https://github.com/18211167516/Go-Gin-Api
// @contact.name atlas
// @contact.email wu_yiqiang@outlook.com
// @host 127.0.0.1:9527
func main() {
  // 初始化配置
  bootstrap.InitializeConfig();

  // 初始化校验器
  bootstrap.InitializeValidator()

  // 初始化日志
  global.App.Log = bootstrap.InitializeLog()
  global.App.Log.Info("log init success!")

  // 初始化数据库
  global.App.DB = bootstrap.InitializeDB()


  // 程序关闭前，释放数据库连接
  defer func() {
    if global.App.DB != nil {
      db, _ := global.App.DB.DB()
      db.Close()
    }
  }()

  // 启动服务器
  bootstrap.RunServer()
}