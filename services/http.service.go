package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"gorm.io/gorm"
)

type httpRuleService struct {
}

var HttpRuleService = new(httpRuleService)

func (h *httpRuleService) TableName() string {
	return "gateway_service_http_rule"
}

func (h *httpRuleService) FindHttpRule(servicesId int64) (dao.HttpRule, error) {
	http := dao.HttpRule{}
	err := global.App.DB.Table(h.TableName()).Where("service_id = ?", servicesId).Find(&http).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return http, err
	}
	return http, nil
}
