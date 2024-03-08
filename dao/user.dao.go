package dao

import (
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	Username  string    `json:"user_name" gorm:"column:user_name" description:"用户名称"`
	Password  string    `json:"password" gorm:"column:password" description:"用户密码"`
	Avator    string    `json:"avator" gorm:"column:avator" description:"头像"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (admin *Admin) GetName() string {
	return admin.Username
}
