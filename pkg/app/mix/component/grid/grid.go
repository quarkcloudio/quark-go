package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Column      int         `json:"column"`
	BorderColor string      `json:"borderColor"`
	ShowBorder  bool        `json:"showBorder"`
	Square      bool        `json:"square"`
	Highlight   bool        `json:"highlight"`
	Body        interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "grid"
	p.SetKey("grid", component.DEFAULT_CRYPT)
	p.Column = 3

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 每列显示个数
func (p *Component) SetColumn(column int) *Component {
	p.Column = column

	return p
}

// 边框颜色
func (p *Component) SetBorderColor(borderColor string) *Component {
	p.BorderColor = borderColor
	return p
}

// 是否显示边框
func (p *Component) SetShowBorder(showBorder bool) *Component {
	p.ShowBorder = showBorder

	return p
}

// 是否方形显示
func (p *Component) SetSquare(square bool) *Component {
	p.Square = square

	return p
}

// 点击背景是否高亮
func (p *Component) SetHighlight(highlight bool) *Component {
	p.Highlight = highlight

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "grid"

	return p
}
