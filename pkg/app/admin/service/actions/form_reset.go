package actions

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type FormResetAction struct {
	actions.Action
}

// 表单重置，FormReset() | FormReset("重置")
func FormReset(options ...interface{}) *FormResetAction {
	action := &FormResetAction{}

	// 文字
	action.Name = "重置"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *FormResetAction) Init(ctx *builder.Context) interface{} {

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "reset"

	// 设置展示位置
	p.SetShowOnForm()

	return p
}
