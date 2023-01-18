package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder/actions"
)

type FormReset struct {
	actions.Action
}

// 初始化
func (p *FormReset) Init() *FormReset {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "default"

	// 文字
	p.Name = "重置"

	// 行为类型
	p.ActionType = "reset"

	// 设置展示位置
	p.SetShowOnForm()

	return p
}
