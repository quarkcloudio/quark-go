package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/card"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/types"
)

// 内容页标题
func (p *Template) ContentTitle(ctx *quark.Context) string {
	template := ctx.Template.(types.Resourcer)
	title := template.GetTitle()

	return title
}

// 渲染自定义内容页组件
func (p *Template) ContentComponentRender(ctx *quark.Context) interface{} {

	// 内容页标题
	title := p.ContentTitle(ctx)

	// 内容页右上角自定义区域行为
	extraActions := p.DetailExtraActions(ctx)

	// 包裹在组件内的内容页字段
	body := p.Content(ctx)

	return p.ContentWithinCard(
		ctx,
		title,
		extraActions,
		body,
	)
}

// 在卡片内的内容页组件
func (p *Template) ContentWithinCard(
	ctx *quark.Context,
	title string,
	extra interface{},
	content interface{}) interface{} {

	return (&card.Component{}).
		Init().
		SetTitle(title).
		SetHeaderBordered(true).
		SetExtra(extra).
		SetBody(content)
}
