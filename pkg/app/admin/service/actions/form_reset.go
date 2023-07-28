package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type FormResetAction struct {
	actions.Action
}

// 表单重置
func FormReset() *FormResetAction {
	return &FormResetAction{}
}

// 初始化
func (p *FormResetAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "重置"

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "reset"

	// 设置展示位置
	p.SetShowOnForm()

	return p
}
