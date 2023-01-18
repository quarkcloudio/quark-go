package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/card"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
)

// 详情页标题
func (p *Template) DetailTitle(request *builder.Request, templateInstance interface{}) string {
	value := reflect.ValueOf(templateInstance).Elem()
	title := value.FieldByName("Title").String()

	return title + "详情"
}

// 渲染详情页组件
func (p *Template) DetailComponentRender(request *builder.Request, templateInstance interface{}, data map[string]interface{}) interface{} {
	title := p.DetailTitle(request, templateInstance)
	formExtraActions := p.DetailExtraActions(request, templateInstance)
	fields := p.DetailFieldsWithinComponents(request, templateInstance, data)
	formActions := p.DetailActions(request, templateInstance)

	return p.DetailWithinCard(
		request,
		templateInstance,
		title,
		formExtraActions,
		fields,
		formActions,
		data,
	)
}

// 在卡片内的详情页组件
func (p *Template) DetailWithinCard(
	request *builder.Request,
	templateInstance interface{},
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
	request *builder.Request,
	templateInstance interface{},
	title string,
	extra interface{},
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	return (&tabs.Component{}).Init().SetTabPanes(fields).SetTabBarExtraContent(extra)
}

// 详情页页面显示前回调
func (p *Template) BeforeDetailShowing(request *builder.Request, data map[string]interface{}) map[string]interface{} {
	return data
}
