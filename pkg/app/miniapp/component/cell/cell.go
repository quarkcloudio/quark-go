package cell

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Title         string `json:"title"`
	SubTitle      string `json:"subTitle"`
	Desc          string `json:"desc"`
	DescTextAlign string `json:"descTextAlign"`
	IsLink        bool   `json:"isLink"`
	Url           string `json:"url"`
	To            string `json:"to"`
	RoundRadius   int    `json:"roundRadius"`
	Center        bool   `json:"center"`
	Size          string `json:"size"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "cell"
	p.SetKey("cell", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标题名称
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 左侧副标题
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle
	return p
}

// 右侧描述
func (p *Component) SetDesc(desc string) *Component {
	p.Desc = desc
	return p
}

// 右侧描述文本对齐方式 text-align
func (p *Component) SetDescTextAlign(descTextAlign string) *Component {
	p.DescTextAlign = descTextAlign
	return p
}

// 是否展示右侧箭头并开启点击反馈
func (p *Component) SetIsLink(isLink bool) *Component {
	p.IsLink = isLink
	return p
}

// 标签页的跳转链接；如果同时存在 to，优先级高于 to
func (p *Component) SetHref(url string) *Component {
	p.Url = url
	return p
}

// 标签页的路由对象，等于 vue-router 的 to 属性 属性
func (p *Component) SetTo(to string) *Component {
	p.To = to
	return p
}

// 圆角半径
func (p *Component) SetRoundRadius(roundRadius int) *Component {
	p.RoundRadius = roundRadius
	return p
}

// 是否使内容垂直居中
func (p *Component) SetCenter(center bool) *Component {
	p.Center = center
	return p
}

// 单元格大小，可选值为 large
func (p *Component) SetSize(size string) *Component {
	p.Size = size
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "cell"

	return p
}
