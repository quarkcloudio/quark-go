package model

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/form/fields/checkbox"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	"github.com/quarkcloudio/quark-go/v2/pkg/utils/datetime"
)

// 角色
type Role struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Name      string            `json:"name" gorm:"size:255;not null"`
	GuardName string            `json:"guard_name" gorm:"size:100;not null"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}

// 获取角色列表
func (model *Role) List() (list []*checkbox.Option, Error error) {
	roles := []Role{}
	err := db.Client.Find(&roles).Error
	if err != nil {
		return list, err
	}

	for _, v := range roles {
		list = append(list, &checkbox.Option{
			Label: v.Name,
			Value: v.Id,
		})
	}

	return list, nil
}

// 通过id集合获取列表
func (model *Role) GetListByIds(ids interface{}) (roles []*Role, Error error) {
	err := db.Client.Where("id in ?", ids).Find(&roles).Error

	return roles, err
}
