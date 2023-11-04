package services

import (
	"fmt"
	"gateway_go/dao"
	"gateway_go/global"
	"gorm.io/gorm"
)

type grpcRuleService struct {
}

var GrpcRuleService = new(grpcRuleService)

func (g *grpcRuleService) TableName() string {
	return "gateway_service_grpc_rule"
}

func (g *grpcRuleService) FindGrpcRule(servicesId int64) (dao.GrpcRule, error) {
	grpc := dao.GrpcRule{}
	err := global.App.DB.Table(g.TableName()).Where("service_id = ?", servicesId).Find(&grpc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("&& err != gorm.ErrRecordNotFound", err)
		return grpc, err
	}
	return grpc, nil
}
