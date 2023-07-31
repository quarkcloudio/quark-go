package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type FormExtraBackAction struct {
	actions.Action
}

// 表单扩展行为返回上一页
func FormExtraBack() *FormExtraBackAction {
	return &FormExtraBackAction{}
}

// 初始化
func (p *FormExtraBackAction) Init(ctx *builder.Context) interface{} {
	// 文字
	p.Name = "返回上一页"

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
