package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type List struct {
	Item
	Items               interface{} `json:"items"`
	ButtonText          string      `json:"buttonText"`
	ButtonPosition      string      `json:"buttonPosition"`
	AlwaysShowItemLabel bool        `json:"alwaysShowItemLabel"`
}

// 初始化
func (p *List) Init() *List {
	p.Component = "listField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.ButtonText = "添加一行数据"
	p.ButtonPosition = "top"
	p.AlwaysShowItemLabel = true

	return p
}

// 按钮名称
func (p *List) SetButton(text string, position string) *List {
	p.ButtonText = text
	p.ButtonPosition = position

	return p
}

// 表单项
func (p *List) SetItem(callback interface{}) *List {
	getCallback := callback.(func() interface{})

	p.Items = getCallback()

	return p
}

// Item 中总是展示 label
func (p *List) SetAlwaysShowItemLabel(alwaysShowItemLabel bool) *List {
	p.AlwaysShowItemLabel = alwaysShowItemLabel

	return p
}
