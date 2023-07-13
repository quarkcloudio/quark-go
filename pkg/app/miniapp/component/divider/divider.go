package divider

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Dashed          bool        `json:"dashed"`
	Hairline        bool        `json:"hairline"`
	ContentPosition string      `json:"contentPosition"`
	Direction       string      `json:"direction"`
	Body            interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "divider"
	p.SetKey("divider", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 是否使用虚线
func (p *Component) SetDashed(dashed bool) *Component {
	p.Dashed = dashed
	return p
}

// 是否使用 0.5px 线
func (p *Component) SetHairline(hairline bool) *Component {
	p.Hairline = hairline
	return p
}

// 内容位置，可选值为 left、right
func (p *Component) SetContentPosition(contentPosition string) *Component {
	p.ContentPosition = contentPosition
	return p
}

// 水平还是垂直类型
func (p *Component) SetDirection(direction string) *Component {
	p.Direction = direction
	return p
}

// 组件内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "divider"

	return p
}
