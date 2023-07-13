package tabbar

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Item struct {
	component.Element
	TabTitle string `json:"tabTitle"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Href     string `json:"href"`
	To       string `json:"to"`
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "tabbarItem"
	p.SetKey("tabbarItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Item) SetStyle(style interface{}) *Item {
	p.Style = style

	return p
}

// 标签页的标题
func (p *Item) SetTabTitle(tabTitle string) *Item {
	p.TabTitle = tabTitle

	return p
}

// 标签名称，作为匹配的标识符
func (p *Item) SetName(name string) *Item {
	p.Name = name

	return p
}

// 标签页显示的图标名称
func (p *Item) SetIcon(icon string) *Item {
	p.Icon = icon

	return p
}

// 标签页的跳转链接；如果同时存在 to，优先级高于 to
func (p *Item) SetHref(href string) *Item {
	p.Href = href

	return p
}

// 标签页的路由对象，等于 vue-router 的 to 属性 属性
func (p *Item) SetTo(to string) *Item {
	p.To = to

	return p
}

// 组件json序列化
func (p *Item) JsonSerialize() *Item {
	p.Component = "tabbarItem"

	return p
}
