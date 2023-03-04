package mix

import "github.com/quarkcms/quark-go/pkg/app/handler/mix/pages"

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
}
