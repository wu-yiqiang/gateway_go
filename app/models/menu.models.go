package models

type Menus struct {
	Name string `json:"name" gorm:"not null;comment:英文名"`
}
