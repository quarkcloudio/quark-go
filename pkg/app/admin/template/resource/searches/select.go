package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Select struct {
	Search
	SelectOptions []*selectfield.Option
}

// 初始化模板
func (p *Select) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "selectField"

	return p
}

// 设置Option
func (p *Select) Option(value interface{}, label string) *selectfield.Option {

	return &selectfield.Option{
		Value: value,
		Label: label,
	}
}
