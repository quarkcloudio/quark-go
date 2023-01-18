package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Year struct {
	Item
}

// 初始化
func (p *Year) Init() *Year {
	p.Component = "yearField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
