package template

import (
	"net/http"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

// 模板
type ToolTemplate struct {
	DB           *gorm.DB                // DB对象
	Model        interface{}             // DB模型结构体
	RouteMapping []*builder.RouteMapping // 路由映射
}

// 获取路由
func (p *ToolTemplate) GetRouteMapping() []*builder.RouteMapping {
	return p.RouteMapping
}

// 是否存在路由
func (p *ToolTemplate) hasRouteMapping(method string, path string, handlerName string) bool {
	has := false
	for _, v := range p.RouteMapping {
		if v.Method == method && v.Path == path && v.HandlerName == handlerName {
			has = true
		}
	}

	return has
}

// 注册路由
func (p *ToolTemplate) AddRouteMapping(method string, path string, handlerName string) *ToolTemplate {
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

// ANY is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *ToolTemplate) Any(path string, handlerName string) {
	p.AddRouteMapping("Any", path, handlerName)
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *ToolTemplate) GET(path string, handlerName string) {
	p.AddRouteMapping(http.MethodGet, path, handlerName)
}

// HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)
func (p *ToolTemplate) HEAD(path string, handlerName string) {
	p.AddRouteMapping(http.MethodHead, path, handlerName)
}

// OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)
func (p *ToolTemplate) OPTIONS(path string, handlerName string) {
	p.AddRouteMapping(http.MethodOptions, path, handlerName)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (p *ToolTemplate) POST(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPost, path, handlerName)
}

// PUT is a shortcut for router.Handle(http.MethodPut, path, handle)
func (p *ToolTemplate) PUT(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPut, path, handlerName)
}

// PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)
func (p *ToolTemplate) PATCH(path string, handlerName string) {
	p.AddRouteMapping(http.MethodPatch, path, handlerName)
}

// DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)
func (p *ToolTemplate) DELETE(path string, handlerName string) {
	p.AddRouteMapping(http.MethodDelete, path, handlerName)
}

// 默认组件渲染
func (p *ToolTemplate) Render(ctx *builder.Context) interface{} {

	return ctx.JSON(200, msg.Error("请实现组件渲染方法", ""))
}
