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

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/index"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/login/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/handle"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/captchaId",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/captchaId"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/captcha/:id",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/captcha/:id"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/logout/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/logout/:resource/handle"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/dashboard/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/dashboard/:resource/index"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/:resource/index"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/editable",
				Handler: RouteAdapter(b, "/api/admin/:resource/editable"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/action/:uriKey",
				Handler: RouteAdapter(b, "/api/admin/:resource/action/:uriKey"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/action/:uriKey",
				Handler: RouteAdapter(b, "/api/admin/:resource/action/:uriKey"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/create",
				Handler: RouteAdapter(b, "/api/admin/:resource/create"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/store",
				Handler: RouteAdapter(b, "/api/admin/:resource/store"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/edit",
				Handler: RouteAdapter(b, "/api/admin/:resource/edit"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/edit/values",
				Handler: RouteAdapter(b, "/api/admin/:resource/edit/values"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/save",
				Handler: RouteAdapter(b, "/api/admin/:resource/save"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/detail",
				Handler: RouteAdapter(b, "/api/admin/:resource/detail"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/export",
				Handler: RouteAdapter(b, "/api/admin/:resource/export"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/import",
				Handler: RouteAdapter(b, "/api/admin/:resource/import"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/import",
				Handler: RouteAdapter(b, "/api/admin/:resource/import"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/import/template",
				Handler: RouteAdapter(b, "/api/admin/:resource/import/template"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/:uriKey/form",
				Handler: RouteAdapter(b, "/api/admin/:resource/:uriKey/form"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/handle"),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/upload/:resource/getList",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/getList"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/delete",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/delete"),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/crop",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/crop"),
			},
		},
	)
}
