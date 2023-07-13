package navbar

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Title           string      `json:"title"`
	LeftText        string      `json:"leftText"`
	RightText       string      `json:"rightText"`
	LeftIcon        string      `json:"leftIcon"`
	RightIcon       string      `json:"rightIcon"`
	Color           string      `json:"color"`
	BackgroundColor string      `json:"backgroundColor"`
	Fixed           bool        `json:"fixed"`
	StatusBar       bool        `json:"statusBar"`
	Shadow          bool        `json:"shadow"`
	Border          bool        `json:"border"`
	Height          interface{} `json:"height"`
	Dark            bool        `json:"dark"`
	LeftWidth       interface{} `json:"leftWidth"`
	RightWidth      interface{} `json:"rightWidth"`
	Body            interface{} `json:"body"`
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

// 左侧按钮文本
func (p *Component) SetLeftText(leftText string) *Component {
	p.LeftText = leftText

	return p
}

// 右侧按钮文本
func (p *Component) SetRightText(rightText string) *Component {
	p.RightText = rightText

	return p
}

// 左侧按钮图标（图标类型参考 Icon 图标 (opens new window)type 属性）
func (p *Component) SetLeftIcon(leftIcon string) *Component {
	p.LeftIcon = leftIcon

	return p
}

// 右侧按钮图标（图标类型参考 Icon 图标 (opens new window)type 属性）
func (p *Component) SetRightIcon(rightIcon string) *Component {
	p.RightIcon = rightIcon

	return p
}

// 图标和文字颜色
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 导航栏背景颜色
func (p *Component) SetBackgroundColor(backgroundColor string) *Component {
	p.BackgroundColor = backgroundColor

	return p
}

// 是否固定顶部
func (p *Component) SetFixed(fixed bool) *Component {
	p.Fixed = fixed

	return p
}

// 是否包含状态栏
func (p *Component) SetStatusBar(statusBar bool) *Component {
	p.StatusBar = statusBar

	return p
}

// 导航栏下是否有阴影
func (p *Component) SetShadow(shadow bool) *Component {
	p.Shadow = shadow

	return p
}

// 导航栏下是否有边框
func (p *Component) SetBorder(border bool) *Component {
	p.Border = border

	return p
}

// 导航栏高度
func (p *Component) SetHeight(height interface{}) *Component {
	p.Height = height

	return p
}

// 导航栏开启暗黑模式
func (p *Component) SetDark(dark bool) *Component {
	p.Dark = dark

	return p
}

// 导航栏左侧插槽宽度
func (p *Component) SetLeftWidth(leftWidth interface{}) *Component {
	p.LeftWidth = leftWidth

	return p
}

// 导航栏右侧插槽宽度
func (p *Component) SetRightWidth(rightWidth interface{}) *Component {
	p.RightWidth = rightWidth

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "navbar"

	return p
}
