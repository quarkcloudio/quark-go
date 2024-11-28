package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/checkbox"
)

type RoleService struct{}

// 初始化
func NewRoleService() *RoleService {
	return &RoleService{}
}

// 通过ID获取角色信息
func (p *RoleService) GetInfoById(id interface{}) (role *model.Role, Error error) {
	err := db.Client.Where("id = ?", id).First(&role).Error

	return role, err
}

// 更新角色数据范围
func (p *RoleService) UpdateRoleDataScope(roleId int, dataScope int, departmentIds []int) (err error) {
	err = db.Client.Model(&model.Role{}).Where("id = ?", roleId).Update("data_scope", dataScope).Error
	if err == nil {
		if dataScope == 2 {
			NewCasbinService().AddDepartmentToRole(roleId, departmentIds)
		} else {
			NewCasbinService().RemoveRoleDepartments(roleId)
		}
	}

	return err
}

// 获取角色列表
func (p *RoleService) List() (list []checkbox.Option, Error error) {
	roles := []model.Role{}
	err := db.Client.Where("status = ?", 1).Find(&roles).Error
	if err != nil {
		return list, err
	}

	for _, v := range roles {
		list = append(list, checkbox.Option{
			Label: v.Name,
			Value: v.Id,
		})
	}

	return list, nil
}

// 通过id集合获取列表
func (model *RoleService) GetListByIds(ids interface{}) (roles []model.Role, Error error) {
	err := db.Client.Where("id in ?", ids).Find(&roles).Error

	return roles, err
}
