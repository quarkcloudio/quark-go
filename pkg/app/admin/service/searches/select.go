package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type SelectField struct {
	searches.Select
}

// 下拉框
func Select(column string, name string, options []*selectfield.Option) *SelectField {
	field := &SelectField{}

	field.Column = column
	field.Name = name
	field.SelectOptions = options

	return field
}

// 执行查询
func (p *SelectField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *SelectField) Options(ctx *builder.Context) interface{} {
	return p.SelectOptions
}
