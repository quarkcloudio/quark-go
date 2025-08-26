package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
)

type FormStepAction struct {
	actions.Action
}

// Tabs表单上一步、下一步按钮，FormStep() | FormStep("上一步","下一步")
func FormStep(options ...interface{}) *FormStepAction {
	action := &FormStepAction{}

	// 文字
	action.Name = []string{"上一步", "下一步"}
	if len(options) == 2 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *FormStepAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "step"

	// 设置展示位置
	p.SetShowOnForm()

	return p
}
