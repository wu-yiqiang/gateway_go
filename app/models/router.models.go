package models

type Routers struct {
	Name     string `json:"name" gorm:"not null;comment:英文名"`
	CnName   string `json:"cn_name" gorm:"not null;comment:中文名"`
	Icon     string `json:"icon" gorm:"not null;comment:icon图标"`
	ParentId string `json:"parent_id" gorm:"not null;comment:父节点"`
}
