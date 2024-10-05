package actions

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

type FormExtraBackAction struct {
	actions.Action
}

// 表单扩展行为返回上一页，FormExtraBack() | FormExtraBack("返回上一页")
func FormExtraBack(options ...interface{}) *FormExtraBackAction {
	action := &FormExtraBackAction{}

	// 文字
	action.Name = "返回上一页"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *FormExtraBackAction) Init(ctx *builder.Context) interface{} {

	// 类型
	p.Type = "link"

	// 行为类型
	p.ActionType = "back"

	// 在表单页右上角自定义区域展示
	p.SetShowOnFormExtra()

	// 在详情页右上角自定义区域展示
	p.SetShowOnDetailExtra()

	return p
}
