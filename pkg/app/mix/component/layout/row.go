package layout

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Row struct {
	component.Element
	Gutter interface{} `json:"gutter"`
	Col    *Col        `json:"col"`
	Body   interface{} `json:"body"`
}

// 初始化
func (p *Row) Init() *Row {
	p.Component = "row"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Row) SetStyle(style interface{}) *Row {
	p.Style = style

	return p
}

// 栅格间隔，可以写成像素值或支持响应式的对象写法来设置水平间隔 { xs: 8, sm: 16, md: 24}。
// 或者使用数组形式同时设置 [水平间距, 垂直间距]
func (p *Row) SetGutter(gutter interface{}) *Row {
	p.Gutter = gutter

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
