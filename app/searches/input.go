package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
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
func (p *InputField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" LIKE ?", "%"+value.(string)+"%")
}
