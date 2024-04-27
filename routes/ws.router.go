package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetWsGroupRoutes(router *gin.RouterGroup) {
	router.GET("/p2p/:room_identity/:user_identity", controllers.WsP2PConnection)
}
