package routes

import (
	_ "fmt"
	_ "gateway_go/common/request"
	app2 "gateway_go/controllers/app"
	"gateway_go/middleware"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetUserGroupRoutes 定义 user 分组路由x`x“
func SetUserGroupRoutes(router *gin.RouterGroup) {
	// @Summary 用户列表
	// @Produce  json
	// @Failure 400 {object} tool.JSONRET "参数错误"
	// @Failure 20001 {object} tool.JSONRET "Token鉴权失败"
	// @Failure 20002 {object} tool.JSONRET "Token已超时"
	// @Failure 20004 {object} tool.JSONRET "Token错误"
	// @Failure 20005 {object} tool.JSONRET "Token参数不能为空"
	// @Success 0 {object} models.UserSwagger "查询成功"
	// @Router /api/v1/users [get]
	// user测试路由
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "user测试路由")
	})
	// 用户注册路由
	router.POST("/register", app2.Register)

	// 用户登陆
	router.POST("/login", app2.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app2.Info)
		authRouter.POST("/auth/logout", app2.Logout)
	}
}

//// 通过token解析用户信息
//func ParseToken(token string) (userInfo map[string]string, err error) {
//
//}
