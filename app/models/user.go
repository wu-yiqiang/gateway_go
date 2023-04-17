package models

import "strconv"
type User struct {
  ID
  Username string `json:"username" gorm:"not null;comment:用户名称"`
  Email string `json:"email" gorm:"not null;index;comment:邮箱地址"`
  Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
  Timestamps
  IsDeleted
}

func (user User) GetUid() string {
  return strconv.Itoa(int(user.ID.ID))
}