package resources

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Demo struct {
	resource.Template
}

// 初始化
func (p *Demo) Init(ctx *builder.Context) interface{} {

	// 初始化模板
	p.TemplateInit(ctx)

	return p
}

// 字段
func (p *Demo) Fields(ctx *builder.Context) []interface{} {

	return []interface{}{}
}

// 搜索
func (p *Demo) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{}
}

// 行为
func (p *Demo) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{}
}
