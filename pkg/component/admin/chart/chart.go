package chart

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Component struct {
	component.Element
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "chart"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 折线图表
func (p *Component) Line(data interface{}) *Line {

	return (&Line{}).Init().SetData(data)
}
