package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields"
)

// 后台字段组件
type Field struct{}

// ID组件
func (p *Field) ID(params ...interface{}) *fields.ID {
	field := (&fields.ID{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// Hidden组件
func (p *Field) Hidden(params ...interface{}) *fields.Hidden {
	field := (&fields.Hidden{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 输入框组件
func (p *Field) Text(params ...interface{}) *fields.Text {
	field := (&fields.Text{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 文本域组件
func (p *Field) TextArea(params ...interface{}) *fields.TextArea {
	field := (&fields.TextArea{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 密码框组件
func (p *Field) Password(params ...interface{}) *fields.Password {
	field := (&fields.Password{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 单选组件
func (p *Field) Radio(params ...string) *fields.Radio {
	field := &fields.Radio{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 多选组件
func (p *Field) Checkbox(params ...string) *fields.Checkbox {
	field := &fields.Checkbox{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 日期组件
func (p *Field) Date(params ...interface{}) *fields.Date {
	field := &fields.Date{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 日期范围组件
func (p *Field) DateRange(params ...interface{}) *fields.DateRange {
	field := &fields.DateRange{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 日期时间组件
func (p *Field) Datetime(params ...interface{}) *fields.Datetime {
	field := &fields.Datetime{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 日期时间范围组件
func (p *Field) DatetimeRange(params ...interface{}) *fields.DatetimeRange {
	field := &fields.DatetimeRange{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 开关组件
func (p *Field) Switch(params ...string) *fields.Switch {
	field := &fields.Switch{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 树形组件
func (p *Field) Tree(params ...interface{}) *fields.Tree {
	field := (&fields.Tree{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 图标组件
func (p *Field) Icon(params ...interface{}) *fields.Icon {
	field := (&fields.Icon{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 下拉框组件
func (p *Field) Select(params ...interface{}) *fields.Select {
	field := (&fields.Select{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 级联菜单组件
func (p *Field) Cascader(params ...interface{}) *fields.Cascader {
	field := (&fields.Cascader{}).Init()

	if len(params) >= 2 {
		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 图片组件
func (p *Field) Image(params ...interface{}) *fields.Image {
	field := (&fields.Image{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 文件组件
func (p *Field) File(params ...interface{}) *fields.File {
	field := (&fields.File{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 文本展示组件
func (p *Field) Display(label string) *fields.Display {
	field := (&fields.Display{}).Init()
	field.SetLabel(label)

	return field
}

// 编辑器组件
func (p *Field) Editor(params ...interface{}) *fields.Editor {
	field := (&fields.Editor{}).Init()

	if len(params) >= 2 {
		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 分组组件
func (p *Field) Group(label string, items []interface{}) *fields.Group {
	field := (&fields.Group{}).Init()

	field.SetBody(items).SetLabel(label)

	return field
}

// List组件
func (p *Field) List(params ...interface{}) *fields.List {
	field := (&fields.List{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 地图组件
func (p *Field) Map(params ...interface{}) *fields.Map {
	field := (&fields.Map{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 地图围栏组件
func (p *Field) Geofence(params ...interface{}) *fields.Geofence {
	field := (&fields.Geofence{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 日期-月组件
func (p *Field) Month(params ...interface{}) *fields.Month {
	field := &fields.Month{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 数组输入框组件
func (p *Field) Number(params ...interface{}) *fields.Number {
	field := (&fields.Number{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 日期-季度组件
func (p *Field) Quarter(params ...interface{}) *fields.Quarter {
	field := &fields.Quarter{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 搜索组件
func (p *Field) Search(params ...interface{}) *fields.Search {
	field := (&fields.Search{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 时间范围组件
func (p *Field) TimeRange(params ...interface{}) *fields.TimeRange {
	field := &fields.TimeRange{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 时间组件
func (p *Field) Time(params ...interface{}) *fields.Time {
	field := &fields.Time{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 周组件
func (p *Field) Week(params ...interface{}) *fields.Week {
	field := &fields.Week{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 年组件
func (p *Field) Year(params ...interface{}) *fields.Year {
	field := &fields.Year{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// Select组合组件
func (p *Field) Selects(body interface{}) *fields.Selects {
	field := &fields.Selects{}

	field.Init().SetBody(body)

	return field
}

// 树选择组件
func (p *Field) TreeSelect(params ...interface{}) *fields.TreeSelect {
	field := (&fields.TreeSelect{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}
