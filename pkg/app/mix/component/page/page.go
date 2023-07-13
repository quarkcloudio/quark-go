package page

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Title   string      `json:"title"`
	NavBar  interface{} `json:"navBar"`
	Content interface{} `json:"content"`
	TabBar  interface{} `json:"tabBar"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "page"
	p.SetKey("page", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 头部导航
func (p *Component) SetNavBar(navBar interface{}) *Component {
	p.NavBar = navBar
	return p
}

// 内容
func (p *Component) SetContent(content interface{}) *Component {
	p.Content = content
	return p
}

// 底部导航
func (p *Component) SetTabBar(tabBar interface{}) *Component {
	p.TabBar = tabBar
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "page"

	return p
}
