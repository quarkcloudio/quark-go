package navbar

import "github.com/quarkcms/quark-go/pkg/component/miniapp/component"

type Component struct {
	component.Element
	Title            string `json:"title"`
	LeftText         string `json:"leftText"`
	Desc             string `json:"desc"`
	LeftShow         bool   `json:"leftShow"`
	TitleIcon        bool   `json:"titleIcon"`
	Border           bool   `json:"border"`
	Fixed            bool   `json:"fixed"`
	Placeholder      bool   `json:"placeholder"`
	SafeAreaInsetTop bool   `json:"safeAreaInsetTop"`
	ZIndex           int    `json:"zIndex"`
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

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "navbar"

	return p
}
