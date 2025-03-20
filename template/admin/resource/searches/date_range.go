package searches

import "github.com/quarkcloudio/quark-go/v3"

type DateRange struct {
	Search
}

// 初始化模板
func (p *DateRange) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "dateRangeField"

	return p
}
