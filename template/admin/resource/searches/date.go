package searches

import "github.com/quarkcloudio/quark-go/v3"

type Date struct {
	Search
}

// 加载初始化数据
func (p *Date) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "dateField"

	return p
}
