package services

import (
	"gateway_go/dao"
	"gateway_go/dto"
	"gateway_go/global"
	"gorm.io/gorm"
)

type servicesService struct {
}

var ServicesService = new(servicesService)

func (s *servicesService) TableName() string {
	return "gateway_service_info"
}

func (s *servicesService) ServiceDetail(search *dao.ServicesInfo) (*dao.ServicesDetail, error) {
	grpc, _ := GrpcRuleService.FindGrpcRule(search.ID)
	http, _ := HttpRuleService.FindHttpRule(search.ID)
	loadType, _ := LoadBalanceService.FindLoadBalance(search.ID)
	accessControl, _ := AccessControllService.FindAccessControl(search.ID)
	tcp, _ := TcpRuleService.FindTcpRule(search.ID)
	serviceDetail := &dao.ServicesDetail{
		Http:          &http,
		Tcp:           &tcp,
		Grpc:          &grpc,
		AccessControl: &accessControl,
		LoadBalance:   &loadType,
	}
	return serviceDetail, nil
}

func (s *servicesService) FindList(params *dto.ServicesListInput) ([]dao.ServicesInfo, int64, error) {
	list := []dao.ServicesInfo{}
	total := int64(0)
	offset := (params.PageNo - 1) * params.PageSize
	query := global.App.DB.Table(s.TableName()).Where("is_delete = ?", 0).Count(&total)
	if params.Info != "" {
		query.Where("(service_name like ? or service_desc like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(params.PageSize).Offset(offset)
	return list, total, nil
}

func (s *servicesService) ServiceDelete(id int) error {
	query := global.App.DB.Table(s.TableName()).Where("id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	update := query.Update("is_delete", 1)
	if update.Error != nil {
		return query.Error
	}
	return nil
}
