package page

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Title   string      `json:"title"`
	Navbar  interface{} `json:"navbar"`
	Content interface{} `json:"content"`
	Tabbar  interface{} `json:"tabbar"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
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
func (p *Component) SetNavbar(navbar interface{}) *Component {
	p.Navbar = navbar
	return p
}

// 内容
func (p *Component) SetContent(content interface{}) *Component {
	p.Content = content
	return p
}

// 底部导航
func (p *Component) SetTabbar(tabbar interface{}) *Component {
	p.Tabbar = tabbar
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "page"

	return p
}
