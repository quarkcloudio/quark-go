package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Item struct {
	component.Element
	Body interface{} `json:"body"`
}

// 初始化组件
func NewItem() *Item {
	return (&Item{}).Init()
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "swiperItem"
	p.SetKey("swiperItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Item) SetStyle(style interface{}) *Item {
	p.Style = style

	return p
}

// 内容
func (p *Item) SetBody(body interface{}) *Item {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Item) JsonSerialize() *Item {
	p.Component = "swiperItem"

	return p
}
