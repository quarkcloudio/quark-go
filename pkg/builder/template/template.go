package template

import (
	"net/http"

	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/gorm"
)

// 模板
type Template struct {
	DB           *gorm.DB                // DB对象
	Model        interface{}             // DB模型结构体
	RouteMapping []*builder.RouteMapping // 路由映射
}

// 获取路由
func (p *Template) GetRouteMapping() []*builder.RouteMapping {
	return p.RouteMapping
}

// 是否存在路由
func (p *Template) hasRouteMapping(method string, path string, handler func(ctx *builder.Context) error) bool {
	has := false
	for _, v := range p.RouteMapping {
		if v.Method == method && v.Path == path {
			has = true
		}
	}
	return has
}

// 注册路由
func (p *Template) AddRouteMapping(method string, path string, handler func(ctx *builder.Context) error) *Template {
	if !p.hasRouteMapping(method, path, handler) {
		getRoute := &builder.RouteMapping{
			Method:  method,
			Path:    path,
			Handler: handler,
		}

		p.RouteMapping = append(p.RouteMapping, getRoute)
	}
	return p
}

// ANY请求
func (p *Template) Any(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping("Any", path, handler)
}

// GET请求
func (p *Template) GET(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodGet, path, handler)
}

// HEAD请求
func (p *Template) HEAD(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodHead, path, handler)
}

// OPTIONS请求
func (p *Template) OPTIONS(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodOptions, path, handler)
}

// POST请求
func (p *Template) POST(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodPost, path, handler)
}

// PUT请求
func (p *Template) PUT(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodPut, path, handler)
}

// PATCH请求
func (p *Template) PATCH(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodPatch, path, handler)
}

// DELETE请求
func (p *Template) DELETE(path string, handler func(ctx *builder.Context) error) {
	p.AddRouteMapping(http.MethodDelete, path, handler)
}

// 默认组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {

	return ctx.JSONError("请实现组件渲染方法")
}
