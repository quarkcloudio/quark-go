package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Group struct {
	Item
	Body interface{} `json:"body"`
	Size int         `json:"size"`
}

// 初始化
func (p *Group) Init() *Group {
	p.Component = "groupField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Size = 32
	p.OnlyOnForms()

	return p
}

// 组件内容
func (p *Group) SetBody(body interface{}) *Group {
	p.Body = body

	return p
}

// 子元素个数
func (p *Group) SetSize(size int) *Group {
	p.Size = size

	return p
}
