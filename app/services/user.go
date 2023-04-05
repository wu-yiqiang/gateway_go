package services

import (
	"errors"
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
	var result = global.App.DB.Where("name = ?", params.Name).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("账号已存在")
		return
	}
	user = models.User{Name: params.Name, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}