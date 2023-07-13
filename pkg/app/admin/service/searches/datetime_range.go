package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type DateTimeRange struct {
	searches.DatetimeRange
}

// 初始化
func (p *DateTimeRange) Init(column string, name string) *DateTimeRange {
	p.ParentInit()
	p.Column = column
	p.Name = name

	return p
}

// 执行查询
func (p *DateTimeRange) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	values, ok := value.([]interface{})
	if !ok {
		return query
	}

	return query.Where(p.Column+" BETWEEN ? AND ?", values[0], values[1])
}
