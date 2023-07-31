package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type FormBackAction struct {
	actions.Action
}

// 返回上一页
func FormBack() *FormBackAction {
	return &FormBackAction{}
}

// 初始化
func (p *FormBackAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "返回上一页"

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
