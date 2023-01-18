package adminresource

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
)

//更新表单的接口
func (p *Template) UpdateApi(request *builder.Request, templateInstance interface{}) string {
	formApi := templateInstance.(interface{ FormApi(*builder.Request) string }).FormApi(request)
	if formApi != "" {
		return formApi
	}

	uri := strings.Split(request.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(request.Path()).ReplaceLast("/index", "/save")
	}

	return stringy.New(request.Path()).ReplaceLast("/edit", "/save")
}

// 编辑页面获取表单数据接口
func (p *Template) EditValueApi(request *builder.Request) string {
	uri := strings.Split(request.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(request.Path()).ReplaceLast("/index", "/edit/values?id=${id}")
	}

	return stringy.New(request.Path()).ReplaceLast("/edit", "/edit/values?id=${id}")
}

// 渲染编辑页组件
func (p *Template) UpdateComponentRender(request *builder.Request, templateInstance interface{}, data map[string]interface{}) interface{} {
	title := p.FormTitle(request, templateInstance)
	formExtraActions := p.FormExtraActions(request, templateInstance)
	api := p.UpdateApi(request, templateInstance)
	fields := p.UpdateFieldsWithinComponents(request, templateInstance)
	formActions := p.FormActions(request, templateInstance)

	return p.FormComponentRender(
		request,
		templateInstance,
		title,
		formExtraActions,
		api,
		fields,
		formActions,
		data,
	)
}

// 编辑页面显示前回调
func (p *Template) BeforeEditing(request *builder.Request, data map[string]interface{}) map[string]interface{} {
	return data
}
