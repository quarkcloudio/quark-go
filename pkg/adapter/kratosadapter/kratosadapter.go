package kratosadapter

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

// 适配kratos框架路由
func RouteAdapter(b *builder.Engine, routePath string) http.HandlerFunc {
	return func(ctx http.Context) error {

		// 转换Request对象
		context := b.NewContext(ctx.Response(), ctx.Request())

		// 设置路由
		context.SetFullPath(routePath)

		return b.Render(context)
	}
}

// 适配kratos框架
func Adapter(b *builder.Engine, s *http.Server) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 路由组列表
	r := s.Route("/")

	// 解析服务
	for _, v := range routePaths {

		// 分割路由路径
		paths := strings.Split(v.Path, "/")

		// 将 "/helloworld/:name" 转换成 "/helloworld/{name}" 格式路由
		path := ""
		for _, sv := range paths {
			if strings.Contains(sv, ":") {
				sv = "{" + strings.Trim(sv, ":") + "}"
			}

			path = path + "/" + sv
		}

		if v.Method == "Any" {
			r.Handle(stdhttp.MethodGet, path, RouteAdapter(b, v.Path))
			r.Handle(stdhttp.MethodPost, path, RouteAdapter(b, v.Path))
			r.Handle(stdhttp.MethodDelete, path, RouteAdapter(b, v.Path))
			r.Handle(stdhttp.MethodPut, path, RouteAdapter(b, v.Path))
		} else {
			r.Handle(v.Method, path, RouteAdapter(b, v.Path))
		}
	}
}
