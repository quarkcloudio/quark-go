package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/cascader"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
	"gorm.io/gorm"
)

type CascaderField struct {
	searches.Cascader
	CascaderOptions []cascader.Option
}

// 下拉框
func Cascader(column string, name string) *CascaderField {
	field := &CascaderField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *CascaderField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" IN ?", value)
}

// 属性
func (p *CascaderField) Options(ctx *quark.Context) interface{} {
	return p.CascaderOptions
}

// 可选项数据源
//
// SetOptions([]cascader.Option {{Value :"zhejiang", Label:"Zhejiang"}})
//
// 或者
//
// SetOptions(options, "parent_key_name", "label_name", "value_name")
//
// 或者
//
// SetOptions(options, 0, "parent_key_name", "label_name", "value_name")
func (p *CascaderField) SetOptions(options ...interface{}) *CascaderField {
	if len(options) == 1 {
		getOptions, ok := options[0].([]cascader.Option)
		if ok {
			p.CascaderOptions = getOptions
			return p
		}
	}
	if len(options) == 4 {
		p.CascaderOptions = cascader.New().ListToOptions(options[0], 0, options[1].(string), options[2].(string), options[3].(string))
	}
	if len(options) == 5 {
		p.CascaderOptions = cascader.New().ListToOptions(options[0], options[1].(int), options[2].(string), options[3].(string), options[4].(string))
	}
	return p
}
