package ginadapter

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

const JSON_RESPONSE = "json"   // json类型响应
const IMAGE_RESPONSE = "image" // 图片类型响应
const EXCEL_RESPONSE = "excel" // Excel文件类型响应

// 将gin框架的Ctx转换为builder框架Request
func RequestAdapter(ctx *gin.Context) (*builder.Request, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	data, err := ctx.GetRawData()
	if err != nil {
		return nil, err
	}
	//把读过的字节流重新放到body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	// 将框架请求转换为builder框架请求
	return &builder.Request{
		IPString:       ctx.ClientIP(),
		HeaderString:   ctx.GetHeader(""),
		MethodString:   ctx.Request.Method,
		FullPathString: ctx.FullPath(),
		HostString:     ctx.Request.Host,
		PathString:     ctx.Request.URL.Path,
		QueryString:    ctx.Request.Response.Request.URL.RawQuery,
		BodyBuffer:     body,
	}, nil
}

// 适配gin框架响应
func ResponseAdapter(r *builder.Resource, responseType string, ctx *gin.Context) {
	result, err := r.Run()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	// 响应结果
	switch responseType {
	case JSON_RESPONSE:
		ctx.JSON(200, result)
		return
	case IMAGE_RESPONSE:
		ctx.Writer.Write(result.([]byte))
		return
	case EXCEL_RESPONSE:
		ctx.Header("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Writer.Write(result.([]byte))
		return
	}
}

// 适配gin框架路由
func RouteAdapter(b *builder.Resource, responseType string, ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	headerString := ""
	for hk, hvs := range ctx.Request.Header {
		tmp := ""
		for _, v := range hvs {
			tmp = tmp + "," + v
		}
		tmp = strings.Trim(tmp, ",")
		headerString = headerString + hk + ": " + tmp + "\r\n"
	}

	//把读过的字节流重新放到body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	// 将框架请求转换为builder框架请求
	request := &builder.Request{
		IPString:       ctx.ClientIP(),
		HeaderString:   headerString,
		MethodString:   ctx.Request.Method,
		FullPathString: ctx.FullPath(),
		HostString:     ctx.Request.Host,
		PathString:     ctx.Request.URL.Path,
		QueryString:    ctx.Request.URL.RawQuery,
		BodyBuffer:     body,
	}

	// 转换Request对象
	result, err := b.TransformRequest(request).Run()
	if err != nil {
		ctx.JSON(200, msg.Error(err.Error(), ""))
		return
	}

	// 响应结果
	switch responseType {
	case JSON_RESPONSE:
		ctx.JSON(200, result)
		return
	case IMAGE_RESPONSE:
		ctx.Writer.Write(result.([]byte))
		return
	case EXCEL_RESPONSE:
		ctx.Header("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Writer.Write(result.([]byte))
		return
	}
}

// 适配gin框架
func Adapter(b *builder.Resource, app *gin.Engine) {

	// 后台路由组
	rg := app.Group("/api/admin")

	// 登录
	rg.GET("/login/:resource/index", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/login/:resource/handle", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/login/:resource/captchaId", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/login/:resource/captcha/:id", func(ctx *gin.Context) {
		RouteAdapter(b, IMAGE_RESPONSE, ctx)
	})
	rg.GET("/logout/:resource/handle", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 仪表盘
	rg.GET("/dashboard/:resource/index", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})

	// 增删改查
	rg.GET("/:resource/index", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/editable", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.Any("/:resource/action/:uriKey", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/create", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/:resource/store", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/edit", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/edit/values", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.POST("/:resource/save", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/detail", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/export", func(ctx *gin.Context) {
		RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.Any("/:resource/import", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
	rg.GET("/:resource/import/template", func(ctx *gin.Context) {
		RouteAdapter(b, EXCEL_RESPONSE, ctx)
	})
	rg.GET("/:resource/:uriKey/form", func(ctx *gin.Context) {
		RouteAdapter(b, JSON_RESPONSE, ctx)
	})
}
