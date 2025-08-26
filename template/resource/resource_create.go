package resource

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

// 创建表单的接口
func (p *Template) CreationApi(ctx *quark.Context) string {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 表单接口
	formApi := template.FormApi(ctx)
	if formApi != "" {
		return formApi
	}

	uri := strings.Split(ctx.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(ctx.Path()).ReplaceLast("/index", "/store")
	}
	if uri[len(uri)-1] == "form" {
		return stringy.New(ctx.Path()).ReplaceLast("/form", "/store")
	}

	return stringy.New(ctx.Path()).ReplaceLast("/create", "/store")
}

// 渲染创建页组件
func (p *Template) CreationComponentRender(ctx *quark.Context, data map[string]interface{}) interface{} {

	// 表单标题
	title := p.FormTitle(ctx)

	// 表单页右上角自定义区域行为
	formExtraActions := p.FormExtraActions(ctx)

	// 创建表单的接口
	api := p.CreationApi(ctx)

	// 包裹在组件内的创建页字段
	fields := p.CreationFieldsWithinComponents(ctx)

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

// 创建页面显示前回调
func (p *Template) BeforeCreating(ctx *quark.Context) map[string]interface{} {
	return map[string]interface{}{}
}
