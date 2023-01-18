package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder/actions"
)

type FormSubmit struct {
	actions.Action
}

// 初始化
func (p *FormSubmit) Init() *FormSubmit {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "primary"

	// 文字
	p.Name = "提交"

	// 行为类型
	p.ActionType = "submit"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnForm(true)

	return p
}
