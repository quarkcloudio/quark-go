package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type GridItem struct {
	component.Element
	Index int         `json:"index"`
	Body  interface{} `json:"body"`
}

// 初始化
func (p *GridItem) Init() *GridItem {
	p.Component = "gridItem"
	p.SetKey("gridItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *GridItem) SetStyle(style interface{}) *GridItem {
	p.Style = style

	return p
}

// 每列显示个数
func (p *GridItem) SetIndex(index int) *GridItem {
	p.Index = index

	return p
}

// 内容
func (p *GridItem) SetBody(body interface{}) *GridItem {
	p.Body = body

	return p
}

// 组件json序列化
func (p *GridItem) JsonSerialize() *GridItem {
	p.Component = "gridItem"

	return p
}
