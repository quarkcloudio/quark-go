package row

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/col"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"
)

type Component struct {
	component.Element
	Type     string           `json:"type"`
	Gutter   string           `json:"gutter"`
	Justify  string           `json:"justify"`
	Align    string           `json:"align"`
	FlexWrap string           `json:"flexWrap"`
	Body     []*col.Component `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "row"
	p.SetKey("row", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 布局方式，可选值为 flex
func (p *Component) SetType(rowType string) *Component {
	p.Type = rowType
	return p
}

// 列元素之间的间距（单位为 px）
func (p *Component) SetGutter(gutter string) *Component {
	p.Gutter = gutter
	return p
}

// Flex 主轴对齐方式，可选值为 start end center space-around space-between space-evenly
func (p *Component) SetJustify(justify string) *Component {
	p.Justify = justify
	return p
}

// Flex 交叉轴对齐方式，可选值为 flex-start center flex-end
func (p *Component) SetAlign(align string) *Component {
	p.Align = align
	return p
}

// Flex 是否换行，可选值为 nowrap wrap reverse
func (p *Component) SetFlexWrap(flexWrap string) *Component {
	p.FlexWrap = flexWrap
	return p
}

// 导航栏内容
func (p *Component) SetBody(body []*col.Component) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "row"

	return p
}
