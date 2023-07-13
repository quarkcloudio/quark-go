package tabbar

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Bottom              bool    `json:"bottom"`
	UnactiveColor       string  `json:"unactiveColor"`
	ActiveColor         string  `json:"activeColor"`
	SafeAreaInsetBottom bool    `json:"safeAreaInsetBottom"`
	Placeholder         bool    `json:"placeholder"`
	Items               []*Item `json:"items"`
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
	p.Component = "tabbar"
	p.SetKey("tabbar", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 是否固定在页面底部
func (p *Component) SetBottom(bottom bool) *Component {
	p.Bottom = bottom

	return p
}

// icon未激活的颜色
func (p *Component) SetUnactiveColor(unactiveColor string) *Component {
	p.UnactiveColor = unactiveColor

	return p
}

// icon激活的颜色
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor

	return p
}

// 是否开启iphone系列全面屏底部安全区适配
func (p *Component) SetSafeAreaInsetBottom(safeAreaInsetBottom bool) *Component {
	p.SafeAreaInsetBottom = safeAreaInsetBottom

	return p
}

// 固定在底部时，是否在标签位置生成一个等高的占位元素
func (p *Component) SetPlaceholder(placeholder bool) *Component {
	p.Placeholder = placeholder

	return p
}

// 标签页Item
func (p *Component) SetItems(items []*Item) *Component {
	p.Items = items

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "tabbar"

	return p
}
