package main

import (
 _ "github.com/gin-gonic/gin"
  "gateway_go/bootstrap"
  "gateway_go/global"
  _"net/http"
 "flag"
)

func main() {
  // 获取运行环境
  mode := flag.String("env", "dev", "asda")
  flag.Parse()
  modeName := *mode
  // 初始化配置
  bootstrap.InitializeConfig(modeName)

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