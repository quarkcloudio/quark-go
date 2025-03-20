package searches

import "github.com/quarkcloudio/quark-go/v3"

type Datetime struct {
	Search
}

// 加载初始化数据
func (p *Datetime) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "datetimeField"

	return p
}
