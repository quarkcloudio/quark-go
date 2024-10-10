package model

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/table"
	"github.com/quarkcloudio/quark-go/v3/pkg/dal/db"
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/datetime"
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

// 获取TreeSelect组件数据
func (model *Department) TreeSelect() (list []*treeselect.TreeData, Error error) {
	list = append(list, model.FindTreeSelectNode(0)...)

	return list, nil
}

// 递归获取TreeSelect组件数据
func (model *Department) FindTreeSelectNode(pid int) (list []*treeselect.TreeData) {
	departments := []Department{}
	db.Client.
		Where("pid = ?", pid).
		Where("status = ?", 1).
		Order("sort asc,id asc").
		Select("name", "id", "pid").
		Find(&departments)

	if len(departments) == 0 {
		return list
	}

	for _, v := range departments {
		item := &treeselect.TreeData{
			Value: v.Id,
			Title: v.Name,
		}

		children := model.FindTreeSelectNode(v.Id)
		if len(children) > 0 {
			item.Children = children
		}

		list = append(list, item)
	}

	return list
}

// 获取TableTree组件数据
func (model *Department) TableTree() (list []*table.TreeData) {
	list = append(list, model.FindTableTreeNode(0)...)

	return list
}

// 递归获取TableTree组件数据
func (model *Department) FindTableTreeNode(pid int) (list []*table.TreeData) {
	departments := []Department{}
	db.Client.
		Where("pid = ?", pid).
		Where("status = ?", 1).
		Order("sort asc,id asc").
		Select("name", "id", "pid").
		Find(&departments)

	if len(departments) == 0 {
		return list
	}

	for _, v := range departments {
		item := &table.TreeData{
			Key:   v.Id,
			Title: v.Name,
		}

		children := model.FindTableTreeNode(v.Id)
		if len(children) > 0 {
			item.Children = children
		}

		list = append(list, item)
	}

	return list
}

// 递归获取部门ids数据
func (model *Department) GetChildrenIds(pid int) (list []int) {
	departments := []Department{}
	db.Client.
		Where("pid = ?", pid).
		Where("status = ?", 1).
		Select("id", "pid").
		Find(&departments)

	if len(departments) == 0 {
		return list
	}

	for _, v := range departments {
		children := model.GetChildrenIds(v.Id)
		if len(children) > 0 {
			list = append(list, children...)
		}
		list = append(list, v.Id)
	}

	return list
}
