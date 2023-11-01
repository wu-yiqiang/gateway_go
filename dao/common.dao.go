package dao

import (
	_ "gorm.io/gorm"
	"time"
)

// 自增ID主键
type ID struct {
	Id int `json:"id" gorm:"primary_key" description:"自增主键"`
}

// 创建、更新时间
type Timestamps struct {
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

// 软删除
type IsDelete struct {
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}
