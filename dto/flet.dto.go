package dto

import "gateway_go/request"

type FletWeatherInput struct {
	Location string `json:"location" form:"location" description:"位置" example:"上海" binding:"required"`
}

func (fletWeatherInput FletWeatherInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"location.required": "位置不能为空",
	}
}

type FletNewsInput struct {
	Type     string `json:"type" form:"type" gorm:"type" comment:"类型" binding:""`
	PageNo   int    `json:"page_no" form:"page_no" gorm:"page_no" comment:"页码" default:"1" binding:"required,min=1"`
	PageSize int    `json:"page_size" form:"page_size" gorm:"page_size" comment:"页数" default:"10" binding:"required,min=1"`
}

func (fletNewsInput FletNewsInput) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"type.required":      "类型不能为空",
		"page_no.required":   "页码不能为空",
		"page_no.min":        "页码最小不能小于1",
		"page_size.required": "页数不能为空",
		"page_size.min":      "页数最小不能小于1",
	}
}
