package app

import (
	"github.com/quarkcloudio/quark-go/v4/app/dashboards"
	"github.com/quarkcloudio/quark-go/v4/app/layouts"
	"github.com/quarkcloudio/quark-go/v4/app/logins"
	"github.com/quarkcloudio/quark-go/v4/app/pages"
	"github.com/quarkcloudio/quark-go/v4/app/resources"
	"github.com/quarkcloudio/quark-go/v4/app/uploads"
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
