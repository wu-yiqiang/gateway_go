package services

import (
	"gateway_go/dao"
	"gateway_go/dto"
	"gateway_go/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type servicesService struct {
}

var ServicesService = new(servicesService)

func (s *servicesService) TableName() string {
	return "gateway_service_info"
}

func (s *servicesService) FindList(c *gin.Context, params *dto.ServicesListInput) ([]dao.ServicesInfo, int64, error) {
	list := []dao.ServicesInfo{}
	total := int64(0)
	offset := (params.PageNo - 1) * params.PageSize
	query := global.App.DB.Table(s.TableName()).Where("is_delete = ?", 0)
	if params.Info != "" {
		query.Where("(service_name like ? or service_desc like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(params.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}
