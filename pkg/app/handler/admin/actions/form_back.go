package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder/actions"
)

type FormBack struct {
	actions.Action
}

// 初始化
func (p *FormBack) Init() *FormBack {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "default"

	// 文字
	p.Name = "返回上一页"

	// 行为类型
	p.ActionType = "back"

	// 设置展示位置
	p.SetShowOnForm().SetShowOnDetail()

	return p
}
