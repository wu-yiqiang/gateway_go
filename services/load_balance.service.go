package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"gorm.io/gorm"
)

type loadBalanceService struct {
}

var LoadBalanceService = new(loadBalanceService)

func (l *loadBalanceService) TableName() string {
	return "gateway_service_load_balance"
}

func (l *loadBalanceService) FindLoadBalance(servicesId int64) (dao.LoadBalance, error) {
	loadBalance := dao.LoadBalance{}
	err := global.App.DB.Table(l.TableName()).Where("service_id = ?", servicesId).Find(&loadBalance).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return loadBalance, err
	}
	return loadBalance, nil
}
