package searches

import "github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selectfield"

type Select struct {
	Search
}

// 初始化
func (p *Select) ParentInit() interface{} {
	p.Component = "selectField"

	return p
}

// 设置Option
func (p *Select) Option(value interface{}, label string) *selectfield.Option {

	return &selectfield.Option{
		Value: value,
		Label: label,
	}
}
