package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
	"gorm.io/gorm"
)

type DateField struct {
	searches.DateRange
}

// 日期范围
func Date(column string, name string) *DateField {
	field := &DateField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *DateField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}
