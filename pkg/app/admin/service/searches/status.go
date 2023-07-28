package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type StatusField struct {
	searches.Select
}

// 状态
func Status() *StatusField {
	field := &StatusField{}
	field.Name = "状态"

	return field
}

// 执行查询
func (p *StatusField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where("status = ?", value)
}

// 属性
func (p *StatusField) Options(ctx *builder.Context) interface{} {

	return []*selectfield.Option{
		p.Option(0, "禁用"),
		p.Option(1, "正常"),
	}
}
