package request

type Register struct {
  Username string `form:"username" json:"username" binding:"required"`
  Email string `form:"email" json:"email" binding:"required,email"`
  Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
  return ValidatorMessages{
    "username.required": "用户名称不能为空",
    "email.required": "邮箱地址不能为空",
    "email.email": "邮箱格式不正确",
    "password.required": "用户密码不能为空",
    }
}


type Login struct {
  Email string `form:"email" json:"email" binding:"required,email"`
  Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
  return ValidatorMessages{
    "email.required": "邮箱不能为空",
    "email.email": "邮箱格式不正确",
    "password.required": "用户密码不能为空",
    }
}
