package pages

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/col"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/image"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/template/page"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Index struct {
	page.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {
	return p
}

// 轮播图
func (p *Index) Banners(ctx *builder.Context) []*image.Component {
	return []*image.Component{
		p.Image("https://storage.360buyimg.com/jdc-article/NutUItaro34.jpg"),
		p.Image("https://storage.360buyimg.com/jdc-article/NutUItaro2.jpg"),
	}
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {
	return p.Row([]*col.Component{
		p.Col(12, "Hello World!"),
		p.Col(12, "你好，世界!"),
		p.Col(24,
			p.Action("确定", "primary").
				SetBlock(true).
				SetLink("/pages/engine/custom?api=/api/miniapp/form/index/index", "navigateTo"),
		),
	})
}
