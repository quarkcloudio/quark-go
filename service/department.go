package service

import (
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type DepartmentService struct{}

// 初始化
func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

// 递归获取部门ids数据
func (p *DepartmentService) GetChildrenIds(pid int) (list []int) {
	departments := []model.Department{}
	db.Client.
		Where("pid = ?", pid).
		Where("status = ?", 1).
		Select("id", "pid").
		Find(&departments)

	if len(departments) == 0 {
		return list
	}

	for _, v := range departments {
		children := p.GetChildrenIds(v.Id)
		if len(children) > 0 {
			list = append(list, children...)
		}
		list = append(list, v.Id)
	}

	return list
}

// 递归获取部门数据
func (p *DepartmentService) GetChildrenDepartments(pid int) (list []model.Department) {
	departments := []model.Department{}
	db.Client.
		Where("pid = ?", pid).
		Where("status = ?", 1).
		Select("id", "pid").
		Find(&departments)

	if len(departments) == 0 {
		return list
	}

	for _, v := range departments {
		children := p.GetChildrenDepartments(v.Id)
		if len(children) > 0 {
			list = append(list, children...)
		}
		list = append(list, v)
	}

	return list
}

// 通过ID获取信息
func (model *DepartmentService) GetInfoById(id interface{}) (department model.Department, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&department).Error

	return department, err
}

// 获取列表
func (model *DepartmentService) GetList() (departments []model.Department, Error error) {
	err := db.Client.Where("status = ?", 1).Find(&departments).Error

	return departments, err
}

// 通过id集合获取列表
func (model *DepartmentService) GetListByIds(ids interface{}) (departments []model.Department, Error error) {
	err := db.Client.Where("id in ?", ids).Where("status = ?", 1).Find(&departments).Error

	return departments, err
}
