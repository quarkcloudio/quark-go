package collapse

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Accordion bool        `json:"accordion"`
	Body      interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "collapse"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 是否显示边框
func (p *Component) SetAccordion(accordion bool) *Component {
	p.Accordion = accordion

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "collapse"

	return p
}
