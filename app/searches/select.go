package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
	"gorm.io/gorm"
)

type SelectField struct {
	searches.Select
	SelectOptions []selectfield.Option
}

// 下拉框
func Select(column string, name string) *SelectField {
	field := &SelectField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *SelectField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *SelectField) Options(ctx *quark.Context) interface{} {
	return p.SelectOptions
}

//	[]selectfield.Option{
//			{Value: 1, Label: "新闻"},
//			{Value: 2, Label: "音乐"},
//			{Value: 3, Label: "体育"},
//		}
//
// 或者
//
// SetOptions(options, "label_name", "value_name")
func (p *SelectField) SetOptions(options ...interface{}) *SelectField {
	if len(options) == 1 {
		getOptions, ok := options[0].([]selectfield.Option)
		if ok {
			p.SelectOptions = getOptions
			return p
		}
	}
	if len(options) == 3 {
		p.SelectOptions = selectfield.New().ListToOptions(options[0], options[1].(string), options[2].(string))
	}
	return p
}
