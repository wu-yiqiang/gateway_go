package controllers

import (
	"context"
	"fmt"
	"gateway_go/global"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/services"
	"gateway_go/validator"
	"github.com/gin-gonic/gin"
	"time"
)

type adminController struct {
}

var AdminController = new(adminController)

// ListPage godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户管理
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param polygon body dto.RegisterInput true "body"
// @Success 200 {object} response.Response{} "success"
// @Router /user/register [post]
func (admin *adminController) AdminRegister(c *gin.Context) {
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

// ListPage godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户管理
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.LoginInput true "body"
// @Success 200 {object} response.Response{content=dto.LoginOutput} "success"
// @Router /user/login [post]
func (admin *adminController) AdminLogin(c *gin.Context) {
	var form validator.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
		return
	} else {
		// 查询redis中token
		joinUnixStr, err := global.App.Redis.Get(context.Background(), user.Username).Result()
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		if joinUnixStr != "" {
			response.Success(c, &services.TokenOutPut{Token: joinUnixStr})
			return
		}
		// 生成token
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		// 将token存储到redis中
		err = global.App.Redis.SetNX(context.Background(), user.Username, tokenData.Token, time.Duration(global.App.Config.Jwt.JwtTtl)*time.Second).Err()
		fmt.Println("sad", err)
		if err != nil {
			response.Success(c, "token存储失败")
		}
		response.Success(c, tokenData)

	}

}

// ListPage godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 用户管理
// @ID /user/auth/changePassword
// @Accept  json
// @Produce  json
// @Param polygon body dto.ChangePasswordInput true "body"
// @Success 200 {object} response.Response{} "success"
// @Router /user/auth/changePassword [post]
func (admin *adminController) AdminChangePassword(c *gin.Context) {
	var form validator.ChangePassword
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if form.Password == form.NewPassword {
		response.ValidateFail(c, "新密码不能和旧密码相同")
		return
	}

	// 查询旧密码是否相同
	if err, _ := services.UserService.Changepassword(form); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	// 修改密码
	err := services.UserService.ModifyPassword(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, "修改密码成功")
	return
}

func (admin *adminController) AdminLogout(c *gin.Context) {

}
