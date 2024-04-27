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
	Uuid        string `json:"id" gorm:"uuid" description:"用户ID"`
	Username    string `json:"username" gorm:"column:username" description:"用户名称"`
	Avatar      string `json:"avatar" gorm:"column:avatar" description:"头像"`
	Nickname    string `json:"nickname" gorm:"column:nickname" description:"昵称"`
	Email       string `json:"email" gorm:"column:email" description:"邮箱"`
	Phone       string `json:"phone" gorm:"column:phone" description:"电话"`
	Role        string `json:"role" gorm:"column:role" description:"角色"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time" description:"更新时间"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time" description:"创建时间"`
}

type AdminInfoAvatar struct {
	Avatar []byte `json:"avatar"  gorm:"column:avatar" binding:"required"`
}

func (adminInfoAvator AdminInfoAvatar) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"avatar.required": "文件不能为空",
	}
}

type QueryUser struct {
	Username string `json:"username" gorm:"column:username" description:"用户名称"`
}

func (queryUser QueryUser) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required": "用户名不能为空",
	}
}
