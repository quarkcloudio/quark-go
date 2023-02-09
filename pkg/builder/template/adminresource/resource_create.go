package adminresource

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
)

// 创建表单的接口
func (p *Template) CreationApi(ctx *builder.Context) string {
	formApi := ctx.Template.(interface {
		FormApi(ctx *builder.Context) string
	}).FormApi(ctx)
	if formApi != "" {
		return formApi
	}

	uri := strings.Split(ctx.Path(), "/")
	if uri[len(uri)-1] == "index" {
		return stringy.New(ctx.Path()).ReplaceLast("/index", "/store")
	}

	return stringy.New(ctx.Path()).ReplaceLast("/create", "/store")
}

// 渲染创建页组件
func (p *Template) CreationComponentRender(ctx *builder.Context, data map[string]interface{}) interface{} {
	title := p.FormTitle(ctx)
	formExtraActions := p.FormExtraActions(ctx)
	api := p.CreationApi(ctx)
	fields := p.CreationFieldsWithinComponents(ctx)
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
func (p *Template) BeforeCreating(ctx *builder.Context) map[string]interface{} {
	return map[string]interface{}{}
}
