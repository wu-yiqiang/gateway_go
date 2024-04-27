package dao

type Admin struct {
	Id          int    `json:"id" gorm:"primary_key" description:"自增主键"`
	Uuid        string `json:"uuid" gorm:"uuid" description:"uuid"`
	Username    string `json:"username" gorm:"column:username" description:"用户名称"`
	Password    string `json:"password" gorm:"column:password" description:"用户密码"`
	Avatar      string `json:"avatar" gorm:"column:avatar" description:"头像"`
	Nickname    string `json:"nickname" gorm:"column:nickname" description:"昵称"`
	Email       string `json:"email" gorm:"column:email" description:"邮箱"`
	Phone       string `json:"phone" gorm:"column:phone" description:"电话"`
	Role        string `json:"role" gorm:"column:role" description:"角色"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time" description:"更新时间"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time" description:"创建时间"`
	IsDelete    int    `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (admin *Admin) GetUsername() string {
	return admin.Username
}
func (admin *Admin) GetPassword() string {
	return admin.Password
}

func (admin *Admin) GetUuid() string {
	return admin.Uuid
}
