package searches

import "github.com/quarkcloudio/quark-go/v3"

type DateRange struct {
	Search
}

// 加载初始化数据
func (p *DateRange) New(ctx *quark.Context) interface{} {
	p.Component = "dateRangeField"

	return p
}
