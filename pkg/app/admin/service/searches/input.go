package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/searches"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type InputField struct {
	searches.Search
}

// 输入框
func Input(column string, name string) *InputField {
	field := &InputField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *InputField) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" LIKE ?", "%"+value.(string)+"%")
}
