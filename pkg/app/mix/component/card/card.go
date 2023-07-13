package card

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Title     string      `json:"title"`
	SubTitle  string      `json:"subTitle"`
	Extra     string      `json:"extra"`
	Thumbnail string      `json:"thumbnail"`
	Cover     string      `json:"cover"`
	IsFull    bool        `json:"isFull"`
	IsShadow  bool        `json:"isShadow"`
	Shadow    string      `json:"shadow"`
	Border    bool        `json:"border"`
	Margin    string      `json:"margin"`
	Spacing   string      `json:"spacing"`
	Padding   string      `json:"padding"`
	Body      interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "card"
	p.SetKey("card", component.DEFAULT_CRYPT)
	p.Shadow = "0px 0px 3px 1px rgba(0, 0, 0, 0.08)"
	p.Border = true
	p.Margin = "10px"
	p.Spacing = "10px"
	p.Padding = "10px"

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 子标题
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle

	return p
}

// 额外信息
func (p *Component) SetExtra(extra string) *Component {
	p.Extra = extra

	return p
}

// 略缩图
func (p *Component) SetThumbnail(thumbnail string) *Component {
	p.Thumbnail = thumbnail

	return p
}

// 封面图
func (p *Component) SetCover(cover string) *Component {
	p.Cover = cover

	return p
}

// 卡片内容是否通栏，为true时将去除padding值
func (p *Component) SetIsFull(isFull bool) *Component {
	p.IsFull = isFull

	return p
}

// 卡片阴影,需符合 css 值,0px 0px 3px 1px rgba(0, 0, 0, 0.08)
func (p *Component) SetShadow(shadow string) *Component {
	p.Shadow = shadow

	return p
}

// 卡片边框
func (p *Component) SetBorder(border bool) *Component {
	p.Border = border

	return p
}

// 卡片外边距
func (p *Component) SetMargin(margin string) *Component {
	p.Margin = margin

	return p
}

// 卡片内边距
func (p *Component) SetSpacing(spacing string) *Component {
	p.Spacing = spacing

	return p
}

// 卡片内容内边距
func (p *Component) SetPadding(padding string) *Component {
	p.Padding = padding

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "card"

	return p
}
