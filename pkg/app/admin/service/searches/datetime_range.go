package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type DateTimeRangeField struct {
	searches.DatetimeRange
}

// 日期时间范围
func DatetimeRange(column string, name string) *DateTimeRangeField {
	field := &DateTimeRangeField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *DateTimeRangeField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	values, ok := value.([]interface{})
	if !ok {
		return query
	}

	return query.Where(p.Column+" BETWEEN ? AND ?", values[0], values[1])
}
