package resources

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
)

type Demo struct {
	adminresource.Template
}

// 初始化
func (p *Demo) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

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
