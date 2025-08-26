package model

import (
	"github.com/quarkcloudio/quark-go/v4/utils/datetime"
)

// 字段
type Attachment struct {
	Id         int               `json:"id" gorm:"autoIncrement"`
	Uid        int               `json:"uid" gorm:"size:11;default:0"`
	Source     string            `json:"source" gorm:"size:255"`
	CategoryId int               `json:"category_id" gorm:"size:11;default:0"`
	Name       string            `json:"name" gorm:"size:255;not null"`
	Type       string            `json:"type" gorm:"size:255"`
	Sort       int               `json:"sort" gorm:"size:11;default:0"`
	Size       int64             `json:"size" gorm:"size:20;default:0"`
	Ext        string            `json:"ext" gorm:"size:255"`
	Path       string            `json:"path" gorm:"size:255;not null"`
	Url        string            `json:"url" gorm:"size:255;not null"`
	Hash       string            `json:"hash" gorm:"size:255;not null"`
	Extra      string            `json:"extra" gorm:"size:5000;not null"`
	Status     int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt  datetime.Datetime `json:"created_at"`
	UpdatedAt  datetime.Datetime `json:"updated_at"`
}
