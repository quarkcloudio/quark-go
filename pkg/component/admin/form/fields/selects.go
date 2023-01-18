package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Selects struct {
	Item
	Body interface{} `json:"body"`
}

// 初始化
func (p *Selects) Init() *Selects {
	p.Component = "selects"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 组件内容
func (p *Selects) SetBody(body interface{}) *Selects {
	p.Body = body

	return p
}
