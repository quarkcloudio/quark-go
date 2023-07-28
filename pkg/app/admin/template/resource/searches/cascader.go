package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/cascader"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Cascader struct {
	Search
	CascaderOptions []*cascader.Option
}

// 初始化模板
func (p *Cascader) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "cascaderField"

	return p
}

// 设置Option
func (p *Cascader) Option(value interface{}, label string) *cascader.Option {

	return &cascader.Option{
		Value: value,
		Label: label,
	}
}
