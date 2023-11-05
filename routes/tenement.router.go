package routes

import (
	"gateway_go/controllers"
	"github.com/gin-gonic/gin"
)

func SetTenementGroupRoutes(router *gin.RouterGroup) {
	router.GET("/app_list", controllers.TenementController.TenementList)
}
