package service

import (
	"github.com/quarkcloudio/quark-go/v3/app/tool/captcha"
	"github.com/quarkcloudio/quark-go/v3/app/tool/upload"
)

// 注册服务
var Providers = []interface{}{
	&upload.File{},
	&upload.Image{},
	&captcha.Index{},
}
