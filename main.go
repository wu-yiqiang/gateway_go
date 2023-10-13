package main

import (
	"gateway_go/global"
	"gateway_go/initialize"
	_ "github.com/gin-gonic/gin"
	_ "net/http"
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
	initialize.InitializeConfig()

	// 初始化校验器
	initialize.InitializeValidator()

	// 初始化日志
	global.App.Log = initialize.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = initialize.InitializeDB()
	// 初始化Redis
	global.App.Redis = initialize.InitializeRedis()
	// 初始化定时任务
	initialize.InitCron()
	// 初始化文件系统
	// initialize.InitializeStorage()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 启动服务器
	initialize.RunServer()
}
