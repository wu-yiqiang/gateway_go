package services

import (
	"errors"
	"gateway_go/app/common/request"
	"gateway_go/app/models"
	"gateway_go/global"
	"gateway_go/utils"
	"strconv"
	"time"
)

type userService struct {
}

var UserService = new(userService)

// 用户注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("username = ?", params.Username).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("账号已存在")
		return
	}
	user = models.User{Username: params.Username, Password: utils.BcryptMake([]byte(params.Password)), Email: params.Email, Phone: params.Phone, Nickname: params.Nickname, Timestamps: models.Timestamps{CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}}
	err = global.App.DB.Create(&user).Error
	return
}

// 登陆获取Token
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("username = ?", params.Username).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	if err = global.App.DB.First(&user, intId).Error; err != nil {
		err = errors.New("用户不存在")
		return
	}
	return
}

// 获取角色信息
func (userService *userService) GetRoles(name string) (err error, roles models.Roles) {
	err = global.App.DB.Where("name = ?", name).First(&roles).Error
	if err != nil {
		err = errors.New("角色不存在")
		return
	}
	return
}

// 获取路由信息
func (userService *userService) GetRouters(id string) (err error, routers models.Routers) {
	err = global.App.DB.Where("id = ?", id).First(&routers).Error
	if err != nil {
		err = errors.New("路由不存在")
		return
	}
	return
}

// 获取路由信息
func (userService *userService) GetMenus(id string) (err error, menus models.Menus) {
	err = global.App.DB.Where("id = ?", id).First(&menus).Error
	if err != nil {
		err = errors.New("菜单不存在")
		return
	}
	return
}