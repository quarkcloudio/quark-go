package searches

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Radio struct {
	Search
	RadioOptions []*radio.Option
}

// 初始化模板
func (p *Radio) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "radioField"
	return p
}

// 设置Option
func (p *Radio) Option(value interface{}, label string) *radio.Option {

	return &radio.Option{
		Value: value,
		Label: label,
	}
}
