package controllers

import (
	"context"
	"fmt"
	"gateway_go/common"
	"gateway_go/dto"
	"gateway_go/global"
	"gateway_go/request"
	"gateway_go/response"
	"gateway_go/services"
	"gateway_go/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
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
// @Param polygon body dto.RegisterInput true "body"
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
			response.Success(c, &services.TokenOutPut{Token: common.TokenType + " " + joinUnixStr})
			return
		}
		// 生成token
		fmt.Println("user login", user)
		tokenData, err, _ := services.JwtService.CreateToken(common.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		// 将token存储到redis中
		err = global.App.Redis.SetNX(context.Background(), user.Username, tokenData.Token, time.Duration(global.App.Config.Jwt.JwtTtl)*time.Second).Err()
		if err != nil {
			response.BusinessFail(c, "token存储失败")
			return
		}
		tokenData.Token = common.TokenType + " " + tokenData.Token
		response.Success(c, tokenData)
		return
	}

}

// ListPage godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 用户管理
// @ID /admin/changePassword
// @Accept  json
// @Produce  json
// @Security Auth
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
// @Security Auth
// @Success 200 {object} response.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [post]
func (admin *adminController) AdminInfo(c *gin.Context) {
	userId, idIsExist := c.Get("userId")
	username, nameIsExist := c.Get("userName")
	if idIsExist == false || nameIsExist == false {
		response.BusinessFail(c, "用户不存在")
		return
	}

	tokenStr, err := global.App.Redis.Get(context.Background(), username.(string)).Result()
	fmt.Println("token 不能删除用于打印token", tokenStr)
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}
	//
	err, data := services.UserService.UserInfo(userId.(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, data)
	return
}

// ListPage godoc
// @Summary 管理员注销
// @Description 管理员注销
// @Tags 用户管理
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Security Auth
// @Success 200 {object} response.Response{} "success"
// @Router /admin_login/logout [get]
func (admin *adminController) AdminLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	// 解密token
	err, customClaims := services.JwtService.DecryptToken(token)
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}
	error := global.App.Redis.Del(context.Background(), customClaims.UserName).Err()
	if error != nil {
		response.BusinessFail(c, "用户注销失败")
		return
	}
	response.Success(c, "用户注销成功")
}

// ListPage godoc
// @Summary 管理员头像更新
// @Description 管理员头像更新
// @Tags 用户管理
// @ID /admin/avatar
// @Accept  json
// @Produce  json
// @Security Auth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response{} "success"
// @Router /admin/avatar [post]
func (admin *adminController) AdminInfoAvatar(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	// 解密token
	err, _ := services.JwtService.DecryptToken(token[len(common.TokenType)+1:])
	if err != nil {
		response.BusinessFail(c, "用户信息不存在")
		return
	}
	// username := customClaims.UserName
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BusinessFail(c, "参数错误")
		return
	}
	//获取文件名
	filename := header.Filename
	//写入文件
	out, err := os.Create(global.App.Config.Storage.Disks.LocalStorage.RootImageDir + filename)
	if err != nil {
		response.BusinessFail(c, "服务错误")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	url, err := utils.Upload2Ali(filename)
	if err != nil {
		response.BusinessFail(c, "图片上传失败")
		return
	}
	data := make(map[string]string)
	data["url"] = url
	response.Success(c, data)
	return
	// services.UserService.ModifyAvatar(username, imgdata)
}
