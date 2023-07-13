package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Row struct {
	component.Element
	Align   string      `json:"align"`
	Gutter  interface{} `json:"gutter"`
	Justify string      `json:"justify"`
	Wrap    bool        `json:"wrap"`
	Col     *Col        `json:"col"`
	Body    interface{} `json:"body"`
}

// 获取Row
func (p *Component) Row() *Row {
	return (&Row{}).Init()
}

// 初始化
func (p *Row) Init() *Row {
	p.Component = "row"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Row) SetStyle(style map[string]interface{}) *Row {
	p.Style = style

	return p
}

// 垂直对齐方式
func (p *Row) SetAlign(align string) *Row {
	p.Align = align

	return p
}

// 栅格间隔，可以写成像素值或支持响应式的对象写法来设置水平间隔 { xs: 8, sm: 16, md: 24}。
// 或者使用数组形式同时设置 [水平间距, 垂直间距]
func (p *Row) SetGutter(gutter interface{}) *Row {
	p.Gutter = gutter

	return p
}

// 水平排列方式，start | end | center | space-around | space-between
func (p *Row) SetJustify(justify string) *Row {
	p.Justify = justify

	return p
}

// 是否自动换行
func (p *Row) SetWrap(wrap bool) *Row {
	p.Wrap = wrap

	return p
}

// 设置列
func (p *Row) SetCol(col *Col) *Row {
	p.Col = col

	return p
}

// 内容
func (p *Row) SetBody(body interface{}) *Row {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Row) JsonSerialize() *Row {
	p.Component = "row"

	return p
}
