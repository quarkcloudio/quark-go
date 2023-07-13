package fiberadapter

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 适配gofiber框架路由
func RouteAdapter(b *builder.Engine, ctx *fiber.Ctx) error {

	header := make(map[string][]string)
	ctx.Request().Header.VisitAll(func(k, v []byte) {
		header[string(k)] = append(header[string(k)], string(v))
	})

	// 转换Request对象
	context := b.TransformContext(
		ctx.Route().Path,
		header,
		string(ctx.Method()),
		ctx.Request().URI().String(),
		bytes.NewReader(ctx.Body()),
		ctx.Response().BodyWriter(),
	)

	return b.Render(context)
}

// 适配gofiber框架
func Adapter(b *builder.Engine, app *fiber.App) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 解析服务
	for _, v := range routePaths {
		switch v.Method {
		case "GET":
			app.Get(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "HEAD":
			app.Head(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "OPTIONS":
			app.Options(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "POST":
			app.Post(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "PUT":
			app.Put(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "PATCH":
			app.Patch(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "DELETE":
			app.Delete(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		case "Any":
			app.All(v.Path, func(ctx *fiber.Ctx) error {
				return RouteAdapter(b, ctx)
			})
		}
	}
}
