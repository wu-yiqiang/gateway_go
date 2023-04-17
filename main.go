package main

import (
 _ "github.com/gin-gonic/gin"
  "gateway_go/bootstrap"
  "gateway_go/global"
  _"net/http"
 "flag"
)
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {


  // 获取运行环境
  mode := flag.String("env", "dev", "asda")
  flag.Parse()
  modeName := *mode
  // 初始化配置
  bootstrap.InitializeConfig(modeName)
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