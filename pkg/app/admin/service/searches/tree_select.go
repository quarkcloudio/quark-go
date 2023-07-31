package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type TreeSelectField struct {
	searches.TreeSelect
}

// 树形下拉框
func TreeSelect(column string, name string, options []*treeselect.TreeData) *TreeSelectField {
	field := &TreeSelectField{}

	field.Column = column
	field.Name = name
	field.TreeSelectOptions = options

	return field
}

// 执行查询
func (p *TreeSelectField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *TreeSelectField) Options(ctx *builder.Context) interface{} {
	return p.TreeSelectOptions
}
