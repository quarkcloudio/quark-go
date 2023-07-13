package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
)

type FormExtraBack struct {
	actions.Action
}

// 初始化
func (p *FormExtraBack) Init() *FormExtraBack {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "link"

	// 文字
	p.Name = "返回上一页"

	// 行为类型
	p.ActionType = "back"

	// 设置展示位置
	p.SetShowOnFormExtra().SetShowOnDetailExtra()

	return p
}
