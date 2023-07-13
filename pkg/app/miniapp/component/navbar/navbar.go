package navbar

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Title            string      `json:"title"`
	LeftText         string      `json:"leftText"`
	Desc             string      `json:"desc"`
	LeftShow         bool        `json:"leftShow"`
	TitleIcon        bool        `json:"titleIcon"`
	Border           bool        `json:"border"`
	Fixed            bool        `json:"fixed"`
	Placeholder      bool        `json:"placeholder"`
	SafeAreaInsetTop bool        `json:"safeAreaInsetTop"`
	ZIndex           int         `json:"zIndex"`
	Body             interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "navbar"
	p.SetKey("navbar", component.DEFAULT_CRYPT)

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

// 左侧文案
func (p *Component) SetLeftText(leftText string) *Component {
	p.LeftText = leftText
	return p
}

// 右侧描述
func (p *Component) SetDesc(desc string) *Component {
	p.Desc = desc
	return p
}

// 是否展示左侧箭头
func (p *Component) SetLeftShow(leftShow bool) *Component {
	p.LeftShow = leftShow
	return p
}

// 标题中是否展示 icon
func (p *Component) SetTitleIcon(titleIcon bool) *Component {
	p.TitleIcon = titleIcon
	return p
}

// 是否显示下边框
func (p *Component) SetBorder(border bool) *Component {
	p.Border = border
	return p
}

// 是否固定到顶部
func (p *Component) SetFixed(fixed bool) *Component {
	p.Fixed = fixed
	return p
}

// 固定在顶部时，是否在标签位置生成一个等高的占位元素
func (p *Component) SetPlaceholder(placeholder bool) *Component {
	p.Placeholder = placeholder
	return p
}

// 是否开启顶部安全区适配
func (p *Component) SetSafeAreaInsetTop(safeAreaInsetTop bool) *Component {
	p.SafeAreaInsetTop = safeAreaInsetTop
	return p
}

// 导航栏 z-index
func (p *Component) SetZIndex(zIndex int) *Component {
	p.ZIndex = zIndex
	return p
}

// 导航栏内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "navbar"

	return p
}
