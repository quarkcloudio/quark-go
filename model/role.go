package model

import (
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/utils/datetime"
)

// 角色
type Role struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Name      string            `json:"name" gorm:"size:255;not null"`
	DataScope int               `json:"data_scope" gorm:"size:1;not null;default:1"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	GuardName string            `json:"guard_name" gorm:"size:100;not null"`
	Status    int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}

// Seeder
func (model *Role) Seeder() {
	seeders := []Role{
		{Name: "普通角色", GuardName: "admin", DataScope: 1},
	}

	db.Client.Create(&seeders)
}
