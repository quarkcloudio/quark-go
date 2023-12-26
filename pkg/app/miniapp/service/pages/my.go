package pages

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/template/page"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type My struct {
	page.Template
}

// 初始化
func (p *My) Init(ctx *builder.Context) interface{} {
	return p
}

// 组件渲染
func (p *My) Content(ctx *builder.Context) interface{} {
	return "我的"
}
