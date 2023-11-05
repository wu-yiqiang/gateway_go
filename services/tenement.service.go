package services

import (
	"gateway_go/dao"
	"gateway_go/dto"
	"gateway_go/global"
	"gorm.io/gorm"
)

type tenementService struct{}

var TenementService = new(tenementService)

func (t *tenementService) FindTenementList(params dto.TenementListInput) ([]dao.Tenement, int64, error) {
	list := []dao.Tenement{}
	tenement := dao.Tenement{}
	total := int64(0)
	offset := (params.PageNo - 1) * params.PageSize
	query := global.App.DB.Table(tenement.TableName()).Where("is_delete = ?", 0).Count(&total)
	if params.Info != "" {
		query.Where("(app_id like ? or name like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(params.PageSize).Offset(offset)
	return list, total, nil
}

func (t *tenementService) TenementDelete(id int) error {
	tenment := dao.Tenement{}
	query := global.App.DB.Table(tenment.TableName()).Where("id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	update := query.Update("is_delete", 1)
	if update.Error != nil {
		return query.Error
	}
	return nil
}
