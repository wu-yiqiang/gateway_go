package services

import (
	"gateway_go/dao"
	"gateway_go/global"
	"gorm.io/gorm"
)

type accessControllService struct {
}

var AccessControllService = new(accessControllService)

func (a *accessControllService) TableName() string {
	return "gateway_service_access_control"
}

func (a *accessControllService) FindAccessControl(servicesId int64) (dao.AccessControl, error) {
	accessControl := dao.AccessControl{}
	err := global.App.DB.Table(a.TableName()).Where("service_id = ?", servicesId).Find(&accessControl).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return accessControl, err
	}
	return accessControl, nil
}
