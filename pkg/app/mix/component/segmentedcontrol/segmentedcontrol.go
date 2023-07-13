package segmentedcontrol

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Current     int         `json:"current"`
	StyleType   string      `json:"styleType"`
	ActiveColor string      `json:"activeColor"`
	Titles      interface{} `json:"titles"`
	Items       interface{} `json:"items"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "segmentedControl"
	p.Current = 0
	p.StyleType = "button"
	p.ActiveColor = "#007aff"
	p.SetKey("segmentedControl", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 当前选中的tab索引值，从0计数
func (p *Component) SetCurrent(current int) *Component {
	p.Current = current
	return p
}

// 跳转方式
func (p *Component) SetStyleType(styleType string) *Component {
	p.StyleType = styleType

	return p
}

// 当 open-type 为 'navigateBack' 时有效，表示回退的层数
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口的显示/关闭动画效果，详见：窗口动画
func (p *Component) SetTitles(titles interface{}) *Component {
	p.Titles = titles

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口显示/关闭动画的持续时间。
func (p *Component) SetItems(items interface{}) *Component {
	p.Items = items

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "segmentedControl"

	return p
}
