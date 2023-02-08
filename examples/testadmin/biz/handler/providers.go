package handler

import (
	"github.com/quarkcms/quark-go/examples/hertzadmin/biz/handler/resources"
)

// 注册服务
var Providers = []interface{}{
	&resources.Demo{},
}
