package initialize

import (
	"context"
	_ "gateway_go/docs"
	"gateway_go/global"
	"gateway_go/routes"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//// 前端项目静态资源
	//router.StaticFile("/", "./static/dist/index.html")
	//router.Static("/assets", "./static/dist/assets")
	//router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	//// 其他静态资源
	//router.Static("/public", "./static")
	//router.Static("/storage", "./storage/app/public")

	// api 分组路由 /api
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	// user 分组路由 /user
	userGroup := router.Group("/user")
	routes.SetUserGroupRoutes(userGroup)

	// demo 分组路由 /demo
	DemoGroup := router.Group("/demo")
	routes.SetDemoGroupRoutes(DemoGroup)

	return router
}
