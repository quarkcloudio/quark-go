package builder

import (
	"net/http"

	"gorm.io/gorm"
)

// 模版接口
type Templater interface {

	// 初始化
	Init(ctx *Context) interface{}

	// 初始化模板
	TemplateInit(ctx *Context) interface{}

	// 初始化路由
	RouteInit() interface{}

	// 自定义路由
	Route() interface{}

	// 获取路由
	GetRouteMapping() []*RouteMapping

	// 添加路由
	AddRouteMapping(method string, path string, handler func(ctx *Context) error) *Template

	// ANY请求
	Any(path string, handler func(ctx *Context) error)

	// GET请求
	GET(path string, handler func(ctx *Context) error)

	// HEAD请求
	HEAD(path string, handler func(ctx *Context) error)

	// OPTIONS请求
	OPTIONS(path string, handler func(ctx *Context) error)

	// POST请求
	POST(path string, handler func(ctx *Context) error)

	// PUT请求
	PUT(path string, handler func(ctx *Context) error)

	// PATCH请求
	PATCH(path string, handler func(ctx *Context) error)

	// DELETE请求
	DELETE(path string, handler func(ctx *Context) error)
}

// 模板
type Template struct {
	DB           *gorm.DB        // DB对象
	RouteMapping []*RouteMapping // 路由映射
}

// 获取路由
func (p *Template) GetRouteMapping() []*RouteMapping {
	return p.RouteMapping
}

// 自定义路由
func (p *Template) Route() interface{} {
	return p
}

// 是否存在路由
func (p *Template) hasRouteMapping(method string, path string, handler func(ctx *Context) error) bool {
	has := false
	for _, v := range p.RouteMapping {
		if v.Method == method && v.Path == path {
			has = true
		}
	}
	return has
}

// 注册路由
func (p *Template) AddRouteMapping(method string, path string, handler func(ctx *Context) error) *Template {
	if !p.hasRouteMapping(method, path, handler) {
		getRoute := &RouteMapping{
			Method:  method,
			Path:    path,
			Handler: handler,
		}

		p.RouteMapping = append(p.RouteMapping, getRoute)
	}
	return p
}

// ANY请求
func (p *Template) Any(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping("Any", path, handler)
}

// GET请求
func (p *Template) GET(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodGet, path, handler)
}

// HEAD请求
func (p *Template) HEAD(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodHead, path, handler)
}

// OPTIONS请求
func (p *Template) OPTIONS(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodOptions, path, handler)
}

// POST请求
func (p *Template) POST(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodPost, path, handler)
}

// PUT请求
func (p *Template) PUT(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodPut, path, handler)
}

// PATCH请求
func (p *Template) PATCH(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodPatch, path, handler)
}

// DELETE请求
func (p *Template) DELETE(path string, handler func(ctx *Context) error) {
	p.AddRouteMapping(http.MethodDelete, path, handler)
}
