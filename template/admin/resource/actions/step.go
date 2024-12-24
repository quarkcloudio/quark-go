package actions

import "github.com/quarkcloudio/quark-go/v3"

type Step struct {
	Action
}

// 初始化
func (p *Step) TemplateInit(ctx *quark.Context) interface{} {
	p.ActionType = "step"
	p.Name = []string{"上一步", "下一步"}
	return p
}
