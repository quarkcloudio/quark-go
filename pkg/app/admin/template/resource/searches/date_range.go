package searches

import "github.com/quarkcloudio/quark-go/v3/pkg/builder"

type DateRange struct {
	Search
}

// 初始化模板
func (p *DateRange) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "dateRangeField"

	return p
}
