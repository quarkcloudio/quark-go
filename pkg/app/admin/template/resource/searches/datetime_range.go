package searches

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type DatetimeRange struct {
	Search
}

// 初始化模板
func (p *DatetimeRange) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "datetimeRangeField"

	return p
}
