package model

import (
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/utils/datetime"
)

// 字段
type Menu struct {
	Id         int               `json:"id" gorm:"autoIncrement"`
	Name       string            `json:"name" gorm:"size:100;not null"`
	GuardName  string            `json:"group_name" gorm:"size:100;not null"`
	Icon       string            `json:"icon" gorm:"size:100;"`
	Type       int               `json:"type" gorm:"size:100;not null"` // 菜单类型：1目录，2菜单，3按钮
	Pid        int               `json:"pid" gorm:"size:11;default:0"`
	Order      int               `json:"order" gorm:"size:11;default:0"`
	Path       string            `json:"path" gorm:"size:255"`
	Query      string            `json:"query" gorm:"size:255"`
	Component  string            `json:"component" gorm:"size:255"`
	Show       int               `json:"show" gorm:"size:1;not null;default:1"`
	IsEngine   int               `json:"is_engine" gorm:"size:1;not null;default:0"`
	IsLink     int               `json:"is_link" gorm:"size:1;not null;default:0"`
	Status     int               `json:"status" gorm:"size:1;not null;default:1"`
	Key        string            `json:"key" gorm:"<-:false"`
	Locale     string            `json:"locale" gorm:"<-:false"`
	HideInMenu bool              `json:"hide_in_menu" gorm:"<-:false"`
	CreatedAt  datetime.Datetime `json:"created_at"`
	UpdatedAt  datetime.Datetime `json:"updated_at"`
}

// 菜单表
func (p *Menu) Seeder() {
	seeders := []Menu{
		{Id: 1, Name: "控制台", GuardName: "admin", Icon: "icon-home", Type: 1, Pid: 0, Order: 0, Path: "/dashboard", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 2, Name: "主页", GuardName: "admin", Icon: "", Type: 2, Pid: 1, Order: 0, Path: "/api/admin/dashboard/index/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 3, Name: "用户管理", GuardName: "admin", Icon: "icon-admin", Type: 1, Pid: 0, Order: 100, Path: "/user", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 4, Name: "用户列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Order: 0, Path: "/api/admin/user/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 5, Name: "权限列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Order: 0, Path: "/api/admin/permission/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 6, Name: "角色列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Order: 0, Path: "/api/admin/role/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 7, Name: "系统配置", GuardName: "admin", Icon: "icon-setting", Type: 1, Pid: 0, Order: 100, Path: "/system", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 8, Name: "设置管理", GuardName: "admin", Icon: "", Type: 1, Pid: 7, Order: 0, Path: "/system/config", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 9, Name: "网站设置", GuardName: "admin", Icon: "", Type: 2, Pid: 8, Order: 0, Path: "/api/admin/webConfig/form", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 10, Name: "配置管理", GuardName: "admin", Icon: "", Type: 2, Pid: 8, Order: 0, Path: "/api/admin/config/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 11, Name: "菜单管理", GuardName: "admin", Icon: "", Type: 2, Pid: 7, Order: 0, Path: "/api/admin/menu/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 12, Name: "操作日志", GuardName: "admin", Icon: "", Type: 2, Pid: 7, Order: 100, Path: "/api/admin/actionLog/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 13, Name: "附件空间", GuardName: "admin", Icon: "icon-attachment", Type: 1, Pid: 0, Order: 100, Path: "/attachment", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 14, Name: "文件管理", GuardName: "admin", Icon: "", Type: 2, Pid: 13, Order: 0, Path: "/api/admin/file/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 15, Name: "图片管理", GuardName: "admin", Icon: "", Type: 2, Pid: 13, Order: 0, Path: "/api/admin/image/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 16, Name: "我的账号", GuardName: "admin", Icon: "icon-user", Type: 1, Pid: 0, Order: 100, Path: "/account", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 17, Name: "个人设置", GuardName: "admin", Icon: "", Type: 2, Pid: 16, Order: 0, Path: "/api/admin/account/form", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 18, Name: "部门列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Order: 0, Path: "/api/admin/department/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 19, Name: "职位列表", GuardName: "admin", Icon: "", Type: 2, Pid: 3, Order: 0, Path: "/api/admin/position/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}

	db.Client.Create(&seeders)
}
