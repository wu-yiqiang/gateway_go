package services

import (
	"errors"
	_"fmt"
	"gateway_go/app/common/request"
	"gateway_go/app/models"
	"gateway_go/global"
	"gateway_go/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
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