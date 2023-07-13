package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Item struct {
	component.Element
	Text string      `json:"text"`
	Body interface{} `json:"body"`
}

// 初始化组件
func NewItem() *Item {
	return (&Item{}).Init()
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "gridItem"
	p.SetKey("gridItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Item) SetStyle(style interface{}) *Item {
	p.Style = style

	return p
}

// 文字
func (p *Item) SetText(text string) *Item {
	p.Text = text
	return p
}

// 内容
func (p *Item) SetBody(body interface{}) *Item {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Item) JsonSerialize() *Item {
	p.Component = "gridItem"

	return p
}
