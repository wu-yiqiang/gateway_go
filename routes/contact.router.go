package routes

import (
	_ "gateway_go/common"
	_ "gateway_go/controllers"
	v1 "gateway_go/internal/api/v1"
	"gateway_go/internal/route"
	_ "gateway_go/middleware"
	"github.com/gin-gonic/gin"
)

func SetContactGroupRoutes(router *gin.RouterGroup) {
	// router.GET("/socket.io", RunSocekt)
	//authRouter := router.Group("").Use(middleware.JWTAuth(common.AppGuardName))
	//{
	//	authRouter.POST("/queryFriends", controllers.ContactController.QueryUserFriends)
	//	authRouter.POST("/add", controllers.ContactController.AddFriend)
	//	// authRouter.GET("/socket.io", RunSocekt)
	//}
	socket := route.RunSocekt
	{
		router.GET("/user", v1.GetUserList)
		router.GET("/user/:uuid", v1.GetUserDetails)
		router.GET("/user/name", v1.GetUserOrGroupByName)
		// 地址冲突先注释
		// router.POST("/user/register", v1.Register)
		// router.POST("/user/login", v1.Login)
		router.PUT("/user", v1.ModifyUserInfo)
		router.POST("/friend", v1.AddFriend)
		router.GET("/message", v1.GetMessage)
		router.GET("/file/:fileName", v1.GetFile)
		router.POST("/file", v1.SaveFile)
		router.GET("/group/:uuid", v1.GetGroup)
		router.POST("/group/:uuid", v1.SaveGroup)
		router.POST("/group/join/:userUuid/:groupUuid", v1.JoinGroup)
		router.GET("/group/user/:uuid", v1.GetGroupUsers)
		router.GET("/socket.io", socket)
	}
}
