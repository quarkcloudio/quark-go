package searches

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/radio"
)

type Radio struct {
	Search
	RadioOptions []radio.Option
}

// 加载初始化数据
func (p *Radio) New(ctx *quark.Context) interface{} {
	p.Component = "radioField"
	return p
}

// 设置Option
func (p *Radio) Option(label string, value interface{}) radio.Option {

	return radio.Option{
		Value: value,
		Label: label,
	}
}

// 设置属性，示例：[]radio.Option{{Value: 1, Label: "男"}, {Value: 2, Label: "女"}}
//
// 或者
//
// SetOptions(options, "label_name", "value_name")
func (p *Radio) SetOptions(options ...interface{}) *Radio {
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
