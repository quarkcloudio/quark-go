package handler

import (
	"github.com/quarkcloudio/quark-go/v2/examples/hertzadmin/biz/handler/resources"
)

// 注册服务
var Providers = []interface{}{
	&resources.Demo{},
}
