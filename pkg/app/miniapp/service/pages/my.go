package pages

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/navbar"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/template/page"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type My struct {
	page.Template
}

// 初始化
func (p *My) Init(ctx *builder.Context) interface{} {
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
