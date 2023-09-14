package models

import "strconv"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;comment:用户名称"`
	Nickname string `json:"nickname" gorm:"not null;comment:昵称"`
	Email    string `json:"email" gorm:"not null;index;comment:邮箱地址"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Phone    string `json:"phone" gorm:"not null;default:'';comment:电话号码"`
	Role     string `json:"role" gorm:"not null;default:'';comment:角色"`
	Timestamps
	IsDelete
}

type Roles struct {
	Name   string `json:"name" gorm:"not null;comment:英文名"`
	CnName string `json:"cn_name" gorm:"not null;comment:中文名"`
	Router string `json:"Router" gorm:"not null;comment:路由"`
	Menu   string `json:"menu" gorm:"not null;comment:菜单"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
