package actions

import "github.com/quarkcloudio/quark-go/v4"

type Batch struct {
	Action
}

// 初始化
func (p *Batch) New(ctx *quark.Context) interface{} {
	p.ActionType = "ajax"
	p.Type = "default"
	p.Batch = true
	p.Size = "small"

	return p
}
