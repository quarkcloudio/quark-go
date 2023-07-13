package col

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Span   int         `json:"span"`
	Offset int         `json:"offset"`
	Body   interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "col"
	p.SetKey("col", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetSpan(span int) *Component {
	p.Span = span
	return p
}

// 左侧文案
func (p *Component) SetOffset(offset int) *Component {
	p.Offset = offset
	return p
}

// 导航栏内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "col"

	return p
}
