package dto

type RegisterInput struct {
	Username string `json:"username" gorm:"column:username" description:"用户名称" example:"admin"`
	Password string `json:"password" gorm:"column:password" description:"用户密码" example:"1234_abcd"`
}

type LoginInput struct {
	Username string `json:"username" gorm:"column:username" description:"用户名称" example:"admin"`
	Password string `json:"password" gorm:"column:password" description:"用户密码" example:"1234_abcd"`
}

type LoginOutput struct {
	Token string `json:"token" gorm:"column:token" description:"token" example:""`
}

type ChangePasswordInput struct {
	Username    string `form:"username" gorm:"column:username" description:"用户名称" example:"admin"`
	Password    string `form:"password" gorm:"column:password" description:"旧密码" example:"1234_abcd"`
	NewPassword string `form:"newpassword" gorm:"column:newpassword" description:"新密码" example:"1234_abcd"`
}

type AdminInfoOutput struct {
	Avatar       string   `json:"avatar"`
	Id           int64    `json:"id"`
	Introduction string   `json:"introduction"`
	LoginTime    string   `json:"login_time"`
	Name         string   `json:"name"`
	Roles        []string `json:"roles"`
}

//type RegisterOut struct {
//	Username:
//}
