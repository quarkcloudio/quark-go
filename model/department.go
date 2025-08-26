package model

import (
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/utils/datetime"
)

// 部门
type Department struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Pid       int               `json:"pid" gorm:"size:11;default:0"`
	Name      string            `json:"name" gorm:"size:500;not null"`
	Sort      int               `json:"sort" gorm:"size:11;default:0"`
	Status    int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}

// Seeder
func (model *Department) Seeder() {
	seeders := []Department{
		{Pid: 0, Name: "夸克云科技", Sort: 0, Status: 1},
		{Pid: 1, Name: "研发中心", Sort: 0, Status: 1},
		{Pid: 1, Name: "营销中心", Sort: 0, Status: 1},
	}

	db.Client.Create(&seeders)
}
