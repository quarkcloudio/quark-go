package miniapp

import "github.com/quarkcms/quark-go/pkg/app/handler/miniapp/pages"

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
}
