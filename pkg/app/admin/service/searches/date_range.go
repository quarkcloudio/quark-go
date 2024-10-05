package searches

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/gorm"
)

type DateRangeField struct {
	searches.DateRange
}

// 日期
func DateRange(column string, name string) *DateRangeField {
	field := &DateRangeField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *DateRangeField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	values, ok := value.([]interface{})
	if !ok {
		return query
	}

	return query.Where(p.Column+" BETWEEN ? AND ?", values[0], values[1])
}
