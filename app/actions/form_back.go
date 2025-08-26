package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
)

type FormBackAction struct {
	actions.Action
}

// 返回上一页，FormBack() | FormBack("返回上一页")
func FormBack(options ...interface{}) *FormBackAction {
	action := &FormBackAction{}

	// 文字
	action.Name = "返回上一页"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *FormBackAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "back"

	// 在表单页展示
	p.SetShowOnForm()

	// 在详情页展示
	p.SetShowOnDetail()

	return p
}
