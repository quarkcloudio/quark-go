package hertzadapter

import (
	"bytes"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 适配hertz框架路由
func RouteAdapter(b *builder.Engine, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(200, builder.Error(err.Error()))
		return
	}

	header := make(map[string][]string)
	ctx.Request.Header.VisitAll(func(k, v []byte) {
		header[string(k)] = append(header[string(k)], string(v))
	})

	// 转换Request对象
	context := b.TransformContext(
		ctx.FullPath(),
		header,
		string(ctx.Request.Method()),
		ctx.URI().String(),
		bytes.NewReader(body),
		ctx.Response.BodyWriter(),
	)

	b.Render(context)
}

// 适配hertz框架
func Adapter(b *builder.Engine, r *server.Hertz) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 解析服务
	for _, v := range routePaths {
		switch v.Method {
		case "GET":
			r.GET(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "HEAD":
			r.HEAD(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "OPTIONS":
			r.OPTIONS(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "POST":
			r.POST(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "PUT":
			r.PUT(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "PATCH":
			r.PATCH(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "DELETE":
			r.DELETE(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		case "Any":
			r.Any(v.Path, func(c context.Context, ctx *app.RequestContext) {
				RouteAdapter(b, ctx)
			})
		}
	}
}
