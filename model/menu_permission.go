package model

import (
	"github.com/quarkcloudio/quark-go/v4/utils/datetime"
)

// 字段
type MenuPermission struct {
	Id           int               `json:"id" gorm:"autoIncrement"`
	MenuId       int               `json:"menu_id" gorm:"size:11;not null"`
	PermissionId int               `json:"permission_id" gorm:"size:11;not null"`
	GuardName    string            `json:"guard_name" gorm:"size:20"`
	CreatedAt    datetime.Datetime `json:"created_at"`
	UpdatedAt    datetime.Datetime `json:"updated_at"`
}
