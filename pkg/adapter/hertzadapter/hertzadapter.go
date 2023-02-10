package hertzadapter

import (
	"bytes"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 适配hertz框架路由
func RouteAdapter(b *builder.Engine, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
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

	err = b.Render(context)
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}
}

// 适配hertz框架
func Adapter(b *builder.Engine, r *server.Hertz) {

	// 后台路由组
	rg := r.Group("/api/admin")

	// 登录
	rg.GET("/login/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.POST("/login/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/login/:resource/captchaId", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/login/:resource/captcha/:id", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/logout/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})

	// 仪表盘
	rg.GET("/dashboard/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})

	// 增删改查
	rg.GET("/:resource/index", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/editable", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.Any("/:resource/action/:uriKey", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/create", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.POST("/:resource/store", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/edit", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/edit/values", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.POST("/:resource/save", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/detail", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/export", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.Any("/:resource/import", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/import/template", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/:resource/:uriKey/form", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})

	// 文件上传
	rg.POST("/upload/:resource/handle", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.GET("/upload/:resource/getList", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.POST("/upload/:resource/delete", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
	rg.POST("/upload/:resource/crop", func(c context.Context, ctx *app.RequestContext) {
		RouteAdapter(b, ctx)
	})
}
