package admin

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/dashboards"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/layouts"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/logins"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/resources"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/uploads"
)

// 注册服务
var Providers = []interface{}{
	&logins.Index{},
	&layouts.Index{},
	&dashboards.Index{},
	&resources.Admin{},
	&resources.Role{},
	&resources.Permission{},
	&resources.Menu{},
	&resources.ActionLog{},
	&resources.Config{},
	&resources.File{},
	&resources.Picture{},
	&resources.WebConfig{},
	&resources.Account{},
	&uploads.File{},
	&uploads.Image{},
}
