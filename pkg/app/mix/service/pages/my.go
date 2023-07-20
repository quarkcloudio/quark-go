package pages

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/mix/template/page"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
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
	return "我的页面"
}
