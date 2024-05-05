package model

import (
	"gorm.io/plugin/soft_delete"
)

type UserFriend struct {
	ID          int32                 `json:"id" gorm:"primarykey"`
	CreatedTime int64                 `json:"created_time"`
	UpdatedTime int64                 `json:"updated_time"`
	ISDeleted   soft_delete.DeletedAt `json:"is_deleted"`
	UserId      int32                 `json:"userId" gorm:"index;comment:'用户ID'"`
	FriendId    int32                 `json:"friendId" gorm:"index;comment:'好友ID'"`
}
