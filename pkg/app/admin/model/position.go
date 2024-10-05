package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/datetime"
)

// 职务
type Position struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Name      string            `json:"name" gorm:"size:500;not null"`
	GuardName string            `json:"group_name" gorm:"size:100;not null"`
	Path      string            `json:"path" gorm:"size:500;not null"`
	Method    string            `json:"method" gorm:"size:500;not null"`
	Remark    string            `json:"remark" gorm:"size:100"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}
