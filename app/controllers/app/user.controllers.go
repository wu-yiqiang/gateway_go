package app

import (
	"gateway_go/app/common/request"
	"gateway_go/app/common/response"
	"gateway_go/app/services"
	"github.com/gin-gonic/gin"
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
		response.Success(c, "新用户注册成功")
	}
}
