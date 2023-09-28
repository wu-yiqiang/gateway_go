package app

import (
	"gateway_go/app/common/request"
	"gateway_go/app/common/response"
	"gateway_go/app/models"
	"gateway_go/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type responseUser struct {
	models.User
	Menus   []string
	Routers []models.Routers
	services.TokenOutPut
}

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		// 存储到 redis
		// 查询用户信息
		err, userInfo := services.UserService.GetUserInfo(strconv.FormatUint(uint64(user.ID.ID), 10))
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
		var rolesArr = make([]models.Roles, 0)
		for _, v := range roleName {
			_, role := services.UserService.GetRoles(v)
			rolesArr = append(rolesArr, role)
		}
		// 通过角色查找路由
		var routerArr = make([]models.Routers, 0)
		var menuArr = make([]string, 0)
		var routerMap = map[string]struct{}{}
		var menuMap = map[string]struct{}{}
		for _, v := range rolesArr {
			var routerSlice = strings.Split(v.Router, ",")
			var menuSlice = strings.Split(v.Menu, ",")

			for _, routerId := range routerSlice {
				if _, ok := routerMap[routerId]; !ok {
					routerMap[routerId] = struct{}{}
					_, router := services.UserService.GetRouters(routerId)
					routerArr = append(routerArr, router)
				}
			}
			for _, menuId := range menuSlice {
				if _, ok := menuMap[menuId]; !ok {
					menuMap[menuId] = struct{}{}
					_, menu := services.UserService.GetMenus(menuId)
					menuArr = append(menuArr, menu.Name)
				}
			}
		}
		user := models.User{ID: userInfo.ID, Username: userInfo.Username, Nickname: userInfo.Nickname, Email: userInfo.Email, Phone: userInfo.Phone, Role: userInfo.Role}
		token := services.TokenOutPut{Token: tokenData.Token, Type: tokenData.Type, Expires: tokenData.Expires}
		resonseUserInfo := responseUser{User: user, TokenOutPut: token, Menus: menuArr, Routers: routerArr}
		response.Success(c, resonseUserInfo)
	}
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
