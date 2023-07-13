package section

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Type             string      `json:"type"`
	Title            string      `json:"title"`
	TitleFontSize    string      `json:"titleFontSize"`
	TitleColor       string      `json:"titleColor"`
	SubTitle         string      `json:"subTitle"`
	SubTitleFontSize string      `json:"subTitleFontSize"`
	SubTitleColor    string      `json:"subTitleColor"`
	Padding          interface{} `json:"padding"`
	Body             interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "section"
	p.SetKey("section", component.DEFAULT_CRYPT)
	p.TitleFontSize = "14px"
	p.TitleColor = "#333"
	p.SubTitleFontSize = "12px"
	p.SubTitleColor = "#999"

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 装饰类型，可选值：line（竖线）、circle（圆形）、square（方形）
func (p *Component) SetType(sectionType string) *Component {
	p.Type = sectionType

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 主标题字体大小
func (p *Component) SetTitleFontSize(titleFontSize string) *Component {
	p.TitleFontSize = titleFontSize

	return p
}

// 主标题字体颜色
func (p *Component) SetTitleColor(titleColor string) *Component {
	p.TitleColor = titleColor

	return p
}

// 副标题
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle

	return p
}

// 副标题字体大小
func (p *Component) SetSubTitleFontSize(subTitleFontSize string) *Component {
	p.SubTitleFontSize = subTitleFontSize

	return p
}

// 副标题字体颜色
func (p *Component) SetSubTitleColor(subTitleColor string) *Component {
	p.SubTitleColor = subTitleColor

	return p
}

// 卡片边框
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
	p.Component = "section"

	return p
}
