package service

import "github.com/quarkcms/quark-go/v2/pkg/app/tool/service/upload"

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
}
