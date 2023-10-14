package app

import (
	"fmt"
	"gateway_go/common/request"
	"gateway_go/common/response"
	"gateway_go/global"
	"gateway_go/models"
	"gateway_go/services"
	"gateway_go/utils"
	"gateway_go/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 用户注册
func Register(c *gin.Context) {
	var form validator.Register
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

// 修改密码
func ChangePassword(c *gin.Context) {
	var form validator.ChangePassword
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	oldPassword := form.Password
	newPassword := form.NewPassword
	if oldPassword == newPassword {
		response.ValidateFail(c, "新旧密码一致")
		return
	}
	// 旧密码正确性校验
	err, userInfo := services.UserService.GetUserInfo(strconv.FormatUint(uint64(form.UserId), 10))
	if err != nil {
		response.ValidateFail(c, "查询不到该用户")
	}
	isEquil := utils.BcryptMakeCheck([]byte(oldPassword), userInfo.Password)
	if !isEquil {
		response.ValidateFail(c, "旧密码输入错误")
		return
	}
	// 插入密码
	// 密码加密
	hashPwd := utils.BcryptMake([]byte(newPassword))
	fmt.Println("sad", hashPwd, form.UserId)
	error := global.App.DB.Where("id = ?", form.UserId).First(&models.User{}).Update("password", hashPwd).Error
	if error != nil {
		response.Success(c, "密码修改失败")
		return
	}
	response.Success(c, "密码修改成功")
	return
}
