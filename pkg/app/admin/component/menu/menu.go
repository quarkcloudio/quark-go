package menu

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	DefaultOpenKeys      interface{} `json:"defaultOpenKeys"`
	DefaultSelectedKeys  interface{} `json:"defaultSelectedKeys"`
	InlineCollapsed      bool        `json:"inlineCollapsed"`
	InlineIndent         int         `json:"inlineIndent"`
	Mode                 string      `json:"mode"`
	Multiple             bool        `json:"multiple"`
	Selectable           bool        `json:"selectable"`
	SubMenuCloseDelay    float64     `json:"subMenuCloseDelay"`
	SubMenuOpenDelay     float64     `json:"subMenuOpenDelay"`
	Theme                string      `json:"theme"`
	TriggerSubMenuAction string      `json:"triggerSubMenuAction"`
	Items                interface{} `json:"items"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Divider
func NewDivider() *Divider {
	return (&Divider{}).Init()
}

// 获取ItemGroup
func NewItemGroup() *ItemGroup {
	return (&ItemGroup{}).Init()
}

// 获取Item
func NewItem() *Item {
	return (&Item{}).Init()
}

// 获取SubMenu
func NewSubMenu() *SubMenu {
	return (&SubMenu{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "menu"
	p.InlineIndent = 24
	p.Mode = "vertical"
	p.Selectable = true
	p.SubMenuCloseDelay = 0.1
	p.SubMenuOpenDelay = 0
	p.Theme = "light"
	p.TriggerSubMenuAction = "hover"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 初始展开的 SubMenu 菜单项 key 数组
func (p *Component) SetDefaultOpenKeys(defaultOpenKeys interface{}) *Component {
	p.DefaultOpenKeys = defaultOpenKeys

	return p
}

// 初始选中的菜单项 key 数组
func (p *Component) SetDefaultSelectedKeys(defaultSelectedKeys interface{}) *Component {
	p.DefaultSelectedKeys = defaultSelectedKeys

	return p
}

// inline 时菜单是否收起状态
func (p *Component) SetInlineCollapsed(inlineCollapsed bool) *Component {
	p.InlineCollapsed = inlineCollapsed

	return p
}

// inline 模式的菜单缩进宽度
func (p *Component) SetInlineIndent(inlineIndent int) *Component {
	p.InlineIndent = inlineIndent

	return p
}

// 菜单类型，现在支持垂直、水平、和内嵌模式三种,vertical | horizontal | inline
func (p *Component) SetMode(mode string) *Component {
	p.Mode = mode

	return p
}

// 是否允许多选
func (p *Component) SetMultiple(multiple bool) *Component {
	p.Multiple = multiple

	return p
}

// 是否允许选中
func (p *Component) SetSelectable(selectable bool) *Component {
	p.Selectable = selectable

	return p
}

// 用户鼠标离开子菜单后关闭延时，单位：秒
func (p *Component) SetSubMenuCloseDelay(subMenuCloseDelay float64) *Component {
	p.SubMenuCloseDelay = subMenuCloseDelay

	return p
}

// 主题颜色,light | dark
func (p *Component) SetTheme(theme string) *Component {
	p.Theme = theme

	return p
}

// SubMenu 展开/关闭的触发行为,hover | click
func (p *Component) SetTriggerSubMenuAction(triggerSubMenuAction string) *Component {
	p.TriggerSubMenuAction = triggerSubMenuAction

	return p
}

// 设置菜单项
func (p *Component) SetItems(items interface{}) *Component {
	p.Items = items

	return p
}

// 菜单分隔符
func (p *Component) Divider() *Divider {
	return (&Divider{}).Init()
}

// 菜单分组
func (p *Component) ItemGroup(title string, items interface{}) *ItemGroup {
	return (&ItemGroup{}).Init().SetTitle(title).SetItems(items)
}

// 菜单项
func (p *Component) Item(label string, title string) *Item {
	return (&Item{}).Init().SetLabel(label).SetTitle(title)
}

// 子菜单
func (p *Component) SubMenu(title string, items interface{}) *SubMenu {
	return (&SubMenu{}).Init().SetTitle(title).SetItems(items)
}
