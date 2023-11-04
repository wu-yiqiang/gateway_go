package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"gorm.io/gorm"
)

type tcpRuleService struct {
}

var TcpRuleService = new(tcpRuleService)

func (h *tcpRuleService) TableName() string {
	return "gateway_service_tcp_rule"
}

func (h *tcpRuleService) FindTcpRule(servicesId int64) (dao.TcpRule, error) {
	tcp := dao.TcpRule{}
	err := global.App.DB.Table(h.TableName()).Where("service_id = ?", servicesId).Find(&tcp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tcp, err
	}
	return tcp, nil
}
