package controllers

import (
	"context"
	"fmt"
	"gateway_go/dto"
	"gateway_go/global"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/services"
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
// @ID /admin/register
// @Accept  json
// @Produce  json
// @Param polygon body dto.RegisterInput true "body"
// @Success 200 {object} response.Response{} "success"
// @Router /admin/register [post]
func (admin *adminController) AdminRegister(c *gin.Context) {
	var form dto.RegisterInput
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
// @ID /admin/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.LoginInput true "body"
// @Success 200 {object} response.Response{data=dto.LoginOutput} "success"
// @Router /admin/login [post]
func (admin *adminController) AdminLogin(c *gin.Context) {
	var form dto.RegisterInput
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
		//if err != nil {
		//	response.BusinessFail(c, err.Error())
		//	return
		//}
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
// @ID /admin/changePassword
// @Accept  json
// @Produce  json
// @Param polygon body dto.ChangePasswordInput true "body"
// @Success 200 {object} response.Response{} "success"
// @Router /admin/changePassword [post]
func (admin *adminController) AdminChangePassword(c *gin.Context) {
	var form dto.ChangePasswordInput
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

// ListPage godoc
// @Summary 管理员信息获取
// @Description 管理员信息获取
// @Tags 用户管理
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Success 200 {object} response.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (admin *adminController) AdminInfo(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.BusinessFail(c, "token username参数不能为空")
		return
	}
	// 解密token
	err, data := services.JwtService.DecryptToken(token)
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}

	tokenStr, err := global.App.Redis.Get(context.Background(), data).Result()
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}
	out := &dto.AdminInfoOutput{
		Name:      data,
		Id:        1,
		Avatar:    "http://e.hiphotos.baidu.com/image/pic/item/a1ec08fa513d2697e542494057fbb2fb4316d81e.jpg",
		LoginTime: "2023-10-26",
		Roles:     []string{"admin"},
	}
	fmt.Println(tokenStr)
	response.Success(c, out)
}

// ListPage godoc
// @Summary 管理员注销
// @Description 管理员注销
// @Tags 用户管理
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{} "success"
// @Router /admin_login/logout [get]
func (admin *adminController) AdminLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		response.TokenFail(c)
		return
	}
	// 解密token
	err, data := services.JwtService.DecryptToken(token)
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}
	error := global.App.Redis.Del(context.Background(), data).Err()
	if error != nil {
		response.BusinessFail(c, "用户注销失败")
		return
	}
	response.Success(c, "用户注销成功")
}
