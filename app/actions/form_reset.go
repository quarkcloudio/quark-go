package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
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
func (p *FormResetAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "reset"

	// 设置展示位置
	p.SetShowOnForm()

	return p
}
