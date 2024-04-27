package dto

import "gateway_go/request"

type AddFriend struct {
	UserId string `json:"userId" gorm:"column:user_id" description:"用户id" example:"admin" binding:"required"`
}

// 自定义错误信息
func (add AddFriend) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"userId.required": "用户ID不能为空",
	}
}
