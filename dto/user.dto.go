package dto

import "gateway_go/request"

type RegisterInput struct {
	Username string `json:"username" gorm:"column:username" json:"username" description:"用户名称" example:"admin" binding:"required"`
	Password string `json:"password" gorm:"column:password" json:"password" description:"用户密码" example:"1234_abcd" binding:"required,password"`
}

//type LoginInput struct {
//	Username string `json:"username" gorm:"column:username" description:"用户名称" example:"admin"`
//	Password string `json:"password" gorm:"column:password" description:"用户密码" example:"1234_abcd"`
//}

// 自定义错误信息
func (register RegisterInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required": "用户名称不能为空",
		"password.required": "用户密码不能为空",
		"password.password": "用户密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}

type LoginOutput struct {
	Token string `json:"token" gorm:"column:token" description:"token" example:""`
}

type ChangePasswordInput struct {
	Username    string `form:"username" gorm:"column:username" json:"username" description:"用户名称" example:"admin" binding:"required"`
	Password    string `form:"password" gorm:"column:password" json:"password" description:"旧密码" example:"1234_abcd" binding:"required,password"`
	NewPassword string `form:"newpassword" gorm:"column:newpassword" json:"newpassword" description:"新密码" example:"1234_abcd" binding:"required,password"`
}

func (changePassword ChangePasswordInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required":    "用户名不能为空",
		"password.required":    "旧密码不能为空",
		"password.password":    "旧密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
		"newpassword.required": "新密码不能为空",
		"newpassword.password": "新密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}

type AdminInfoOutput struct {
	Avatar       string   `json:"avatar"`
	Id           int64    `json:"id"`
	Introduction string   `json:"introduction"`
	LoginTime    string   `json:"login_time"`
	Name         string   `json:"name"`
	Roles        []string `json:"roles"`
}
