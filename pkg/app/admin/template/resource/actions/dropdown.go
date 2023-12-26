package actions

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/menu"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Dropdown struct {
	Action
	Arrow        bool                   `json:"arrow"`        // 下拉框箭头是否显示
	Placement    string                 `json:"placement"`    // 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
	Trigger      []string               `json:"trigger"`      // 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
	OverlayStyle map[string]interface{} `json:"overlayStyle"` // 下拉根元素的样式
	Actions      []interface{}          `json:"actions"`      // 下拉菜单行为
}

// 初始化
func (p *Dropdown) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "dropdown"
	p.Placement = "bottomLeft"
	p.Trigger = append(p.Trigger, "hover")

	return p
}

// 是否显示箭头图标
func (p *Dropdown) GetArrow() bool {
	return p.Arrow
}

// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
func (p *Dropdown) GetPlacement() string {
	return p.Placement
}

// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
func (p *Dropdown) GetTrigger() []string {
	return p.Trigger
}

// 下拉根元素的样式
func (p *Dropdown) GetOverlayStyle() map[string]interface{} {
	return p.OverlayStyle
}

// 菜单
func (p *Dropdown) GetMenu(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := p.GetActions()

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		items = append(items, template.BuildAction(ctx, actionInstance))
	}

	return (&menu.Component{}).Init().SetItems(items)
}

// 下拉菜单行为
func (p *Dropdown) SetActions(actions []interface{}) interface{} {
	p.Actions = actions

	return p
}

// 获取下拉菜单行为
func (p *Dropdown) GetActions() []interface{} {
	return p.Actions
}
