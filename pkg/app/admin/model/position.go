package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/checkbox"
	"github.com/quarkcloudio/quark-go/v3/pkg/dal/db"
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

// Seeder
func (model *Position) Seeder() {
	seeders := []Position{
		{Name: "董事长", Sort: 0, Status: 1},
		{Name: "项目经理", Sort: 0, Status: 1},
		{Name: "普通员工", Sort: 0, Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取角色列表
func (model *Position) List() (list []*checkbox.Option, Error error) {
	positions := []Position{}
	err := db.Client.Find(&positions).Error
	if err != nil {
		return list, err
	}
	for _, v := range positions {
		list = append(list, &checkbox.Option{
			Label: v.Name,
			Value: v.Id,
		})
	}
	return list, nil
}
