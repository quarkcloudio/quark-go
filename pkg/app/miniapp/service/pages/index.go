package pages

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/component/col"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/template/page"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Index struct {
	page.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {
	return p
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {
	return []interface{}{
		p.Row([]*col.Component{
			p.Col(24, "Hello World!"),
		}).SetStyle("text-align:center;"),
		p.Row([]*col.Component{
			p.Col(24,
				p.Action("跳转到表单页", "primary").
					SetBlock(true).
					SetLink("/pages/engine/default?api=/api/miniapp/form/demo/index", "navigateTo"),
			),
		}).SetStyle("text-align:center;"),
	}
}
