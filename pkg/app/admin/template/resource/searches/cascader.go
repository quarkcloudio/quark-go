package searches

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/cascader"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
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
