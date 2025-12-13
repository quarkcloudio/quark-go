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
	Permission string            `json:"permission" gorm:"size:100"`
	Icon       string            `json:"icon" gorm:"size:100;"`
	Type       int               `json:"type" gorm:"size:100;not null"` // 菜单类型：1目录，2菜单，3按钮
	Pid        int               `json:"pid" gorm:"size:11;default:0"`
	Sort       int               `json:"sort" gorm:"size:11;default:0"`
	Path       string            `json:"path" gorm:"size:255"`
	Query      string            `json:"query" gorm:"size:255"`
	Component  string            `json:"component" gorm:"size:255"`
	Visible    int               `json:"visible" gorm:"size:1;not null;default:1"`
	IsEngine   int               `json:"is_engine" gorm:"size:1;not null;default:0"`
	IsLink     int               `json:"is_link" gorm:"size:1;not null;default:0"`
	IsFrame    int               `json:"is_frame" gorm:"size:1;not null;default:0"`
	Status     int               `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt  datetime.Datetime `json:"created_at"`
	UpdatedAt  datetime.Datetime `json:"updated_at"`
}

// 菜单表
func (p *Menu) Seeder() {
	seeders := []Menu{
		{Id: 1, Name: "首页", GuardName: "admin", Icon: "ant-design:home-outlined", Type: 1, Pid: 0, Sort: 0, Path: "home", Query: "{\"api\":\"/api/admin/dashboard/index/index\"}", Component: "home/index", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 3, Name: "用户管理", GuardName: "admin", Icon: "ant-design:usergroup-add-outlined", Type: 1, Pid: 0, Sort: 100, Path: "user", Query: "", Component: "", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 4, Name: "用户列表", GuardName: "admin", Icon: "ant-design:user-add-outlined", Type: 2, Pid: 3, Sort: 0, Path: "user", Query: "{\"api\":\"/api/admin/user/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 5, Name: "权限列表", GuardName: "admin", Icon: "ant-design:profile-outlined", Type: 2, Pid: 3, Sort: 0, Path: "permission", Query: "{\"api\":\"/api/admin/permission/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 6, Name: "角色列表", GuardName: "admin", Icon: "ant-design:idcard-outlined", Type: 2, Pid: 3, Sort: 0, Path: "role", Query: "{\"api\":\"/api/admin/role/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 7, Name: "系统配置", GuardName: "admin", Icon: "ant-design:setting-outlined", Type: 1, Pid: 0, Sort: 100, Path: "system", Query: "", Component: "", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 8, Name: "设置管理", GuardName: "admin", Icon: "ant-design:appstore-add-outlined", Type: 1, Pid: 7, Sort: 0, Path: "config", Query: "", Component: "", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 9, Name: "网站设置", GuardName: "admin", Icon: "ant-design:cluster-outlined", Type: 2, Pid: 8, Sort: 0, Path: "webConfig", Query: "{\"api\":\"/api/admin/webConfig/form\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 10, Name: "配置管理", GuardName: "admin", Icon: "ant-design:tool-outlined", Type: 2, Pid: 8, Sort: 0, Path: "config", Query: "{\"api\":\"/api/admin/config/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 11, Name: "菜单管理", GuardName: "admin", Icon: "ant-design:menu-outlined", Type: 2, Pid: 7, Sort: 0, Path: "menu", Query: "{\"api\":\"/api/admin/menu/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 12, Name: "操作日志", GuardName: "admin", Icon: "ant-design:file-done-outlined", Type: 2, Pid: 7, Sort: 100, Path: "actionLog", Query: "{\"api\":\"/api/admin/actionLog/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 13, Name: "附件空间", GuardName: "admin", Icon: "ant-design:folder-outlined", Type: 1, Pid: 0, Sort: 100, Path: "attachment", Query: "", Component: "", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 14, Name: "文件管理", GuardName: "admin", Icon: "ant-design:file-outlined", Type: 2, Pid: 13, Sort: 0, Path: "file", Query: "{\"api\":\"/api/admin/file/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 15, Name: "图片管理", GuardName: "admin", Icon: "ant-design:picture-outlined", Type: 2, Pid: 13, Sort: 0, Path: "image", Query: "{\"api\":\"/api/admin/image/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 16, Name: "我的账号", GuardName: "admin", Icon: "ant-design:user-outlined", Type: 1, Pid: 0, Sort: 110, Path: "account", Query: "", Component: "", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 17, Name: "个人设置", GuardName: "admin", Icon: "ant-design:user-switch-outlined", Type: 2, Pid: 16, Sort: 0, Path: "setting", Query: "{\"api\":\"/api/admin/account/form\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 18, Name: "部门列表", GuardName: "admin", Icon: "ant-design:apartment-outlined", Type: 2, Pid: 3, Sort: 0, Path: "department", Query: "{\"api\":\"/api/admin/department/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 19, Name: "职位列表", GuardName: "admin", Icon: "ant-design:bars-outlined", Type: 2, Pid: 3, Sort: 0, Path: "position", Query: "{\"api\":\"/api/admin/position/index\"}", Component: "", Visible: 1, IsEngine: 1, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 20, Name: "组件调试", GuardName: "admin", Icon: "ant-design:appstore-outlined", Type: 1, Pid: 0, Sort: 100, Path: "develop", Query: "", Component: "", Visible: 0, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
		{Id: 21, Name: "组件开发", GuardName: "admin", Icon: "ant-design:experiment-outlined", Type: 2, Pid: 20, Sort: 0, Path: "index", Query: "", Component: "develop/index", Visible: 1, IsEngine: 0, IsLink: 0, IsFrame: 0, Status: 1, Permission: ""},
	}

	db.Client.Create(&seeders)
}
