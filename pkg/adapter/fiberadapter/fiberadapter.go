package fiberadapter

import (
	"bytes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

const JSON_RESPONSE = "json"   // json类型响应
const IMAGE_RESPONSE = "image" // 图片类型响应
const EXCEL_RESPONSE = "excel" // Excel文件类型响应

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
	case JSON_RESPONSE:
		return ctx.JSON(result)
	case IMAGE_RESPONSE:
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	case EXCEL_RESPONSE:
		ctx.Set("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		ctx.Set("Content-Type", "application/octet-stream")
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	}

	return responseError
}

// 适配gofiber框架路由
func RouteAdapter(b *builder.Resource, responseType string, ctx *fiber.Ctx) error {
	var responseError error

	// 将框架请求转换为builder框架请求
	request := &builder.Request{
		IPString:       ctx.IP(),
		HeaderString:   ctx.Request().Header.String(),
		MethodString:   string(ctx.Method()),
		FullPathString: ctx.Route().Path,
		HostString:     string(ctx.Hostname()),
		PathString:     string(ctx.Path()),
		QueryString:    string(ctx.Context().QueryArgs().QueryString()),
		BodyBuffer:     ctx.Body(),
	}

	// 转换Request对象
	result, err := b.TransformRequest(request).Run()
	if err != nil {
		return ctx.JSON(msg.Error(err.Error(), ""))
	}

	// 响应结果
	switch responseType {
	case JSON_RESPONSE:
		return ctx.JSON(result)
	case IMAGE_RESPONSE:
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	case EXCEL_RESPONSE:
		ctx.Set("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		ctx.Set("Content-Type", "application/octet-stream")
		return ctx.SendStream(bytes.NewReader(result.([]byte)))
	}

	return responseError
}

// 适配gofiber框架
func Adapter(b *builder.Resource, app *fiber.App) {

	// 后台路由组
	rg := app.Group("/api/admin")

	// 登录
	rg.Get("/login/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Post("/login/:resource/handle", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/login/:resource/captchaId", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/login/:resource/captcha/:id", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, IMAGE_RESPONSE, ctx)
	})
	rg.Get("/logout/:resource/handle", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 仪表盘
	rg.Get("/dashboard/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 增删改查
	rg.Get("/:resource/index", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/editable", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.All("/:resource/action/:uriKey", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/create", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Post("/:resource/store", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/edit", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/edit/values", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Post("/:resource/save", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/detail", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/export", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.All("/:resource/import", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Get("/:resource/import/template", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.Get("/:resource/:uriKey/form", func(ctx *fiber.Ctx) error {
		return RouteAdapter(b, JSON_RESPONSE, ctx)
	})

}
