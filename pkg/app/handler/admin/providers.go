package admin

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/dashboards"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/login"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&login.Index{},
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
}
