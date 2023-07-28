package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type FormBack struct {
	actions.Action
}

// 初始化
func (p *FormBack) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "返回上一页"

	// 类型
	p.Type = "default"

	// 行为类型
	p.ActionType = "back"

	// 设置展示位置
	p.SetShowOnForm().SetShowOnDetail()

	return p
}
