package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
)

type MoreActions struct {
	actions.Dropdown
}

// 初始化
func (p *MoreActions) Init(name string) *MoreActions {

	// 初始化父结构
	p.ParentInit()

	// 下拉框箭头是否显示
	p.Arrow = true

	// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
	p.Placement = "bottomLeft"

	// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
	p.Trigger = []string{"hover"}

	// 下拉根元素的样式
	p.OverlayStyle = map[string]interface{}{
		"zIndex": 999,
	}

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 文字
	p.Name = name

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}
