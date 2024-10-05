package searches

import "github.com/quarkcloudio/quark-go/v3/pkg/builder"

type Date struct {
	Search
}

// 初始化模板
func (p *Date) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "dateField"

	return p
}
