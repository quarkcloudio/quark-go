package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
	"gorm.io/gorm"
)

type RadioField struct {
	searches.Radio
	RadioOptions []radio.Option
}

// 下拉框
func Radio(column string, name string) *RadioField {
	field := &RadioField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *RadioField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *RadioField) Options(ctx *quark.Context) interface{} {
	return p.RadioOptions
}

// 设置属性，示例：[]radio.Option{{Value: 1, Label: "男"}, {Value: 2, Label: "女"}}
//
// 或者
//
// SetOptions(options, "label_name", "value_name")
func (p *RadioField) SetOptions(options ...interface{}) *RadioField {
	if len(options) == 1 {
		getOptions, ok := options[0].([]radio.Option)
		if ok {
			p.RadioOptions = getOptions
			return p
		}
	}
	if len(options) == 3 {
		p.RadioOptions = radio.New().ListToOptions(options[0], options[1].(string), options[2].(string))
	}
	return p
}
