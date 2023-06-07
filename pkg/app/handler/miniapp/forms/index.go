package forms

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/miniappform"
)

type Index struct {
	miniappform.Template
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 字段
func (p *Index) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{
		p.Field().Input("username", "姓名"),
	}
}
