package tabs

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Centered           bool        `json:"centered"`
	DefaultActiveKey   string      `json:"defaultActiveKey"`
	Size               string      `json:"size"`
	TabBarExtraContent interface{} `json:"tabBarExtraContent"`
	TabBarGutter       int         `json:"tabBarGutter"`
	TabBarStyle        interface{} `json:"tabBarStyle"`
	TabPosition        string      `json:"tabPosition"`
	Type               string      `json:"type"`
	TabPanes           interface{} `json:"tabPanes"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取TabPane
func NewTabPane() *TabPane {

	return (&TabPane{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "tabs"
	p.Size = "default"
	p.TabPosition = "top"
	p.Type = "line"
	p.TabBarGutter = 35
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 标签居中展示
func (p *Component) SetCentered(centered bool) *Component {
	p.Centered = centered

	return p
}

// 初始化选中面板的 key，如果没有设置 activeKey
func (p *Component) SetDefaultActiveKey(defaultActiveKey string) *Component {
	p.DefaultActiveKey = defaultActiveKey

	return p
}

// 大小，提供 large default 和 small 三种大小
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// tab bar 上额外的元素
func (p *Component) SetTabBarExtraContent(tabBarExtraContent interface{}) *Component {
	p.TabBarExtraContent = tabBarExtraContent

	return p
}

// tabs 之间的间隙
func (p *Component) SetTabBarGutter(tabBarGutter int) *Component {
	p.TabBarGutter = tabBarGutter

	return p
}

// tab bar 的样式对象
func (p *Component) SetTabBarStyle(style map[string]interface{}) *Component {
	p.TabBarStyle = style

	return p
}

// 页签位置，可选值有 top right bottom left
func (p *Component) SetTabPosition(tabPosition string) *Component {
	p.TabPosition = tabPosition

	return p
}

// 页签的基本样式，可选 line、card editable-card 类型
func (p *Component) SetType(tabType string) *Component {
	p.Type = tabType

	return p
}

// tab 的内容
func (p *Component) SetTabPanes(tabPanes interface{}) *Component {
	p.TabPanes = tabPanes

	return p
}
