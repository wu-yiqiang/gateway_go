package services

import (
	"errors"
	"gateway_go/dao"
	"gateway_go/global"
	"gateway_go/utils"
	"gateway_go/validator"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) TableName() string {
	return "gateway_admin"
}

// 用户注册
func (userService *userService) Register(params validator.Register) (err error, user dao.Admin) {
	var result = global.App.DB.Table(userService.TableName()).Where("user_name = ?", params.Username).Select("id").First(&dao.Admin{})
	if result.RowsAffected != 0 {
		err = errors.New("账号已存在")
		return
	}
	user = dao.Admin{Username: params.Username, Password: utils.BcryptMake([]byte(params.Password)), IsDelete: 0}
	err = global.App.DB.Table(userService.TableName()).Create(&user).Error
	if err != nil {
		return
	}
	return
}

//
//// 登陆获取Token
//func (userService *userService) Login(params validator.Login) (err error, user *dao.User) {
//	err = global.App.DB.Where("username = ?", params.Username).First(&user).Error
//	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
//		err = errors.New("用户名不存在或密码错误")
//	}
//	return
//}
//
//// GetUserInfo 获取用户信息
//func (userService *userService) GetUserInfo(id string) (err error, user dao.User) {
//	intId, err := strconv.Atoi(id)
//	if err = global.App.DB.First(&user, intId).Error; err != nil {
//		err = errors.New("用户不存在")
//		return
//	}
//	return
//}
//
//// 获取角色信息
//func (userService *userService) GetRoles(name string) (err error, roles dao.Roles) {
//	err = global.App.DB.Where("name = ?", name).First(&roles).Error
//	if err != nil {
//		err = errors.New("角色不存在")
//		return
//	}
//	return
//}
//
//// 获取路由信息
//func (userService *userService) GetRouters(id string) (err error, routers dao.Routers) {
//	err = global.App.DB.Where("id = ?", id).First(&routers).Error
//	if err != nil {
//		err = errors.New("路由不存在")
//		return
//	}
//	return
//}
//
//// 获取路由信息
//func (userService *userService) GetMenus(id string) (err error, menus dao.Menus) {
//	err = global.App.DB.Where("id = ?", id).First(&menus).Error
//	if err != nil {
//		err = errors.New("菜单不存在")
//		return
//	}
//	return
//}
