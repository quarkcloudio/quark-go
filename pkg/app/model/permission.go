package model

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selectfield"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 权限
type Permission struct {
	Id        int       `json:"id" gorm:"autoIncrement"`
	MenuId    int       `json:"menu_id" gorm:"size:11;default:0"`
	Name      string    `json:"name" gorm:"size:500;not null"`
	GuardName string    `json:"group_name" gorm:"size:100;not null"`
	Path      string    `json:"path" gorm:"size:500;not null"`
	Method    string    `json:"method" gorm:"size:500;not null"`
	Remark    string    `json:"remark" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 获取列表
func (model *Permission) List() (list []*selectfield.Option, Error error) {
	permissions := []Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return list, err
	}

	for _, v := range permissions {
		option := &selectfield.Option{
			Label: v.Name,
			Value: v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}

// 通过权限id集合获取权限列表
func (model *Permission) GetListByIds(permissionIds []int) (permissions []Permission, Error error) {
	err := db.Client.Where("id in ?", permissionIds).Find(&permissions).Error

	return permissions, err
}

// 通过管理员ID获取权限列表
func (model *Permission) GetListByAdminId(id int) (permissions []Permission, Error error) {

	// 管理员拥有的角色id集合
	roleIds, err := (&ModelHasRole{}).GetRoleIdsByAdminId(id)
	if err != nil {
		return permissions, nil
	}

	// 角色拥有的权限id集合
	permissionIds, err := (&RoleHasPermission{}).GetPermissionIdsByRoleIds(roleIds)
	if err != nil {
		return permissions, err
	}

	// 角色权限列表
	return (&Permission{}).GetListByIds(permissionIds)
}
