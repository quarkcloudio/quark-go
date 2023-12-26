package searches

import "github.com/quarkcloudio/quark-go/v2/pkg/builder"

type Datetime struct {
	Search
}

// 初始化模板
func (p *Datetime) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "datetimeField"

	return p
}
