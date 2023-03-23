package tool

import "github.com/quarkcms/quark-go/pkg/app/handler/tool/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
