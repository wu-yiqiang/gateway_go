package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetFileGroupRoutes(router *gin.RouterGroup) {
	router.POST("/upload", controllers.FileController.Upload)
	router.POST("/mergechunks", controllers.FileController.MergeChunks)
	router.POST("/uploadFile", controllers.FileController.UploadFile)
	//// 需要token验证的接口
	//authRouter := router.Group("/file").Use(middleware.JWTAuth(common.AppGuardName))
	//{
	//	authRouter.POST("/upload", controllers.FileController.Upload)
	//	//authRouter.POST("/user", controllers.ServicesController.ServicesDelete)
	//	// authRouter.POST("/download", controllers.FileController.Download)
	//}
}
