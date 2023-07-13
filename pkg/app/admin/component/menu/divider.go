package menu

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Divider struct {
	component.Element
	Dashed bool `json:"dashed"`
}

// 初始化
func (p *Divider) Init() *Divider {
	p.Component = "menuDivider"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 子菜单项值
func (p *Divider) SetDashed(dashed bool) *Divider {
	p.Dashed = dashed

	return p
}
