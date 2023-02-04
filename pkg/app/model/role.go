package model

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 角色
type Role struct {
	Id        int       `json:"id" gorm:"autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	GuardName string    `json:"guard_name" gorm:"size:100;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 模型角色关联表
type ModelHasRole struct {
	RoleId    int    `json:"role_id" gorm:"index:model_has_roles_model_id_model_type_index"`
	ModelType string `json:"model_type" gorm:"size:255;not null"`
	ModelId   int    `json:"model_id" gorm:"index:model_has_roles_model_id_model_type_index"`
}

// 角色权限关联表
type RoleHasPermission struct {
	PermissionId int `json:"permission_id" gorm:"size:11;not null"`
	RoleId       int `json:"role_id" gorm:"index:role_has_permissions_role_id_foreign"`
}

// 模型权限关联表
type ModelHasPermission struct {
	PermissionId int    `json:"permission_id" gorm:"index:model_has_permissions_model_id_model_type_index"`
	ModelType    string `json:"model_type" gorm:"size:255;not null"`
	ModelId      int    `json:"model_id" gorm:"index:model_has_permissions_model_id_model_type_index"`
}

// 获取角色列表
func (model *Role) List() (list []map[string]interface{}, Error error) {
	roles := []Role{}
	err := db.Client.Find(&roles).Error
	if err != nil {
		return list, err
	}

	for _, v := range roles {
		option := map[string]interface{}{
			"label": v.Name,
			"value": v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}

// 通过管理员ID获取角色相关
func (model *ModelHasRole) GetListByAdminId(id int) (modelHasRole *ModelHasRole, Error error) {
	err := db.Client.Where("model_id", id).Where("model_type", "admin").First(&modelHasRole).Error

	return modelHasRole, err
}

// 通过管理员ID获取角色id集合
func (model *ModelHasRole) GetRoleIdsByAdminId(id int) (roleIds []int, Error error) {
	err := db.Client.Model(model).Where("model_id", id).Where("model_type", "admin").Pluck("id", &roleIds).Error

	return roleIds, err
}

// 通过角色id集合获取权限id集合
func (model *RoleHasPermission) GetPermissionIdsByRoleIds(roleIds []int) (permissionIds []int, Error error) {
	// 角色权限id
	err := db.Client.Model(model).Where("role_id in (?)", roleIds).Pluck("id", &permissionIds).Error

	return permissionIds, err
}
