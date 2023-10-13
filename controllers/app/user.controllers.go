package app

import (
	request2 "gateway_go/common/request"
	"gateway_go/common/response"
	"gateway_go/services"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {
	var form request2.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request2.GetErrorMsg(form, err))
		return
	}
	if err, _ := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, "新用户注册成功")
	}
}
