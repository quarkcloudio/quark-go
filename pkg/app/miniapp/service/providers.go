package service

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/service/forms"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/service/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
	&forms.Index{},
}
