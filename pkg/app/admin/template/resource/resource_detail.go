package resource

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/card"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tabs"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 详情页标题
func (p *Template) DetailTitle(ctx *builder.Context) string {
	template := ctx.Template.(types.Resourcer)
	title := template.GetTitle()

	return title + "详情"
}

// 渲染详情页组件
func (p *Template) DetailComponentRender(ctx *builder.Context, data map[string]interface{}) interface{} {

	// 详情页标题
	title := p.DetailTitle(ctx)

	// 详情页右上角自定义区域行为
	formExtraActions := p.DetailExtraActions(ctx)

	// 包裹在组件内的详情页字段
	fields := p.DetailFieldsWithinComponents(ctx, data)

	// 包裹在组件内的详情页字段
	formActions := p.DetailActions(ctx)

	return p.DetailWithinCard(
		ctx,
		title,
		formExtraActions,
		fields,
		formActions,
		data,
	)
}

// 在卡片内的详情页组件
func (p *Template) DetailWithinCard(
	ctx *builder.Context,
	title string,
	extra interface{},
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	return (&card.Component{}).
		Init().
		SetTitle(title).
		SetHeaderBordered(true).
		SetExtra(extra).
		SetBody(fields)
}

// 在标签页内的详情页组件
func (p *Template) DetailWithinTabs(
	ctx *builder.Context,
	title string,
	extra interface{},
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	return (&tabs.Component{}).Init().SetTabPanes(fields).SetTabBarExtraContent(extra)
}

// 详情页页面显示前回调
func (p *Template) BeforeDetailShowing(ctx *builder.Context, data map[string]interface{}) map[string]interface{} {
	return data
}
