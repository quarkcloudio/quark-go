package zeroadapter

import (
	"net/http"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 适配gozero框架路由
func RouteAdapter(b *builder.Engine, routePath string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// 转换Request对象
		context := b.NewContext(w, r)

		// 设置路由
		context.SetFullPath(routePath)

		err := b.Render(context)
		if err != nil {
			httpx.OkJson(w, msg.Error(err.Error(), ""))
			return
		}
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
		routes = append(routes, rest.Route{
			Method:  v.Method,
			Path:    v.Path,
			Handler: RouteAdapter(b, v.Path),
		})
	}

	server.AddRoutes(routes)
}
