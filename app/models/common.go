package models

import (
  _"gorm.io/gorm"
  "time"
)

// 自增ID主键
type ID struct {
  ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
  CreatedTime time.Time `json:"created_time"`
  UpdatedTime time.Time `json:"updated_time"`
}

// 软删除
type IsDeleted struct {
  IsDeleted bool `json:"is_deleted" gorm:"index"`
}