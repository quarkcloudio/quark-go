package adminresource

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
)

// 创建表单的接口
func (p *Template) CreationApi(request *builder.Request, templateInstance interface{}) string {
	formApi := templateInstance.(interface {
		FormApi(request *builder.Request) string
	}).FormApi(request)
	if formApi != "" {
		return formApi
	}

	uri := strings.Split(request.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(request.Path()).ReplaceLast("/index", "/store")
	}

	return stringy.New(request.Path()).ReplaceLast("/create", "/store")
}

// 渲染创建页组件
func (p *Template) CreationComponentRender(request *builder.Request, templateInstance interface{}, data map[string]interface{}) interface{} {
	title := p.FormTitle(request, templateInstance)
	formExtraActions := p.FormExtraActions(request, templateInstance)
	api := p.CreationApi(request, templateInstance)
	fields := p.CreationFieldsWithinComponents(request, templateInstance)
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

// 创建页面显示前回调
func (p *Template) BeforeCreating(request *builder.Request) map[string]interface{} {
	return map[string]interface{}{}
}
