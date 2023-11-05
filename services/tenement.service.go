package services

import (
	"gateway_go/dao"
)

type tenementService struct {
}

var TenementService = new(tenementService)

func (t *tenementService) TableName() string {
	return "gateway_app"
}

func (h *tenementService) FindTenementList(servicesId int64) (*dao.Tenement, error) {
	tenement := &dao.Tenement{}
	return tenement, nil
}
