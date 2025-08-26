package searches

import "github.com/quarkcloudio/quark-go/v4"

type DatetimeRange struct {
	Search
}

// 加载初始化数据
func (p *DatetimeRange) New(ctx *quark.Context) interface{} {
	p.Component = "datetimeRangeField"

	return p
}
