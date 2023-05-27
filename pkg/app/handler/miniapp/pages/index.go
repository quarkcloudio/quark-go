package pages

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/miniapppage"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/navbar"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/tabbar"
)

type Index struct {
	miniapppage.Template
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 头部导航
func (p *Index) Navbar(ctx *builder.Context, navbar *navbar.Component) interface{} {
	return navbar.SetTitle("首页")
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {
	return "Hello World!"
}

// 底部导航
func (p *Index) Tabbar(ctx *builder.Context, tabbar *tabbar.Component) interface{} {
	return tabbar.SetBottom(true)
}
