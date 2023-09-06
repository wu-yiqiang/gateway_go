package services

import (
	"errors"
	_"fmt"
	"gateway_go/app/common/request"
	"gateway_go/app/models"
	"gateway_go/global"
	"gateway_go/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// 用户注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("账号已存在")
		return
	}
	user = models.User{Username: params.Username, Password: utils.BcryptMake([]byte(params.Password)), Email: params.Email}
	err = global.App.DB.Create(&user).Error
	return
}

// 登陆获取Token
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}

