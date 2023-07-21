package resource

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 更新表单的接口
func (p *Template) UpdateApi(ctx *builder.Context) string {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 表单接口
	formApi := template.FormApi(ctx)
	if formApi != "" {
		return formApi
	}

	uri := strings.Split(ctx.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(ctx.Path()).ReplaceLast("/index", "/save")
	}

	return stringy.New(ctx.Path()).ReplaceLast("/edit", "/save")
}

// 编辑页面获取表单数据接口
func (p *Template) EditValueApi(request *builder.Context) string {
	uri := strings.Split(request.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(request.Path()).ReplaceLast("/index", "/edit/values?id=${id}")
	}

	return stringy.New(request.Path()).ReplaceLast("/edit", "/edit/values?id=${id}")
}

// 渲染编辑页组件
func (p *Template) UpdateComponentRender(ctx *builder.Context, data map[string]interface{}) interface{} {

	// 表单标题
	title := p.FormTitle(ctx)

	// 表单页右上角自定义区域行为
	formExtraActions := p.FormExtraActions(ctx)

	// 更新表单的接口
	api := p.UpdateApi(ctx)

	// 包裹在组件内的编辑页字段
	fields := p.UpdateFieldsWithinComponents(ctx)

	// 表单页行为
	formActions := p.FormActions(ctx)

	return p.FormComponentRender(
		ctx,
		title,
		formExtraActions,
		api,
		fields,
		formActions,
		data,
	)
}

// 编辑页面显示前回调
func (p *Template) BeforeEditing(request *builder.Context, data map[string]interface{}) map[string]interface{} {
	return data
}
