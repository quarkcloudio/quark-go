package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type SwiperDot struct {
	component.Element
	Current    int         `json:"current"`
	Mode       string      `json:"mode"`
	Field      string      `json:"field"`
	Items      interface{} `json:"items"`
	DotsStyles interface{} `json:"dotsStyles"`
}

// 初始化
func (p *SwiperDot) Init() *SwiperDot {
	p.Component = "swiperDot"
	p.SetKey("swiperDot", component.DEFAULT_CRYPT)
	p.DotsStyles = nil

	return p
}

// Set style.
func (p *SwiperDot) SetStyle(style interface{}) *SwiperDot {
	p.Style = style

	return p
}

// 当前指示点索引，必须是通过 swiper 的 change 事件获取到的 e.detail.current
func (p *SwiperDot) SetCurrent(current int) *SwiperDot {
	p.Current = current

	return p
}

// 指示点的类型，可选值：default 、round 、nav 、 indexes
func (p *SwiperDot) SetMode(mode string) *SwiperDot {
	p.Mode = mode
	return p
}

// mode 为 nav 时，显示的内容字段（mode = nav 时必填）
func (p *SwiperDot) SetField(field string) *SwiperDot {
	p.Field = field

	return p
}

// 轮播图的数据，通过数组长度决定指示点个数
func (p *SwiperDot) SetItems(items interface{}) *SwiperDot {
	p.Items = items

	return p
}

// 轮播图的数据，通过数组长度决定指示点个数
func (p *SwiperDot) SetBody(body interface{}) *SwiperDot {
	p.Items = body

	return p
}

// 指示点样式
func (p *SwiperDot) SetDotsStyles(dotsStyles interface{}) *SwiperDot {
	p.DotsStyles = dotsStyles

	return p
}

// 组件json序列化
func (p *SwiperDot) JsonSerialize() *SwiperDot {
	p.Component = "swiperDot"

	return p
}
