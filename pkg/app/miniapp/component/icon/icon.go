package icon

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Name  string `json:"name"`
	Color string `json:"color"`
	Size  string `json:"size"`
	Tag   string `json:"tag"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "icon"
	p.SetKey("icon", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 图标名称或图片链接
func (p *Component) SetName(name string) *Component {
	p.Name = name
	return p
}

// 图标颜色
func (p *Component) SetColor(color string) *Component {
	p.Color = color
	return p
}

// 图标大小，如 20px 2em 2rem
func (p *Component) SetSize(size string) *Component {
	p.Size = size
	return p
}

// 小程序标签
func (p *Component) SetTag(tag string) *Component {
	p.Tag = tag
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "icon"

	return p
}
