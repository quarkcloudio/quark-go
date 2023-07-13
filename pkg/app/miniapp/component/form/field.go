package form

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/calendar"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/cascader"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/checkbox"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/input"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/picker"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/radio"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/switchfield"
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

// 数字输入框
func (p *Field) InputNumber(params ...interface{}) *input.Component {
	field := input.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 多选框
func (p *Field) Checkbox(params ...interface{}) *checkbox.Component {
	field := checkbox.New()

	field.SetName(params[0].(string)).SetLabel(params[1].(string))

	return field
}

// 从底部弹起的滚动选择器
func (p *Field) Picker(name string, label string, pickerRange []interface{}) *picker.Component {
	field := picker.New()

	field.SetName(name).SetLabel(label)

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
