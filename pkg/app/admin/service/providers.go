package service

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/dashboards"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/layouts"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/logins"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/resources"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/uploads"
)

// 注册服务
var Providers = []interface{}{
	&logins.Index{},
	&layouts.Index{},
	&dashboards.Index{},
	&resources.User{},
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
