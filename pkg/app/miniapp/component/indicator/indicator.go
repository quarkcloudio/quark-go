package indicator

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Current  int         `json:"current"`
	Size     int         `json:"size"`
	Block    bool        `json:"block"`
	Align    string      `json:"align"`
	FillZero bool        `json:"fillZero"`
	Body     interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "indicator"
	p.SetKey("indicator", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 当前步骤
func (p *Component) SetCurrent(current int) *Component {
	p.Current = current
	return p
}

// 步骤长度
func (p *Component) SetSize(size int) *Component {
	p.Size = size
	return p
}

// 是否启用块级布局
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block
	return p
}

// 对齐方式，仅在 block 为 true 时生效, 可选值 left, right, center
func (p *Component) SetAlign(align string) *Component {
	p.Align = align
	return p
}

// 对齐方式，仅在 block 为 true 时生效, 可选值 left, right, center
func (p *Component) SetFillZero(fillZero bool) *Component {
	p.FillZero = fillZero
	return p
}

// 组件内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "indicator"

	return p
}
