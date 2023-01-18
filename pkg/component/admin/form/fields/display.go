package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Display struct {
	Item
}

// 初始化
func (p *Display) Init() *Display {
	p.Component = "displayField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
