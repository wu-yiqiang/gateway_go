package validator

import "gateway_go/common/request"

type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,password"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"mobile"`
}

// 自定义错误信息
func (register Register) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required": "用户名称不能为空",
		"email.required":    "邮箱地址不能为空",
		"email.email":       "邮箱地址不合法",
		"phone.mobile":      "号码不合法",
		"password.required": "用户密码不能为空",
		"password.password": "用户密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func (login Login) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"username.required": "用户名不能为空",
		"password.required": "用户密码不能为空",
		"password.password": "用户密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}

type ChangePassword struct {
	UserId      int64  `form:"userid" json:"userid" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required,password"`
	NewPassword string `form:"newpassword" json:"newpassword" binding:"required,password"`
}

func (changePassword ChangePassword) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"userid.required":      "用户id不能为空",
		"password.required":    "旧密码不能为空",
		"password.password":    "旧密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
		"newpassword.required": "新密码不能为空",
		"newpassword.password": "新密码必须8-16位，必须包含有一个大写字母，一个小写字母，一个数字",
	}
}
