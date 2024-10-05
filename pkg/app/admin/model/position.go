package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/datetime"
)

// 职务
type Position struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Name      string            `json:"name" gorm:"size:500;not null"`
	Sort      int               `json:"sort" gorm:"size:11;default:0"`
	Status    int               `json:"status" gorm:"size:1;not null;default:1"`
	Remark    string            `json:"remark" gorm:"size:100"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}
