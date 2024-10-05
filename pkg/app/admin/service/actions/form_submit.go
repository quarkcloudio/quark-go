package actions

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

type FormSubmitAction struct {
	actions.Action
}

// 表单提交，FormSubmit() | FormSubmit("提交")
func FormSubmit(options ...interface{}) *FormSubmitAction {
	action := &FormSubmitAction{}

	// 文字
	action.Name = "提交"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *FormSubmitAction) Init(ctx *builder.Context) interface{} {

	// 类型
	p.Type = "primary"

	// 行为类型
	p.ActionType = "submit"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnForm(true)

	return p
}
