package model

import (
	"strings"
	"time"

	"github.com/go-basic/uuid"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/lister"
)

// 字段
type Menu struct {
	Key        string    `json:"key" gorm:"<-:false"`
	Id         int       `json:"id" gorm:"autoIncrement"`
	Name       string    `json:"name" gorm:"size:100;not null"`
	GuardName  string    `json:"group_name" gorm:"size:100;not null"`
	Icon       string    `json:"icon" gorm:"size:100;"`
	Type       string    `json:"type" gorm:"size:100;not null"`
	Pid        int       `json:"pid" gorm:"size:11;default:0"`
	Sort       int       `json:"sort" gorm:"size:11;default:0"`
	Path       string    `json:"path" gorm:"size:255"`
	Show       int       `json:"show" gorm:"size:1;not null;default:1"`
	Status     int       `json:"status" gorm:"size:1;not null;default:1"`
	Locale     string    `json:"locale" gorm:"<-:false"`
	HideInMenu bool      `json:"hide_in_menu" gorm:"<-:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// 菜单表
func (p *Menu) Seeder() {
	seeders := []Menu{
		{Id: 1, Name: "控制台", GuardName: "admin", Icon: "icon-home", Type: "default", Pid: 0, Sort: 0, Path: "/dashboard", Show: 1, Status: 1},
		{Id: 2, Name: "主页", GuardName: "admin", Icon: "", Type: "engine", Pid: 1, Sort: 0, Path: "/api/admin/dashboard/index/index", Show: 1, Status: 1},
		{Id: 3, Name: "管理员", GuardName: "admin", Icon: "icon-admin", Type: "default", Pid: 0, Sort: 0, Path: "/admin", Show: 1, Status: 1},
		{Id: 4, Name: "管理员列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 3, Sort: 0, Path: "/api/admin/admin/index", Show: 1, Status: 1},
		{Id: 5, Name: "权限列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 3, Sort: 0, Path: "/api/admin/permission/index", Show: 1, Status: 1},
		{Id: 6, Name: "角色列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 3, Sort: 0, Path: "/api/admin/role/index", Show: 1, Status: 1},
		{Id: 7, Name: "系统配置", GuardName: "admin", Icon: "icon-setting", Type: "default", Pid: 0, Sort: 0, Path: "/system", Show: 1, Status: 1},
		{Id: 8, Name: "设置管理", GuardName: "admin", Icon: "", Type: "default", Pid: 7, Sort: 0, Path: "/system/config", Show: 1, Status: 1},
		{Id: 9, Name: "网站设置", GuardName: "admin", Icon: "", Type: "engine", Pid: 8, Sort: 0, Path: "/api/admin/webConfig/setting/form", Show: 1, Status: 1},
		{Id: 10, Name: "配置管理", GuardName: "admin", Icon: "", Type: "engine", Pid: 8, Sort: 0, Path: "/api/admin/config/index", Show: 1, Status: 1},
		{Id: 11, Name: "菜单管理", GuardName: "admin", Icon: "", Type: "engine", Pid: 7, Sort: 0, Path: "/api/admin/menu/index", Show: 1, Status: 1},
		{Id: 12, Name: "操作日志", GuardName: "admin", Icon: "", Type: "engine", Pid: 7, Sort: 0, Path: "/api/admin/actionLog/index", Show: 1, Status: 1},
		{Id: 13, Name: "附件空间", GuardName: "admin", Icon: "icon-attachment", Type: "default", Pid: 0, Sort: 0, Path: "/attachment", Show: 1, Status: 1},
		{Id: 14, Name: "文件管理", GuardName: "admin", Icon: "", Type: "engine", Pid: 13, Sort: 0, Path: "/api/admin/file/index", Show: 1, Status: 1},
		{Id: 15, Name: "图片管理", GuardName: "admin", Icon: "", Type: "engine", Pid: 13, Sort: 0, Path: "/api/admin/picture/index", Show: 1, Status: 1},
		{Id: 16, Name: "我的账号", GuardName: "admin", Icon: "icon-user", Type: "default", Pid: 0, Sort: 0, Path: "/account", Show: 1, Status: 1},
		{Id: 17, Name: "个人设置", GuardName: "admin", Icon: "", Type: "engine", Pid: 16, Sort: 0, Path: "/api/admin/account/setting/form", Show: 1, Status: 1},
	}

	db.Client.Create(&seeders)
}

// 获取菜单的有序列表
func (model *Menu) OrderedList() (list []map[string]interface{}, Error error) {
	var menus []map[string]interface{}
	err := db.Client.
		Model(&model).
		Where("guard_name = ?", "admin").
		Order("sort asc,id asc").
		Find(&menus).Error
	if err != nil {
		return list, err
	}

	menuTrees, err := lister.ListToTree(menus, "id", "pid", "children", 0)
	if err != nil {
		return list, err
	}

	menuTreeList, err := lister.TreeToOrderedList(menuTrees, 0, "name", "children")
	if err != nil {
		return list, err
	}

	list = append(list, map[string]interface{}{
		"label": "根节点",
		"value": 0,
	})
	for _, v := range menuTreeList {
		option := map[string]interface{}{
			"label": v.((map[string]interface{}))["name"],
			"value": v.(map[string]interface{})["id"],
		}
		list = append(list, option)
	}

	return list, nil
}

// 获取菜单的tree
func (model *Menu) Tree() (list []interface{}, Error error) {
	menus := []Menu{}
	err := db.Client.Where("status = ?", 1).Select("name", "id", "pid").Find(&menus).Error
	if err != nil {
		return list, err
	}

	menuList := []map[string]interface{}{}
	for _, v := range menus {
		item := map[string]interface{}{
			"key":   v.Id,
			"pid":   v.Pid,
			"title": v.Name,
		}
		menuList = append(menuList, item)
	}

	return lister.ListToTree(menuList, "key", "pid", "children", 0)
}

// 通过管理员ID权限菜单
func (model *Menu) GetListByAdminId(adminId int) (menuList interface{}, Error error) {
	menus := []Menu{}
	var menuKey int

	if adminId == 1 {
		db.Client.Where("status = ?", 1).Where("guard_name", "admin").Order("sort asc").Find(&menus)
	} else {
		var menuIds []int
		permissions, err := (&Permission{}).GetListByAdminId(adminId)
		if err != nil {
			return menuList, err
		}

		if permissions != nil {
			for key, v := range permissions {
				menuIds[key] = v.MenuId
			}
		}

		var pids1 []int
		// 三级查询列表
		db.Client.
			Where("status = ?", 1).
			Where("id in (?)", menuIds).
			Where("pid <> ?", 0).
			Order("sort asc").
			Find(&menus)
		for key, v := range menus {
			if v.Pid != 0 {
				pids1[key] = v.Pid
			}
			menuKey = key
		}

		var pids2 []int
		menu2 := []Menu{}
		// 二级查询列表
		db.Client.
			Where("status = ?", 1).
			Where("id in (?)", pids1).
			Where("pid <> ?", 0).
			Order("sort asc").
			Find(&menu2)
		for key, v := range menu2 {
			if v.Pid != 0 {
				pids2[key] = v.Pid
			}
			menuKey = menuKey + key
			menus[menuKey] = v
		}

		menu3 := []Menu{}
		// 一级查询列表
		db.Client.
			Where("status = ?", 1).
			Where("id in (?)", pids2).
			Where("pid", 0).
			Order("sort asc").
			Find(&menu3)
		for key, v := range menu3 {
			menuKey = menuKey + key
			menus[menuKey] = v
		}
	}

	for k, v := range menus {
		v.Key = uuid.New()
		v.Locale = "menu" + strings.Replace(v.Path, "/", ".", -1)

		if v.Show == 1 {
			v.HideInMenu = false
		} else {
			v.HideInMenu = true
		}

		if v.Type == "engine" {
			v.Path = "/index?api=" + v.Path
		}

		menus[k] = v
	}

	return lister.ListToTree(menus, "id", "pid", "routes", 0)
}
