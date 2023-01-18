package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Hidden struct {
	Item
}

// 初始化
func (p *Hidden) Init() *Hidden {
	p.Component = "hiddenField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
