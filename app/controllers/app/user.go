package app

import (
	"github.com/gin-gonic/gin"
	"gateway_go/app/common/request"
	"gateway_go/app/common/response"
	"gateway_go/app/services"
)

// 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, _ := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "注册成功")
	}
}