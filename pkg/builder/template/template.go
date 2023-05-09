package template

import (
	"net/http"

	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/gorm"
)

// 模板
type Template struct {
	DB    *gorm.DB    // DB对象
	Model interface{} // DB模型结构体
}

// ANY请求
func (p *Template) Any(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping("Any", path, handler)
}

// GET请求
func (p *Template) GET(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodGet, path, handler)
}

// HEAD请求
func (p *Template) HEAD(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodHead, path, handler)
}

// OPTIONS请求
func (p *Template) OPTIONS(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodOptions, path, handler)
}

// POST请求
func (p *Template) POST(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodPost, path, handler)
}

// PUT请求
func (p *Template) PUT(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodPut, path, handler)
}

// PATCH请求
func (p *Template) PATCH(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodPatch, path, handler)
}

// DELETE请求
func (p *Template) DELETE(path string, handler func(ctx *builder.Context) error) {
	builder.AddRouteMapping(http.MethodDelete, path, handler)
}

// 默认组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {

	return ctx.JSONError("请实现组件渲染方法")
}
