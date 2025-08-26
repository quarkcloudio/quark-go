package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
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
func (p *DateRangeField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	values, ok := value.([]interface{})
	if !ok {
		return query
	}

	return query.Where(p.Column+" BETWEEN ? AND ?", values[0], values[1])
}
