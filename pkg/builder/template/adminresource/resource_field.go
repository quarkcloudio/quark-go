package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/cascader"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/checkbox"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/date"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/daterange"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/datetime"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/datetimerange"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/display"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/editor"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/file"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/geofence"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/group"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/hidden"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/icon"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/id"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/image"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/list"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/mapfield"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/month"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/number"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/password"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/quarter"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/radio"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/search"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selectfield"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selects"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/switchfield"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/text"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/textarea"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/time"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/timerange"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/tree"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/treeselect"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/week"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/year"
)

// 后台字段组件
type Field struct{}

// ID组件
func (p *Field) ID(params ...interface{}) *id.Component {
	field := id.New()

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
func (p *Field) Hidden(params ...interface{}) *hidden.Component {
	field := hidden.New()

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
func (p *Field) Text(params ...interface{}) *text.Component {
	field := text.New()

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
func (p *Field) TextArea(params ...interface{}) *textarea.Component {
	field := textarea.New()

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
func (p *Field) Password(params ...interface{}) *password.Component {
	field := password.New()

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
func (p *Field) Radio(params ...string) *radio.Component {
	field := radio.New()

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 多选组件
func (p *Field) Checkbox(params ...string) *checkbox.Component {
	field := checkbox.New()

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 日期组件
func (p *Field) Date(params ...interface{}) *date.Component {
	field := date.New()

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
func (p *Field) DateRange(params ...interface{}) *daterange.Component {
	field := daterange.New()

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
func (p *Field) Datetime(params ...interface{}) *datetime.Component {
	field := datetime.New()

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
func (p *Field) DatetimeRange(params ...interface{}) *datetimerange.Component {
	field := datetimerange.New()

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
func (p *Field) Switch(params ...string) *switchfield.Component {
	field := switchfield.New()

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 树形组件
func (p *Field) Tree(params ...interface{}) *tree.Component {
	field := tree.New()

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
func (p *Field) Icon(params ...interface{}) *icon.Component {
	field := icon.New()

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
func (p *Field) Select(params ...interface{}) *selectfield.Component {
	field := selectfield.New()

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
func (p *Field) Cascader(params ...interface{}) *cascader.Component {
	field := cascader.New()

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
func (p *Field) Image(params ...interface{}) *image.Component {
	field := image.New()

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
func (p *Field) File(params ...interface{}) *file.Component {
	field := file.New()

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
func (p *Field) Display(label string) *display.Component {
	field := display.
		New().
		SetLabel(label)

	return field
}

// 编辑器组件
func (p *Field) Editor(params ...interface{}) *editor.Component {
	field := editor.New()

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
func (p *Field) Group(label string, items []interface{}) *group.Component {
	field := group.
		New().
		SetBody(items).
		SetLabel(label)

	return field
}

// List组件
func (p *Field) List(params ...interface{}) *list.Component {
	field := list.New()

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
func (p *Field) Map(params ...interface{}) *mapfield.Component {
	field := mapfield.New()

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
func (p *Field) Geofence(params ...interface{}) *geofence.Component {
	field := geofence.New()

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
func (p *Field) Month(params ...interface{}) *month.Component {
	field := month.New()

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
func (p *Field) Number(params ...interface{}) *number.Component {
	field := number.New()

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
func (p *Field) Quarter(params ...interface{}) *quarter.Component {
	field := quarter.New()

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
func (p *Field) Search(params ...interface{}) *search.Component {
	field := search.New()

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
func (p *Field) TimeRange(params ...interface{}) *timerange.Component {
	field := timerange.New()

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
func (p *Field) Time(params ...interface{}) *time.Component {
	field := time.New()

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
func (p *Field) Week(params ...interface{}) *week.Component {
	field := week.New()

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
func (p *Field) Year(params ...interface{}) *year.Component {
	field := year.New()

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
func (p *Field) Selects(body interface{}) *selects.Component {
	field := selects.
		New().
		SetBody(body)

	return field
}

// 树选择组件
func (p *Field) TreeSelect(params ...interface{}) *treeselect.Component {
	field := treeselect.New()

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
