package icon

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Size         int    `json:"size"`
	Type         string `json:"type"`
	Color        string `json:"color"`
	CustomPrefix string `json:"customPrefix"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "icon"
	p.SetKey("icon", component.DEFAULT_CRYPT)
	p.Size = 24

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetSize(size int) *Component {
	p.Size = size
	return p
}

// 子标题
func (p *Component) SetType(fontType string) *Component {
	p.Type = fontType

	return p
}

// 额外信息
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 略缩图
func (p *Component) SetCustomPrefix(customPrefix string) *Component {
	p.CustomPrefix = customPrefix

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "icon"

	return p
}
