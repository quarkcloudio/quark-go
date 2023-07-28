package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type FormReset struct {
	actions.Action
}

// 初始化
func (p *FormReset) Init(ctx *builder.Context) interface{} {

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
