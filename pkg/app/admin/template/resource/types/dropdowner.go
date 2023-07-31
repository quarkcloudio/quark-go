package types

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Dropdowner interface {
	Actioner

	// 是否显示箭头图标
	GetArrow() bool

	// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
	GetPlacement() string

	// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
	GetTrigger() []string

	// 下拉根元素的样式
	GetOverlayStyle() map[string]interface{}

	// 菜单
	GetMenu(ctx *builder.Context) interface{}

	// 下拉菜单行为
	SetActions(actions []interface{}) interface{}

	// 获取下拉菜单行为
	GetActions() []interface{}
}
