package admin

import (
	"github.com/quarkcloudio/quark-go/v3/app/admin/dashboards"
	"github.com/quarkcloudio/quark-go/v3/app/admin/layouts"
	"github.com/quarkcloudio/quark-go/v3/app/admin/logins"
	"github.com/quarkcloudio/quark-go/v3/app/admin/pages"
	"github.com/quarkcloudio/quark-go/v3/app/admin/resources"
	"github.com/quarkcloudio/quark-go/v3/app/admin/uploads"
)

// 注册服务
var Providers = []interface{}{
	&logins.Index{},
	&layouts.Index{},
	&dashboards.Index{},
	&resources.User{},
	&resources.Role{},
	&resources.Permission{},
	&resources.Department{},
	&resources.Position{},
	&resources.Menu{},
	&resources.ActionLog{},
	&resources.Config{},
	&resources.File{},
	&resources.Image{},
	&resources.WebConfig{},
	&resources.Account{},
	&pages.Index{},
	&uploads.File{},
	&uploads.Image{},
}
