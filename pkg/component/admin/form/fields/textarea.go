package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type TextArea struct {
	Item
	MinRows int `json:"-"`
	MaxRows int `json:"-"`
}

// 初始化
func (p *TextArea) Init() *TextArea {
	p.Component = "textAreaField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(400)
	p.AutoSize = map[string]int{
		"minRows": 2,
		"maxRows": 5,
	}

	return p
}

// autoSize 属性适用于 textarea 节点，并且只有高度会自动变化。另外 autoSize 可以设定为一个对象，指定最小行数和最大行数。
func (p *TextArea) SetAutoSize(autoSize interface{}) *TextArea {
	p.AutoSize = autoSize
	return p
}

// 指定最小行数。
func (p *TextArea) SetMinRows(rows int) *TextArea {
	p.MinRows = rows
	autoSize := map[string]int{
		"minRows": rows,
		"maxRows": p.MaxRows,
	}
	p.AutoSize = autoSize

	return p
}

// 指定最大行数。
func (p *TextArea) SetMaxRows(rows int) *TextArea {
	p.MaxRows = rows
	autoSize := map[string]int{
		"minRows": p.MinRows,
		"maxRows": rows,
	}
	p.AutoSize = autoSize

	return p
}
