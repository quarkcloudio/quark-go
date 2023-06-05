package form

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/component/miniapp/component"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields"
)

type Field struct {
	component.Element
	Api             string      `json:"api"`
	Model           interface{} `json:"model"`
	Rules           interface{} `json:"rules"`
	ValidateTrigger string      `json:"validateTrigger"`
	LabelPosition   string      `json:"labelPosition"`
	LabelWidth      int         `json:"labelWidth"`
	LabelAlign      string      `json:"labelAlign"`
	ErrShowType     string      `json:"errShowType"`
	Border          bool        `json:"border"`
	Body            interface{} `json:"body"`
}

// 输入框
func (p *Field) Input(params ...interface{}) *fields.Input {
	field := (&fields.Input{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if placeholder == "" {
		field.SetPlaceholder("请输入" + params[1].(string))
	}

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 多选框
func (p *Field) Checkbox(params ...interface{}) *fields.Checkbox {
	field := (&fields.Checkbox{}).Init()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 单选框
func (p *Field) Radio(params ...interface{}) *fields.Radio {
	field := (&fields.Radio{}).Init()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 开关选择器
func (p *Field) Switch(params ...interface{}) *fields.Switch {
	field := (&fields.Switch{}).Init()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 滑动选择器
func (p *Field) Slider(params ...interface{}) *fields.Slider {
	field := (&fields.Slider{}).Init()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 从底部弹起的滚动选择器
func (p *Field) Picker(name string, label string, pickerRange []interface{}) *fields.Picker {
	field := (&fields.Picker{}).Init().SetRange(pickerRange)

	field.SetName(name).SetLabel(label)

	return field
}
