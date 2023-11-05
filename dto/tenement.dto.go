package dto

import (
	"gateway_go/dao"
	"gateway_go/request"
)

type TenementListInput struct {
	Info     string `json:"info" form:"info" gorm:"info" comment:"租户名称" binding:""`
	PageNo   int    `json:"page_no" form:"page_no" gorm:"page_no" comment:"页码" binding:"required,min=1"`
	PageSize int    `json:"page_size" form:"page_size" gorm:"page_size" comment:"页数" binding:"required,min=1"`
}

func (t TenementListInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"page_no.required":   "页码不能为空",
		"page_no.min":        "页码最小不能小于1",
		"page_size.required": "页数不能为空",
		"page_size.min":      "页数最小不能小于1",
	}
}

type TenementListOutput struct {
	Info  string          `json:"info" form:"info" gorm:"info" comment:"租户名称"`
	Total int64           `json:"total" form:"total" gorm:"total" comment:"总数"`
	List  *[]dao.Tenement `json:"list" form:"list" gorm:"list" comment:"列表"`
}
