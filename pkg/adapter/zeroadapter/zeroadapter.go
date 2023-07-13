package zeroadapter

import (
	"net/http"

	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/zeromicro/go-zero/rest"
)

// 适配gozero框架路由
func RouteAdapter(b *builder.Engine, routePath string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// 转换Request对象
		context := b.NewContext(w, r)

		// 设置路由
		context.SetFullPath(routePath)

		b.Render(context)
	}
}

// 适配gozero框架
func Adapter(b *builder.Engine, server *rest.Server) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 路由组列表
	routes := []rest.Route{}

	// 解析服务
	for _, v := range routePaths {
		if v.Method == "Any" {
			routes = append(routes, rest.Route{
				Method:  http.MethodGet,
				Path:    v.Path,
				Handler: RouteAdapter(b, v.Path),
			})
			routes = append(routes, rest.Route{
				Method:  http.MethodPost,
				Path:    v.Path,
				Handler: RouteAdapter(b, v.Path),
			})
			routes = append(routes, rest.Route{
				Method:  http.MethodDelete,
				Path:    v.Path,
				Handler: RouteAdapter(b, v.Path),
			})
			routes = append(routes, rest.Route{
				Method:  http.MethodPut,
				Path:    v.Path,
				Handler: RouteAdapter(b, v.Path),
			})
		} else {
			routes = append(routes, rest.Route{
				Method:  v.Method,
				Path:    v.Path,
				Handler: RouteAdapter(b, v.Path),
			})
		}
	}

	server.AddRoutes(routes)
}
