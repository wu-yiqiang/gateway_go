package models

import "strconv"
type User struct {
  ID
  Username string `json:"username" gorm:"not null;comment:用户名称"`
  nickname string `json:"username" gorm:"not null;comment:昵称"`
  Email string `json:"email" gorm:"not null;index;comment:邮箱地址"`
  Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
  Phone string `json:"phone" gorm:"not null;default:'';comment:电话号码"`
  Timestamps
  IsDelete
}

func (user User) GetUid() string {
  return strconv.Itoa(int(user.ID.ID))
}