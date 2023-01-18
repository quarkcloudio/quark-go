package fiberadapter

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

const COMPONENT_RESPONSE = "component" // 组件类型响应
const ACTION_RESPONSE = "action"       // 行为类型响应
const FILE_RESPONSE = "file"           // 文件类型响应

// 将gofiber框架的Ctx转换为builder框架Request
func RequestAdapter(ctx *fiber.Ctx) (*builder.Request, error) {

	return &builder.Request{
		IPString:       ctx.IP(),
		HeaderString:   ctx.Request().Header.String(),
		MethodString:   string(ctx.Method()),
		FullPathString: ctx.Route().Path,
		HostString:     string(ctx.Hostname()),
		PathString:     string(ctx.Path()),
		QueryString:    string(ctx.Context().QueryArgs().QueryString()),
		BodyBuffer:     ctx.Body(),
	}, nil
}

// 适配gofiber框架响应
func ResponseAdapter(r *builder.Resource, responseType string, ctx *fiber.Ctx) error {
	var responseError error

	result, err := r.Run()
	if err != nil {
		return ctx.JSON(msg.Error(err.Error(), ""))
	}

	switch responseType {
	case "component":
		return ctx.JSON(result)
	case "action":
		return ctx.JSON(msg.Success("操作成功", "", result))
	case "file":
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	}

	return responseError
}

// 适配gofiber框架路由
func RouteAdapter(b *builder.Resource, responseType string, ctx *fiber.Ctx) error {

	// 适配请求
	request, err := RequestAdapter(ctx)
	if err != nil {
		return ctx.JSON(msg.Error(err.Error(), ""))
	}

	// 适配响应
	resource := b.TransformRequest(request)
	return ResponseAdapter(resource, responseType, ctx)
}

// 适配gofiber框架
func Adapter(b *builder.Resource, app *fiber.App) {

	// 后台路由组
	rg := app.Group("/api/admin")

	// 登录
	rg.Get("/login/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, COMPONENT_RESPONSE, ctx)
	})
	rg.Post("/login/:resource/handle", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, ACTION_RESPONSE, ctx)
	})
	rg.Get("/login/:resource/captchaId", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, ACTION_RESPONSE, ctx)
	})
	rg.Get("/login/:resource/captcha/:id", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, FILE_RESPONSE, ctx)
	})

	// 仪表盘
	rg.Get("/dashboard/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, COMPONENT_RESPONSE, ctx)
	})

	// 增删改查
	rg.Get("/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, COMPONENT_RESPONSE, ctx)
	})
}
