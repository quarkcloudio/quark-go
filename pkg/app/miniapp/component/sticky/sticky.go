package divider

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Position     string      `json:"position"`
	Top          int         `json:"top"`
	ZIndex       int         `json:"zIndex"`
	ParentHeight int         `json:"parentHeight"`
	Body         interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "sticky"
	p.SetKey("sticky", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 吸附位置（top、bottom）
func (p *Component) SetPosition(position string) *Component {
	p.Position = position
	return p
}

// 吸顶距离
func (p *Component) SetTop(top int) *Component {
	p.Top = top
	return p
}

// 吸附时的层级
func (p *Component) SetZIndex(zIndex int) *Component {
	p.ZIndex = zIndex
	return p
}

// 设置粘性元素父级高度
func (p *Component) SetParentHeight(parentHeight int) *Component {
	p.ParentHeight = parentHeight
	return p
}

// 组件内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "sticky"

	return p
}
