package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type SwiperItem struct {
	component.Element
	Title string      `json:"title"`
	Body  interface{} `json:"body"`
}

// 初始化
func (p *SwiperItem) Init() *SwiperItem {
	p.Component = "swiperItem"
	p.SetKey("swiperItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *SwiperItem) SetStyle(style interface{}) *SwiperItem {
	p.Style = style

	return p
}

// 标题
func (p *SwiperItem) SetTitle(title string) *SwiperItem {
	p.Title = title

	return p
}

// 内容
func (p *SwiperItem) SetBody(body interface{}) *SwiperItem {
	p.Body = body

	return p
}

// 组件json序列化
func (p *SwiperItem) JsonSerialize() *SwiperItem {
	p.Component = "swiperItem"

	return p
}
