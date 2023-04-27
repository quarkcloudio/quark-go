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
func (p *Template) hasRouteMapping(method string, path string, handlerName string) bool {
	has := false
	for _, v := range p.RouteMapping {
		if v.Method == method && v.Path == path && v.HandlerName == handlerName {
			has = true
		}
	}
	return has
}

// 注册路由
func (p *Template) AddRouteMapping(method string, path string, handlerName string) *Template {
	if !p.hasRouteMapping(method, path, handlerName) {
		getRoute := &builder.RouteMapping{
			Method:      method,
			Path:        path,
			HandlerName: handlerName,
		}

		p.RouteMapping = append(p.RouteMapping, getRoute)
	}
	return p
}

// ANY请求
func (p *Template) Any(path string, handlerName string) {
	p.AddRouteMapping("Any", path, handlerName)
}

// GET请求
func (p *Template) GET(path string, handlerName string) {
	p.AddRouteMapping(http.MethodGet, path, handlerName)
}

// HEAD请求
func (p *Template) HEAD(path string, handlerName string) {
	p.AddRouteMapping(http.MethodHead, path, handlerName)
}

// OPTIONS请求
func (p *Template) OPTIONS(path string, handlerName string) {
	p.AddRouteMapping(http.MethodOptions, path, handlerName)
}

// POST请求
func (p *Template) POST(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPost, path, handlerName)
}

// PUT请求
func (p *Template) PUT(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPut, path, handlerName)
}

// PATCH请求
func (p *Template) PATCH(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPatch, path, handlerName)
}

// DELETE请求
func (p *Template) DELETE(path string, handlerName string) {
	p.AddRouteMapping(http.MethodDelete, path, handlerName)
}

// 默认组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {

	return ctx.JSONError("请实现组件渲染方法")
}
