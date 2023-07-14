package ginadapter

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 适配gin框架路由
func RouteAdapter(b *builder.Engine, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(200, builder.Error(err.Error()))
		return
	}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(200, builder.Error(err.Error()))
		return
	}

	//把读过的字节流重新放到body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	// 转换Request对象
	context := b.TransformContext(
		ctx.FullPath(),
		ctx.Request.Header,
		ctx.Request.Method,
		ctx.Request.URL.String(),
		bytes.NewReader(body),
		ctx.Writer,
	)

	b.Render(context)
}

// 适配gin框架
func Adapter(b *builder.Engine, app *gin.Engine) {

	// 获取注册的服务
	routePaths := b.GetRoutePaths()

	// 解析服务
	for _, v := range routePaths {
		switch v.Method {
		case "GET":
			app.GET(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "HEAD":
			app.HEAD(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "OPTIONS":
			app.OPTIONS(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "POST":
			app.POST(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "PUT":
			app.PUT(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "PATCH":
			app.PATCH(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "DELETE":
			app.DELETE(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		case "Any":
			app.Any(v.Path, func(ctx *gin.Context) {
				RouteAdapter(b, ctx)
			})
		}
	}
}
