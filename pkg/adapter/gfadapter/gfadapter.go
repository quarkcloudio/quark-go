package gfadapter

import (
	"bytes"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

// 适配goframe框架路由
func routeAdapter(b *builder.Engine, r *ghttp.Request) error {

	header := r.Header.Clone()

	// 转换Request对象
	context := b.TransformContext(
		r.Router.Uri,
		header,
		r.Method,
		r.URL.RequestURI(),
		bytes.NewReader(r.GetBody()),
		r.Response.Writer,
	)
	err := b.Render(context)

	return err
}

// 适配goframe框架
func Adapter(b *builder.Engine, server *ghttp.Server) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 解析服务
	for _, v := range routePaths {
		server.BindHandler(v.Path, func(r *ghttp.Request) {
			routeAdapter(b, r)
		})
	}
	return
}
