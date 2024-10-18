package searches

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

type Select struct {
	Search
	SelectOptions []*selectfield.Option
}

// 初始化模板
func (p *Select) TemplateInit(ctx *builder.Context) interface{} {
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

// 单向联动,返回数据类型：map[string]string{"field": "you_want_load_field","api": "admin/resource_name/action/select-options"}
func (p *Select) Load(ctx *builder.Context) map[string]string {
	return nil
}
