package model

import (
	"strings"
	"time"

	"github.com/go-basic/uuid"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/tree"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/quarkcms/quark-go/v2/pkg/utils/lister"
	"gorm.io/gorm"
)

// 字段
type Menu struct {
	Key        string    `json:"key" gorm:"<-:false"`
	Id         int       `json:"id" gorm:"autoIncrement"`
	Name       string    `json:"name" gorm:"size:100;not null"`
	GuardName  string    `json:"group_name" gorm:"size:100;not null"`
	Icon       string    `json:"icon" gorm:"size:100;"`
	Type       int       `json:"type" gorm:"size:100;not null"` // 菜单类型：1目录，2菜单，3按钮
	Pid        int       `json:"pid" gorm:"size:11;default:0"`
	Sort       int       `json:"sort" gorm:"size:11;default:0"`
	Path       string    `json:"path" gorm:"size:255"`
	Show       int       `json:"show" gorm:"size:1;not null;default:1"`
	IsEngine   int       `json:"is_engine" gorm:"size:1;not null;default:0"`
	IsLink     int       `json:"is_link" gorm:"size:1;not null;default:0"`
	Status     int       `json:"status" gorm:"size:1;not null;default:1"`
	Locale     string    `json:"locale" gorm:"<-:false"`
	HideInMenu bool      `json:"hide_in_menu" gorm:"<-:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// 菜单表
func (p *Menu) Seeder() {
	seeders := []Menu{
		{Id: 1, Name: "控制台", GuardName: "admin", Icon: "icon-home", Type: 1, Pid: 0, Sort: 0, Path: "/dashboard", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 2, Name: "主页", GuardName: "admin", Icon: "", Type: 2, Pid: 1, Sort: 0, Path: "/api/admin/dashboard/index/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 3, Name: "管理员", GuardName: "admin", Icon: "icon-admin", Type: 1, Pid: 0, Sort: 100, Path: "/admin", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 4, Name: "管理员列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Sort: 0, Path: "/api/admin/admin/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 5, Name: "权限列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Sort: 0, Path: "/api/admin/permission/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 6, Name: "角色列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Sort: 0, Path: "/api/admin/role/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 7, Name: "系统配置", GuardName: "admin", Icon: "icon-setting", Type: 1, Pid: 0, Sort: 100, Path: "/system", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 8, Name: "设置管理", GuardName: "admin", Icon: "", Type: 1, Pid: 7, Sort: 0, Path: "/system/config", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 9, Name: "网站设置", GuardName: "admin", Icon: "", Type: 2, Pid: 8, Sort: 0, Path: "/api/admin/webConfig/setting/form", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 10, Name: "配置管理", GuardName: "admin", Icon: "", Type: 2, Pid: 8, Sort: 0, Path: "/api/admin/config/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 11, Name: "菜单管理", GuardName: "admin", Icon: "", Type: 2, Pid: 7, Sort: 0, Path: "/api/admin/menu/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 12, Name: "操作日志", GuardName: "admin", Icon: "", Type: 2, Pid: 7, Sort: 100, Path: "/api/admin/actionLog/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 13, Name: "附件空间", GuardName: "admin", Icon: "icon-attachment", Type: 1, Pid: 0, Sort: 100, Path: "/attachment", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 14, Name: "文件管理", GuardName: "admin", Icon: "", Type: 2, Pid: 13, Sort: 0, Path: "/api/admin/file/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 15, Name: "图片管理", GuardName: "admin", Icon: "", Type: 2, Pid: 13, Sort: 0, Path: "/api/admin/picture/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 16, Name: "我的账号", GuardName: "admin", Icon: "icon-user", Type: 1, Pid: 0, Sort: 100, Path: "/account", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 17, Name: "个人设置", GuardName: "admin", Icon: "", Type: 2, Pid: 16, Sort: 0, Path: "/api/admin/account/setting/form", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}

	db.Client.Create(&seeders)
}

// 获取TreeSelect组件数据
func (model *Menu) TreeSelect(root bool) (list []*treeselect.TreeData, Error error) {

	// 是否有根节点
	if root {
		list = append(list, &treeselect.TreeData{
			Title: "根节点",
			Value: 0,
		})
	}

	list = append(list, model.FindTreeSelectNode(0)...)

	return list, nil
}

// 递归获取TreeSelect组件数据
func (model *Menu) FindTreeSelectNode(pid int) (list []*treeselect.TreeData) {
	menus := []Menu{}
	db.Client.
		Where("guard_name = ?", "admin").
		Where("pid = ?", pid).
		Order("sort asc,id asc").
		Select("name", "id", "pid").
		Find(&menus)

	if len(menus) == 0 {
		return list
	}

	for _, v := range menus {
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

// 获取Tree组件数据
func (model *Menu) Tree() (list []*tree.TreeData, Error error) {
	list = append(list, model.FindTreeNode(0)...)

	return list, nil
}

// 递归获取Tree组件数据
func (model *Menu) FindTreeNode(pid int) (list []*tree.TreeData) {
	menus := []Menu{}
	db.Client.
		Where("guard_name = ?", "admin").
		Where("pid = ?", pid).
		Order("sort asc").
		Select("name", "id", "pid").
		Find(&menus)

	if len(menus) == 0 {
		return list
	}

	for _, v := range menus {
		item := &tree.TreeData{
			Key:   v.Id,
			Title: v.Name,
		}

		children := model.FindTreeNode(v.Id)
		if len(children) > 0 {
			item.Children = children
		}

		list = append(list, item)
	}

	return list
}

// 递归获取父数据
func (model *Menu) FindParentTreeNode(chrildPid int) (list []*Menu) {
	menus := []*Menu{}
	db.Client.
		Where("guard_name = ?", "admin").
		Where("id = ?", chrildPid).
		Where("type IN ?", []int{1, 2, 3}).
		Find(&menus)

	if len(menus) == 0 {
		return menus
	}

	for _, v := range menus {
		if v.Pid != 0 {
			children := model.FindParentTreeNode(v.Pid)
			if len(children) > 0 {
				menus = append(menus, children...)
			}
		}
	}

	return menus
}

// 通过管理员ID权限菜单
func (model *Menu) GetListByAdminId(adminId int) (menuList interface{}, err error) {
	menus := []*Menu{}

	if adminId == 1 {
		db.Client.
			Where("status = ?", 1).
			Where("guard_name", "admin").
			Where("type IN ?", []int{1, 2, 3}).
			Order("sort asc").
			Find(&menus)

		return model.MenuParser(menus)
	}

	var menuIds []int
	roleHasMenus, err := (&CasbinRule{}).GetUserMenus(adminId)
	if err != nil {
		return menuList, err
	}
	if len(roleHasMenus) == 0 {
		return
	}

	for _, v := range roleHasMenus {
		menuIds = append(menuIds, v.Id)
	}

	// 最底层列表
	db.Client.
		Where("status = ?", 1).
		Where("guard_name = ?", "admin").
		Where("id in ?", menuIds).
		Where("type IN ?", []int{1, 2, 3}).
		Where("pid <> ?", 0).
		Find(&menus)

	for _, v := range menus {
		list := model.FindParentTreeNode(v.Pid)
		for _, v := range list {
			menuIds = append(menuIds, v.Id)
		}
	}

	// 所有列表
	db.Client.
		Where("guard_name = ?", "admin").
		Where("id in ?", menuIds).
		Order("sort asc").
		Find(&menus)

	return model.MenuParser(menus)
}

// 解析菜单
func (model *Menu) MenuParser(menus []*Menu) (menuList interface{}, Error error) {
	newMenus := []*Menu{}

	for _, v := range menus {
		v.Key = uuid.New()
		v.Locale = "menu" + strings.Replace(v.Path, "/", ".", -1)

		if v.Show == 1 {
			v.HideInMenu = false
		} else {
			v.HideInMenu = true
		}

		if v.Type == 2 && v.IsEngine == 1 {
			v.Path = "/layout/index?api=" + v.Path
		}

		if !model.HasMenu(newMenus, v.Id) && v.Type != 3 {
			newMenus = append(newMenus, v)
		}
	}

	return lister.ListToTree(newMenus, "id", "pid", "routes", 0)
}

// 判断菜单是否已经存在
func (model *Menu) HasMenu(menus []*Menu, id int) (result bool) {
	for _, v := range menus {
		if v.Id == id {
			result = true
		}
	}

	return
}

// 通过ID获取菜单信息
func (model *Menu) GetInfoById(id interface{}) (menu *Menu, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&menu).Error

	return menu, err
}

// 通过名称获取菜单信息
func (model *Menu) GetInfoByName(name string) (menu *Menu, Error error) {
	err := db.Client.Where("status = ?", 1).Where("name = ?", name).First(&menu).Error

	return menu, err
}

// 通过ID判断菜单是否已存在
func (model *Menu) IsExist(id interface{}) bool {
	menu := Menu{}
	err := db.Client.Where("id = ?", id).First(&menu).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	if err != nil {
		panic(err)
	}

	return true
}

// 通过id集合获取列表
func (model *Menu) GetListByIds(menuIds interface{}) (menus []*Menu, Error error) {
	err := db.Client.Where("id in ?", menuIds).Find(&menus).Error

	return menus, err
}
