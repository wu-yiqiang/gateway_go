package dto

type RegisterInput struct {
	Username string `json:"username" gorm:"column:username" description:"用户名称" example:"admin"`
	Password string `json:"password" gorm:"column:password" description:"用户密码" example:"admin"`
}

//type RegisterOut struct {
//	Username:
//}
