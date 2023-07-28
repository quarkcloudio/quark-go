package searches

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Date struct {
	Search
}

// 初始化模板
func (p *Date) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "dateField"

	return p
}
