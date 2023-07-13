package model

import (
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

// 权限
type Permission struct {
	Id        int       `json:"id" gorm:"autoIncrement"`
	Name      string    `json:"name" gorm:"size:500;not null"`
	GuardName string    `json:"group_name" gorm:"size:100;not null"`
	Path      string    `json:"path" gorm:"size:500;not null"`
	Method    string    `json:"method" gorm:"size:500;not null"`
	Remark    string    `json:"remark" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 获取列表
func (model *Permission) List() (list []*selectfield.Option, Error error) {
	permissions := []Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return list, err
	}

	for _, v := range permissions {
		option := &selectfield.Option{
			Label: v.Name,
			Value: v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}

// 通过权限id集合获取权限列表
func (model *Permission) GetListByIds(permissionIds interface{}) (permissions []*Permission, Error error) {
	err := db.Client.Where("id in ?", permissionIds).Find(&permissions).Error

	return permissions, err
}

// 通过权限name集合获取权限列表
func (model *Permission) GetListByNames(permissionNames interface{}) (permissions []*Permission, Error error) {
	err := db.Client.Where("name in ?", permissionNames).Find(&permissions).Error

	return permissions, err
}
