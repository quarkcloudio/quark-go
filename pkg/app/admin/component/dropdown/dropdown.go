package dropdown

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Label              string                 `json:"label"`
	Block              bool                   `json:"block"`
	Danger             bool                   `json:"danger"`
	Disabled           bool                   `json:"disabled"`
	Ghost              bool                   `json:"ghost"`
	Icon               string                 `json:"icon"`
	Shape              string                 `json:"shape"`
	Size               string                 `json:"size"`
	Type               string                 `json:"type"`
	Arrow              bool                   `json:"arrow"`
	DestroyPopupOnHide bool                   `json:"destroyPopupOnHide"`
	Menu               interface{}            `json:"menu"`
	OverlayClassName   string                 `json:"overlayClassName"`
	OverlayStyle       map[string]interface{} `json:"overlayStyle"`
	Placement          string                 `json:"placement"`
	Trigger            []string               `json:"trigger"`
	Visible            bool                   `json:"-"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Item
func NewItem() *Item {

	return (&Item{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "dropdown"
	p.Visible = true

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 将按钮宽度调整为其父宽度的选项
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 设置危险按钮
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = "icon-" + icon

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 下拉框箭头是否显示
func (p *Component) SetArrow(arrow bool) *Component {
	p.Arrow = arrow

	return p
}

// 关闭后是否销毁 Dropdown
func (p *Component) SetDestroyPopupOnHide(destroyPopupOnHide bool) *Component {
	p.DestroyPopupOnHide = destroyPopupOnHide

	return p
}

// 菜单
func (p *Component) SetMenu(menu interface{}) *Component {
	p.Menu = menu

	return p
}

// 下拉根元素的类名称
func (p *Component) SetOverlayClassName(overlayClassName string) *Component {
	p.OverlayClassName = overlayClassName

	return p
}

// 下拉根元素的样式
func (p *Component) SetOverlayStyle(overlayStyle map[string]interface{}) *Component {
	p.OverlayStyle = overlayStyle

	return p
}

// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
func (p *Component) SetPlacement(placement string) *Component {
	p.Placement = placement

	return p
}

// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
func (p *Component) SetTrigger(trigger []string) *Component {
	p.Trigger = trigger

	return p
}

// 菜单是否显示
func (p *Component) SetVisible(visible bool) *Component {
	p.Visible = visible

	return p
}
