package service

import "github.com/quarkcloudio/quark-go/v3/pkg/app/tool/service/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
