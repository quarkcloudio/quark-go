package hertzadapter

import (
	"bytes"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

const JSON_RESPONSE = "json"   // json类型响应
const IMAGE_RESPONSE = "image" // 图片类型响应
const EXCEL_RESPONSE = "excel" // Excel文件类型响应

// 适配hertz框架路由
func RouteAdapter(b *builder.Engine, responseType string, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	// 转换Request对象
	result, err := b.Transform(
		string(ctx.Request.Method()),
		ctx.URI().String(),
		bytes.NewReader(body),
		ctx.Response.BodyWriter(),
	).Render()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	// 响应结果
	switch responseType {
	case JSON_RESPONSE:
		ctx.JSON(200, result)
	case IMAGE_RESPONSE:
		ctx.Write(result.([]byte))
	case EXCEL_RESPONSE:
		ctx.Response.Header.Set("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		ctx.Response.Header.Set("Content-Type", "application/octet-stream")
		ctx.Write(result.([]byte))
	}
}

// 适配hertz框架
func Adapter(b *builder.Engine, r *server.Hertz) {

	// 后台路由组
	rg := r.Group("/api/admin")

	// 登录
	rg.GET("/login/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/login/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/login/:resource/captchaId", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/login/:resource/captcha/:id", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, IMAGE_RESPONSE, ctx)
	})
	rg.GET("/logout/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 仪表盘
	rg.GET("/dashboard/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 增删改查
	rg.GET("/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/editable", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Any("/:resource/action/:uriKey", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/create", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/:resource/store", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/edit", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/edit/values", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/:resource/save", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/detail", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/export", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.Any("/:resource/import", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/import/template", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.GET("/:resource/:uriKey/form", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 文件上传
	rg.POST("/upload/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/upload/:resource/getList", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/upload/:resource/delete", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/upload/:resource/crop", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
}
