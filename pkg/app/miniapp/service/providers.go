package service

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/miniapp/service/forms"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/miniapp/service/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
	&forms.Demo{},
}
