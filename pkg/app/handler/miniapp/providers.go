package miniapp

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/miniapp/forms"
	"github.com/quarkcms/quark-go/pkg/app/handler/miniapp/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
	&forms.Index{},
}
