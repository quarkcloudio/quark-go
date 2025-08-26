package service

import (
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/transfer"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type PermissionService struct{}

// 初始化
func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

// 获取列表
func (p *PermissionService) List() (list []selectfield.Option, Error error) {
	permissions := []model.Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return list, err
	}

	for _, v := range permissions {
		option := selectfield.Option{
			Label: v.Name,
			Value: v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}

// 获取数据源
func (p *PermissionService) DataSource() (dataSource []transfer.DataSource, Error error) {
	permissions := []model.Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return dataSource, err
	}

	for _, v := range permissions {
		option := transfer.DataSource{
			Key:         v.Id,
			Title:       v.Name,
			Description: v.Remark,
		}
		dataSource = append(dataSource, option)
	}

	return dataSource, nil
}

// 通过权限id集合获取权限列表
func (p *PermissionService) GetListByIds(permissionIds interface{}) (permissions []model.Permission, Error error) {
	err := db.Client.Where("id in ?", permissionIds).Find(&permissions).Error

	return permissions, err
}

// 通过权限name集合获取权限列表
func (p *PermissionService) GetListByNames(permissionNames interface{}) (permissions []model.Permission, Error error) {
	err := db.Client.Where("name in ?", permissionNames).Find(&permissions).Error

	return permissions, err
}
