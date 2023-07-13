package service

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/service/pages"

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
}
