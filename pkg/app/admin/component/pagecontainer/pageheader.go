package pagecontainer

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type PageHeader struct {
	component.Element
	Avatar           interface{} `json:"avatar"`
	BackIcon         interface{} `json:"backIcon"`
	Breadcrumb       interface{} `json:"breadcrumb"`
	BreadcrumbRender interface{} `json:"breadcrumbRender"`
	Extra            interface{} `json:"extra"`
	Footer           interface{} `json:"footer"`
	Ghost            bool        `json:"ghost"`
	SubTitle         string      `json:"subTitle"`
	Tags             interface{} `json:"tags"`
	Title            string      `json:"title"`
}

// 初始化
func (p *PageHeader) Init() *PageHeader {
	p.Component = "pageHeader"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *PageHeader) SetStyle(style map[string]interface{}) *PageHeader {
	p.Style = style

	return p
}

// 标题栏旁的头像
func (p *PageHeader) SetAvatar(avatar interface{}) *PageHeader {
	p.Avatar = avatar
	return p
}

// 自定义 back icon ，如果为 false 不渲染 back icon
func (p *PageHeader) SetBackIcon(backIcon interface{}) *PageHeader {
	p.BackIcon = backIcon
	return p
}

// 面包屑的配置
func (p *PageHeader) SetBreadcrumb(breadcrumb interface{}) *PageHeader {
	p.Breadcrumb = breadcrumb
	return p
}

// 自定义面包屑区域的内容
func (p *PageHeader) SetBreadcrumbRender(breadcrumbRender interface{}) *PageHeader {
	p.BreadcrumbRender = breadcrumbRender
	return p
}

// 操作区，位于 title 行的行尾
func (p *PageHeader) SetExtra(extra interface{}) *PageHeader {
	p.Extra = extra
	return p
}

// PageHeader 的页脚，一般用于渲染 TabBar
func (p *PageHeader) SetFooter(footer interface{}) *PageHeader {
	p.Footer = footer
	return p
}

// pageHeader 的类型，将会改变背景颜色
func (p *PageHeader) SetGhost(ghost bool) *PageHeader {
	p.Ghost = ghost
	return p
}

// 自定义的二级标题文字
func (p *PageHeader) SetSubTitle(subTitle string) *PageHeader {
	p.SubTitle = subTitle
	return p
}

// title 旁的 tag 列表
func (p *PageHeader) SetTags(tags interface{}) *PageHeader {
	p.Tags = tags
	return p
}

// 自定义标题文字
func (p *PageHeader) SetTitle(title string) *PageHeader {
	p.Title = title
	return p
}

// 组件json序列化
func (p *PageHeader) JsonSerialize() *PageHeader {
	p.Component = "pageHeader"

	return p
}
