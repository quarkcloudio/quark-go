package view

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Body interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "view"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 容器控件里面的内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}
