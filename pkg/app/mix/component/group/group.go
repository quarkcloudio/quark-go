package group

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Title string      `json:"title"`
	Top   int         `json:"top"`
	Mode  string      `json:"mode"`
	Body  interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "group"
	p.SetKey("group", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 主标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 分组间隔
func (p *Component) SetTop(top int) *Component {
	p.Top = top

	return p
}

// 模式 ，card 为卡片模式
func (p *Component) SetMode(mode string) *Component {
	p.Mode = mode

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "group"

	return p
}
