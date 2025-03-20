package searches

import "github.com/quarkcloudio/quark-go/v3"

type DatetimeRange struct {
	Search
}

// 初始化模板
func (p *DatetimeRange) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "datetimeRangeField"

	return p
}
