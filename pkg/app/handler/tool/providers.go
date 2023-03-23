package tool

import "github.com/quarkcms/quark-go/pkg/app/handler/tool/uploads"

// 注册服务
var Providers = []interface{}{
	&uploads.File{},
	&uploads.Image{},
}
