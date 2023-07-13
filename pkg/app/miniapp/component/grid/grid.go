package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	ColumnNum int     `json:"columnNum"`
	Border    bool    `json:"border"`
	Gutter    int     `json:"gutter"`
	Center    bool    `json:"center"`
	Square    bool    `json:"square"`
	Reverse   bool    `json:"reverse"`
	Direction string  `json:"direction"`
	Clickable bool    `json:"clickable"`
	Body      []*Item `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "grid"
	p.SetKey("grid", component.DEFAULT_CRYPT)
	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style
	return p
}

// 列数
func (p *Component) SetColumnNum(columnNum int) *Component {
	p.ColumnNum = columnNum
	return p
}

// 是否显示边框
func (p *Component) SetBorder(border bool) *Component {
	p.Border = border
	return p
}

// 格子之间的间距，默认单位为 px
func (p *Component) SetGutter(gutter int) *Component {
	p.Gutter = gutter
	return p
}

// 是否将格子内容居中显示
func (p *Component) SetCenter(center bool) *Component {
	p.Center = center
	return p
}

// 是否将格子固定为正方形
func (p *Component) SetSquare(square bool) *Component {
	p.Square = square
	return p
}

// 内容翻转
func (p *Component) SetReverse(reverse bool) *Component {
	p.Reverse = reverse
	return p
}

// 格子内容排列的方向，可选值为 horizontal
func (p *Component) SetDirection(direction string) *Component {
	p.Direction = direction
	return p
}

// 是否开启格子点击反馈
func (p *Component) SetClickable(clickable bool) *Component {
	p.Clickable = clickable
	return p
}

// 内容
func (p *Component) SetBody(body []*Item) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "grid"
	return p
}
