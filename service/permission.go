package service

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v4/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/transfer"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type PermissionService struct{}

// 初始化
func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

// 获取列表
func (p *PermissionService) List() (list []selectfield.Option, Error error) {
	permissions := []model.Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return list, err
	}

	for _, v := range permissions {
		option := selectfield.Option{
			Label: v.Name,
			Value: v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}

// 获取数据源
func (p *PermissionService) DataSource() (dataSource []transfer.DataSource, Error error) {
	permissions := []model.Permission{}
	err := db.Client.Find(&permissions).Error
	if err != nil {
		return dataSource, err
	}

	for _, v := range permissions {
		option := transfer.DataSource{
			Key:         v.Id,
			Title:       v.Name,
			Description: v.Remark,
		}
		dataSource = append(dataSource, option)
	}

	return dataSource, nil
}

// 通过权限id集合获取权限列表
func (p *PermissionService) GetListByIds(permissionIds interface{}) (permissions []model.Permission, Error error) {
	err := db.Client.Where("id in ?", permissionIds).Find(&permissions).Error

	return permissions, err
}

// 通过权限name集合获取权限列表
func (p *PermissionService) GetListByNames(permissionNames interface{}) (permissions []model.Permission, Error error) {
	err := db.Client.Where("name in ?", permissionNames).Find(&permissions).Error

	return permissions, err
}

func (p *PermissionService) HasAnyPermissions(modelId int, path string, method string) (hasPermission bool) {
	var permission model.Permission
	err := db.Client.Where("path = ? AND method = ?", path, method).First(&permission).Error
	if err != nil {
		return false
	}
	roleIds := []int{}
	err = db.Client.Model(&model.UserRole{}).Where("uid = ?", modelId).Pluck("role_id", &roleIds).Error
	if err != nil {
		return false
	}
	rolePermissions := []model.Permission{}
	db.Client.Model(&model.RolePermission{}).Where("role_id IN ?", roleIds).Find(&rolePermissions)

	for _, v := range rolePermissions {
		if v.Id == permission.Id {
			hasPermission = true
		}
	}

	return hasPermission
}

// 添加菜单拥有的权限
func (p *PermissionService) AddMenuPermission(menuId int, permissionIds interface{}) (err error) {
	permissions, err := NewPermissionService().GetListByIds(permissionIds)
	if err != nil {
		return err
	}

	menuPermissions := []model.MenuPermission{}

	// 菜单拥有的权限
	for _, v := range permissions {
		menuPermissions = append(menuPermissions, model.MenuPermission{
			MenuId:       menuId,
			PermissionId: v.Id,
		})
	}

	p.RemoveMenuPermissions(menuId)
	if len(menuPermissions) == 0 {
		return
	}

	err = db.Client.Create(&menuPermissions).Error
	if err != nil {
		return err
	}

	return
}

// 删除菜单拥有的权限
func (p *PermissionService) RemoveMenuPermissions(menuId int) (err error) {
	err = db.Client.Where("menu_id = ?", menuId).Delete(&model.MenuPermission{}).Error
	if err != nil {
		return err
	}
	return
}

// 获取菜单拥有的权限
func (p *PermissionService) GetMenuPermissions(menuId int) (permissions []model.Permission, err error) {
	permissionIds := []int{}
	menuPermissions := []model.MenuPermission{}
	err = db.Client.Model(&model.MenuPermission{}).Where("menu_id = ?", menuId).Find(&menuPermissions).Error
	if err != nil {
		return
	}
	for _, v := range menuPermissions {
		permissionIds = append(permissionIds, v.PermissionId)
	}

	permissions, err = NewPermissionService().GetListByIds(permissionIds)

	return
}

// 给角色添加菜单及权限
func (p *PermissionService) AddMenuAndPermissionToRole(roleId int, menuIds []int) (err error) {
	roleMenus := []model.RoleMenu{}
	rolePermissions := []model.RolePermission{}
	hasAddData := make(map[string]bool)

	// 角色拥有的菜单
	for _, v := range menuIds {
		rule := "role|" + strconv.Itoa(roleId) + "menu|" + strconv.Itoa(v)
		if !hasAddData[rule] {
			roleMenus = append(roleMenus, model.RoleMenu{
				RoleId: roleId,
				MenuId: v,
			})
			hasAddData[rule] = true
		}
	}

	// 角色拥有的权限
	for _, menuId := range menuIds {
		menuPermissions, err := p.GetMenuPermissions(menuId)
		if err == nil {
			for _, menuPermission := range menuPermissions {
				rule := "role|" + strconv.Itoa(roleId) + menuPermission.Path + menuPermission.Method
				if !hasAddData[rule] {
					rolePermissions = append(rolePermissions, model.RolePermission{
						RoleId:       roleId,
						PermissionId: menuPermission.Id,
					})
					hasAddData[rule] = true
				}
			}
		}
	}

	// 先清理数据
	err = p.RemoveRoleMenuAndPermissions(roleId)
	if err != nil {
		return err
	}
	if len(roleMenus) > 0 {
		err = db.Client.Create(&roleMenus).Error
		if err != nil {
			return err
		}
	}
	if len(rolePermissions) > 0 {
		err = db.Client.Create(&rolePermissions).Error
		if err != nil {
			return err
		}
	}

	return
}

// 删除角色拥有的菜单及权限
func (p *PermissionService) RemoveRoleMenuAndPermissions(roleId int) (err error) {
	err = db.Client.Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error
	if err != nil {
		return err
	}
	err = db.Client.Where("role_id = ?", roleId).Delete(&model.RolePermission{}).Error
	if err != nil {
		return err
	}
	return
}

// 获取角色拥有的菜单
func (p *PermissionService) GetRoleMenus(roleId int) (menus []model.Menu, err error) {
	menuIds := []interface{}{}
	db.Client.Model(&model.RoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &menuIds)
	menus, err = NewMenuService().GetListByIds(menuIds)
	return
}

// 获取角色拥有的权限
func (p *PermissionService) GetRolePermissions(roleId int) (permissions []model.Permission, err error) {
	permissionIds := []int{}
	db.Client.Model(&model.RolePermission{}).Where("role_id = ?", roleId).Pluck("permission_id", &permissionIds)
	permissions, err = NewPermissionService().GetListByIds(permissionIds)
	return
}

// 添加用户拥有的角色
func (p *PermissionService) AddUserRole(modelId int, roleIds []int) (err error) {
	userRoles := []model.UserRole{}
	hasAddData := make(map[string]bool)

	for _, v := range roleIds {
		rule := "role|" + strconv.Itoa(v)
		if !hasAddData[rule] {
			userRoles = append(userRoles, model.UserRole{
				Uid:    modelId,
				RoleId: v,
			})
			hasAddData[rule] = true
		}
	}

	p.RemoveUserRoles(modelId)

	if len(userRoles) > 0 {
		err = db.Client.Create(&userRoles).Error
		if err != nil {
			return err
		}
	}
	return
}

// 删除用户拥有的角色
func (p *PermissionService) RemoveUserRoles(modelId int) (err error) {
	err = db.Client.Where("uid = ?", modelId).Delete(&model.UserRole{}).Error
	if err != nil {
		return err
	}
	return
}

// 获取用户拥有的角色
func (p *PermissionService) GetUserRoles(modelId int) (roles []model.Role, err error) {
	roleIds := []int{}
	err = db.Client.Model(&model.UserRole{}).Where("uid = ?", modelId).Pluck("role_id", &roleIds).Error
	if err != nil {
		return roles, err
	}
	roles, err = NewRoleService().GetListByIds(roleIds)
	return roles, err
}

// 获取用户拥有的菜单
func (p *PermissionService) GetUserMenus(modelId int) (menus []model.Menu, err error) {
	getMenus := []model.Menu{}
	roles, err := p.GetUserRoles(modelId)
	if err != nil {
		return menus, err
	}

	for _, v := range roles {
		roleMenus, err := p.GetRoleMenus(v.Id)
		if err == nil {
			getMenus = append(getMenus, roleMenus...)
		}
	}

	return getMenus, nil
}

// 给角色添加数据权限
func (p *PermissionService) AddDepartmentToRole(roleId int, departmentIds []int) (err error) {
	roleDepartments := []model.RoleDepartment{}
	hasAddData := make(map[string]bool)

	// 角色拥有的菜单
	for _, v := range departmentIds {
		rule := "roleDepartment|" + strconv.Itoa(roleId) + strconv.Itoa(v)
		if !hasAddData[rule] {
			roleDepartments = append(roleDepartments, model.RoleDepartment{
				RoleId:       roleId,
				DepartmentId: v,
			})
			hasAddData[rule] = true
		}
	}

	// 先清理数据
	p.RemoveRoleDepartments(roleId)
	if len(roleDepartments) > 0 {
		err = db.Client.Create(&roleDepartments).Error
		if err != nil {
			return err
		}
	}

	return
}

// 删除角色拥有的部门
func (p *PermissionService) RemoveRoleDepartments(roleId int) (err error) {
	err = db.Client.Where("role_id = ?", roleId).Delete(&model.RoleDepartment{}).Error
	if err != nil {
		return err
	}
	return
}

// 获取角色拥有的部门Ids
func (p *PermissionService) GetRoleDepartmentIds(roleId int) (ids []int, err error) {
	err = db.Client.Model(&model.RoleDepartment{}).Where("role_id = ?", roleId).Pluck("department_id", &ids).Error
	return
}

// 获取角色拥有的部门
func (p *PermissionService) GetRoleDepartments(roleId int) (departments []model.Department, err error) {
	departmentIds, err := p.GetRoleDepartmentIds(roleId)
	if err != nil {
		return
	}
	return NewDepartmentService().GetListByIds(departmentIds)
}

// 获取用户拥有的部门
func (p *PermissionService) GetUserDepartments(modelId int) (menus []model.Department, err error) {
	getDepartments := []model.Department{}
	roles, err := p.GetUserRoles(modelId)
	if err != nil {
		return
	}

	userInfo, err := NewUserService().GetInfoById(modelId)
	if err != nil {
		return
	}

	for _, v := range roles {
		switch v.DataScope {
		case 1:
			departments, err := NewDepartmentService().GetList()
			if err == nil {
				getDepartments = append(getDepartments, departments...)
			}
		case 2:
			departments, err := p.GetRoleDepartments(v.Id)
			if err == nil {
				getDepartments = append(getDepartments, departments...)
			}
		case 3:
			department, err := NewDepartmentService().GetInfoById(userInfo.DepartmentId)
			if err == nil {
				getDepartments = append(getDepartments, department)
			}
		case 4:
			department, err := NewDepartmentService().GetInfoById(userInfo.DepartmentId)
			if err == nil {
				getDepartments = append(getDepartments, department)
			}
			departments := NewDepartmentService().GetChildrenDepartments(userInfo.DepartmentId)
			getDepartments = append(getDepartments, departments...)
		case 5:
		}
	}

	return getDepartments, nil
}
