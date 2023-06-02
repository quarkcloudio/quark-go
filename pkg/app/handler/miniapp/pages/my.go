package pages

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/miniapppage"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/navbar"
)

type My struct {
	miniapppage.Template
}

// 初始化
func (p *My) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 头部导航
func (p *My) Navbar(ctx *builder.Context, navbar *navbar.Component) interface{} {
	return navbar.SetTitle("我的")
}

// 组件渲染
func (p *My) Content(ctx *builder.Context) interface{} {
	return "我的"
}
