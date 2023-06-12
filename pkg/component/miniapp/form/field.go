package form

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/calendar"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/cascader"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/checkbox"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/input"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/picker"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/radio"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/slider"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/switchfield"
)

type Field struct{}

// 日历
func (p *Field) Calendar(params ...interface{}) *calendar.Component {
	field := calendar.New()

	field.
		SetName(params[0].(string)).
		SetLabel(params[1].(string))

	return field
}

// 级联选择器
func (p *Field) Cascader(params ...interface{}) *cascader.Component {
	field := cascader.New()

	field.
		SetName(params[0].(string)).
		SetLabel(params[1].(string))

	return field
}

// 输入框
func (p *Field) Input(params ...interface{}) *input.Component {
	field := input.New()

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
func (p *Field) Checkbox(params ...interface{}) *checkbox.Component {
	field := checkbox.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 单选框
func (p *Field) Radio(params ...interface{}) *radio.Component {
	field := radio.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 开关选择器
func (p *Field) Switch(params ...interface{}) *switchfield.Component {
	field := switchfield.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 滑动选择器
func (p *Field) Slider(params ...interface{}) *slider.Component {
	field := slider.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 从底部弹起的滚动选择器
func (p *Field) Picker(name string, label string, pickerRange []interface{}) *picker.Component {
	field := picker.New().SetRange(pickerRange)

	field.SetName(name).SetLabel(label)

	return field
}
