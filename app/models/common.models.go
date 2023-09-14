package models

import (
	_ "gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

// 软删除
type IsDelete struct {
	IsDelete int `json:"is_deleted" gorm:"index"`
}
