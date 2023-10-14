package app

import (
	"context"
	"fmt"
	request2 "gateway_go/common/request"
	"gateway_go/common/response"
	"gateway_go/global"
	models2 "gateway_go/models"
	services2 "gateway_go/services"
	"gateway_go/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type responseUser struct {
	models2.User
	Menus   []string
	Routers []models2.Routers
	services2.TokenOutPut
}

func Login(c *gin.Context) {
	var form validator.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request2.GetErrorMsg(form, err))
		return
	}
	err, user := services2.UserService.Login(form)
	fmt.Println("login", user)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services2.JwtService.CreateToken(services2.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		// token存储到 redis
		err = global.App.Redis.Set(context.Background(), user.Username, tokenData.Token, 120*60*time.Second).Err()
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		// 查询用户信息
		err, userInfo := services2.UserService.GetUserInfo(strconv.FormatUint(uint64(user.ID.ID), 10))
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		var roleName = strings.Split(userInfo.Role, ",")
		if len(roleName) < 0 {
			// 获取用户权限
			response.Success(c, tokenData)
			return
		}
		// 通过角色名查询角色信息
		var rolesArr = make([]models2.Roles, 0)
		for _, v := range roleName {
			_, role := services2.UserService.GetRoles(v)
			rolesArr = append(rolesArr, role)
		}
		// 通过角色查找路由
		var routerArr = make([]models2.Routers, 0)
		var menuArr = make([]string, 0)
		var routerMap = map[string]struct{}{}
		var menuMap = map[string]struct{}{}
		for _, v := range rolesArr {
			var routerSlice = strings.Split(v.Router, ",")
			var menuSlice = strings.Split(v.Menu, ",")

			for _, routerId := range routerSlice {
				if _, ok := routerMap[routerId]; !ok {
					routerMap[routerId] = struct{}{}
					_, router := services2.UserService.GetRouters(routerId)
					routerArr = append(routerArr, router)
				}
			}
			for _, menuId := range menuSlice {
				if _, ok := menuMap[menuId]; !ok {
					menuMap[menuId] = struct{}{}
					_, menu := services2.UserService.GetMenus(menuId)
					menuArr = append(menuArr, menu.Name)
				}
			}
		}
		user := models2.User{ID: userInfo.ID, Username: userInfo.Username, Nickname: userInfo.Nickname, Email: userInfo.Email, Phone: userInfo.Phone, Role: userInfo.Role}
		token := services2.TokenOutPut{Token: tokenData.Token, Type: tokenData.Type, Expires: tokenData.Expires}
		resonseUserInfo := responseUser{User: user, TokenOutPut: token, Menus: menuArr, Routers: routerArr}
		response.Success(c, resonseUserInfo)
	}
}

func Logout(c *gin.Context) {
	err := services2.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

func Info(c *gin.Context) {
	err, user := services2.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
