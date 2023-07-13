package service

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/dashboards"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/layouts"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/logins"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/resources"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/uploads"
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
