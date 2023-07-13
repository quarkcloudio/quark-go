package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type DotStyle struct {
	component.Element
	Width                   int    `json:"width"`
	Bottom                  int    `json:"bottom"`
	Color                   string `json:"color"`
	BackgroundColor         string `json:"backgroundColor"`
	Border                  string `json:"border"`
	SelectedBackgroundColor string `json:"selectedBackgroundColor"`
	SelectedBorder          string `json:"selectedBorder"`
}

// 初始化
func (p *DotStyle) Init() *DotStyle {
	p.Component = "dotStyle"
	p.SetKey("dotStyle", component.DEFAULT_CRYPT)

	return p
}

// 指示点宽度 在 mode = nav、mode = indexes 时不生效
func (p *DotStyle) SetWidth(width int) *DotStyle {
	p.Width = width

	return p
}

// 指示点距 swiper 底部的高度
func (p *DotStyle) SetBottom(bottom int) *DotStyle {
	p.Bottom = bottom

	return p
}

// 指示点前景色，只在 mode = nav ，mode = indexes 时生效
func (p *DotStyle) SetColor(color string) *DotStyle {
	p.Color = color

	return p
}

// 未选择指示点背景色
func (p *DotStyle) SetBackgroundColor(backgroundColor string) *DotStyle {
	p.BackgroundColor = backgroundColor

	return p
}

// 未选择指示点边框样式
func (p *DotStyle) SetBorder(border string) *DotStyle {
	p.Border = border

	return p
}

// 已选择指示点背景色，在 mode = nav 时不生效
func (p *DotStyle) SetSelectedBackgroundColor(selectedBackgroundColor string) *DotStyle {
	p.SelectedBackgroundColor = selectedBackgroundColor

	return p
}

// 已选择指示点边框样式，在 mode = nav 时不生效
func (p *DotStyle) SetSelectedBorder(selectedBorder string) *DotStyle {
	p.SelectedBorder = selectedBorder

	return p
}

// 组件json序列化
func (p *DotStyle) JsonSerialize() *DotStyle {
	p.Component = "dotStyle"

	return p
}
