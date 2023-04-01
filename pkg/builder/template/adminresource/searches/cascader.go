package searches

import "github.com/quarkcms/quark-go/pkg/component/admin/form/fields/cascader"

type Cascader struct {
	Search
}

// 初始化
func (p *Cascader) ParentInit() interface{} {
	p.Component = "cascaderField"

	return p
}

// 设置Option
func (p *Cascader) Option(value interface{}, label string) *cascader.Option {

	return &cascader.Option{
		Value: value,
		Label: label,
	}
}
