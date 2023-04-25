package adminfrontpage

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
)

// 后台登录模板
type Template struct {
	template.Template
}

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 注册路由映射
	p.GET("/api/admin/frontpage/:resource/:component", "Render") // 自定义模板路由

	return p
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {
	data := map[string]interface{}{
		"component": ctx.Param("component"),
	}

	return ctx.JSON(200, data)
}
