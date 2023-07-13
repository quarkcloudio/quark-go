package list

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Border bool        `json:"border"`
	Body   interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "list"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 是否显示边框
func (p *Component) SetBorder(border bool) *Component {
	p.Border = border

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "list"

	return p
}
