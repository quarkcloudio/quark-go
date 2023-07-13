package resource

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/cascader"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/checkbox"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/compact"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/date"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/daterange"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/datetime"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/datetimerange"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/dependency"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/display"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/editor"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/fieldset"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/file"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/geofence"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/group"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/hidden"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/icon"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/id"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/image"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/list"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/mapfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/month"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/number"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/password"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/quarter"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/search"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selects"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/space"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/switchfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/text"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/textarea"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/time"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/timerange"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/tree"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/week"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/year"
)

// 后台字段组件
type Field struct{}

// 解析字段
func fieldParser(field interface{}, params []interface{}, placeholder string) interface{} {
	v := reflect.ValueOf(field).Elem()

	switch len(params) {
	case 1:
		v.FieldByName("Name").Set(reflect.ValueOf(params[0].(string)))
	case 2:
		v.FieldByName("Name").Set(reflect.ValueOf(params[0]))
		v.FieldByName("Label").Set(reflect.ValueOf(params[1]))
	case 3:
		v.FieldByName("Name").Set(reflect.ValueOf(params[0]))
		v.FieldByName("Label").Set(reflect.ValueOf(params[1]))

		// 判断是否为闭包函数
		closure, ok := params[2].(func() interface{})
		if ok {
			v.FieldByName("Callback").Set(reflect.ValueOf(closure))
		}
	}

	if placeholder != "" && len(params) > 1 {
		v.FieldByName("Placeholder").Set(reflect.ValueOf(placeholder + params[1].(string)))
	}

	return field
}

// ID组件
//
// field.ID("id", "ID") 或 field.ID("id", "ID", func() interface{} { return p.Field["username"] })
func (p *Field) ID(params ...interface{}) *id.Component {
	return fieldParser(id.New(), params, "").(*id.Component)
}

// Hidden组件
//
// field.Hidden("id", "ID") 或 field.Hidden("id", "ID", func() interface{} { return p.Field["username"] })
func (p *Field) Hidden(params ...interface{}) *hidden.Component {
	return fieldParser(hidden.New(), params, "").(*hidden.Component)
}

// 输入框组件
//
// field.Text("username", "输入框") 或 field.Text("username", "输入框", func() interface{} { return p.Field["username"] })
func (p *Field) Text(params ...interface{}) *text.Component {
	return fieldParser(text.New(), params, "请输入").(*text.Component)
}

// 文本域组件
//
// field.TextArea("name", "文本域") 或 field.TextArea("name", "文本域", func() interface{} { return p.Field["name"] })
func (p *Field) TextArea(params ...interface{}) *textarea.Component {
	return fieldParser(textarea.New(), params, "请输入").(*textarea.Component)
}

// 密码组件
//
// field.Password("name", "密码框") 或 field.Password("name", "密码", func() interface{} { return p.Field["name"] })
func (p *Field) Password(params ...interface{}) *password.Component {
	return fieldParser(password.New(), params, "请输入").(*password.Component)
}

// 单选组件
//
// field.Radio("name", "单选") 或 field.Radio("name", "单选", func() interface{} { return p.Field["name"] })
func (p *Field) Radio(params ...interface{}) *radio.Component {
	return fieldParser(radio.New(), params, "").(*radio.Component)
}

// 多选组件
//
// field.Checkbox("name", "多选") 或 field.Checkbox("name", "多选", func() interface{} { return p.Field["name"] })
func (p *Field) Checkbox(params ...interface{}) *checkbox.Component {
	return fieldParser(checkbox.New(), params, "").(*checkbox.Component)
}

// 日期组件
//
// field.Date("name", "日期") 或 field.Date("name", "日期", func() interface{} { return p.Field["name"] })
func (p *Field) Date(params ...interface{}) *date.Component {
	return fieldParser(date.New(), params, "").(*date.Component)
}

// 日期范围组件
//
// field.DateRange("name", "日期范围") 或 field.DateRange("name", "日期范围", func() interface{} { return p.Field["name"] })
func (p *Field) DateRange(params ...interface{}) *daterange.Component {
	return fieldParser(daterange.New(), params, "").(*daterange.Component)
}

// 日期时间组件
//
// field.Datetime("name", "日期时间") 或 field.Datetime("name", "日期时间", func() interface{} { return p.Field["name"] })
func (p *Field) Datetime(params ...interface{}) *datetime.Component {
	return fieldParser(datetime.New(), params, "").(*datetime.Component)
}

// 日期时间范围组件
//
// field.DatetimeRange("name", "日期时间范围") 或 field.DatetimeRange("name", "日期时间范围", func() interface{} { return p.Field["name"] })
func (p *Field) DatetimeRange(params ...interface{}) *datetimerange.Component {
	return fieldParser(datetimerange.New(), params, "").(*datetimerange.Component)
}

// 开关组件
//
// field.Switch("name", "开关") 或 field.Switch("name", "开关", func() interface{} { return p.Field["name"] })
func (p *Field) Switch(params ...interface{}) *switchfield.Component {
	return fieldParser(switchfield.New(), params, "").(*switchfield.Component)
}

// 树形组件
//
// field.Tree("name", "树形组件") 或 field.Tree("name", "树形组件", func() interface{} { return p.Field["name"] })
func (p *Field) Tree(params ...interface{}) *tree.Component {
	return fieldParser(tree.New(), params, "请选择").(*tree.Component)
}

// 图标组件
//
// field.Icon("name", "图标") 或 field.Icon("name", "图标", func() interface{} { return p.Field["name"] })
func (p *Field) Icon(params ...interface{}) *icon.Component {
	return fieldParser(icon.New(), params, "请选择").(*icon.Component)
}

// 下拉框组件
//
// field.Select("name", "文本") 或 field.Select("name", "文本", func() interface{} { return p.Field["name"] })
func (p *Field) Select(params ...interface{}) *selectfield.Component {
	return fieldParser(selectfield.New(), params, "请选择").(*selectfield.Component)
}

// 级联菜单组件
//
// field.Cascader("name", "级联菜单") 或 field.Cascader("name", "级联菜单", func() interface{} { return p.Field["name"] })
func (p *Field) Cascader(params ...interface{}) *cascader.Component {
	return fieldParser(cascader.New(), params, "请选择").(*cascader.Component)
}

// 图片组件
//
// field.Image("name", "文本") 或 field.Image("name", "文本", func() interface{} { return p.Field["name"] })
func (p *Field) Image(params ...interface{}) *image.Component {
	return fieldParser(image.New(), params, "").(*image.Component)
}

// 文件组件
//
// field.File("name", "文件") 或 field.File("name", "文件", func() interface{} { return p.Field["name"] })
func (p *Field) File(params ...interface{}) *file.Component {
	return fieldParser(file.New(), params, "").(*file.Component)
}

// 文本展示组件
//
// field.Display("文本")
func (p *Field) Display(label string) *display.Component {
	field := display.
		New().
		SetLabel(label)

	return field
}

// 编辑器组件
//
// field.Editor("name", "编辑器") 或 field.Editor("name", "编辑器", func() interface{} { return p.Field["name"] })
func (p *Field) Editor(params ...interface{}) *editor.Component {
	return fieldParser(editor.New(), params, "").(*editor.Component)
}

// 分组组件
//
// field.Group("分组组件",[]interface{}{field.Text("title", "标题"),field.Number("num","奖品数量")})
func (p *Field) Group(title string, items []interface{}) *group.Component {
	field := group.
		New().
		SetTitle(title).
		SetBody(items)

	return field
}

// List组件
//
// field.List("name", "List") 或 field.List("name", "List", func() interface{} { return p.Field["name"] })
func (p *Field) List(params ...interface{}) *list.Component {
	return fieldParser(list.New(), params, "").(*list.Component)
}

// 地图组件
//
// field.Map("name", "地图") 或 field.Map("name", "地图", func() interface{} { return p.Field["name"] })
func (p *Field) Map(params ...interface{}) *mapfield.Component {
	return fieldParser(mapfield.New(), params, "").(*mapfield.Component)
}

// 地图围栏组件
//
// field.Geofence("name", "地图围栏") 或 field.Geofence("name", "地图围栏", func() interface{} { return p.Field["name"] })
func (p *Field) Geofence(params ...interface{}) *geofence.Component {
	return fieldParser(geofence.New(), params, "").(*geofence.Component)
}

// 日期-月组件
//
// field.Month("name", "月") 或 field.Month("name", "月", func() interface{} { return p.Field["name"] })
func (p *Field) Month(params ...interface{}) *month.Component {
	return fieldParser(month.New(), params, "").(*month.Component)
}

// 数字输入框组件
//
// field.Number("name", "数字") 或 field.Number("name", "数字", func() interface{} { return p.Field["name"] })
func (p *Field) Number(params ...interface{}) *number.Component {
	return fieldParser(number.New(), params, "请输入").(*number.Component)
}

// 日期-季度组件
//
// field.Quarter("name", "季度") 或 field.Quarter("name", "季度", func() interface{} { return p.Field["name"] })
func (p *Field) Quarter(params ...interface{}) *quarter.Component {
	return fieldParser(quarter.New(), params, "").(*quarter.Component)
}

// 搜索组件
//
// field.Search("name", "搜索") 或 field.Search("name", "搜索", func() interface{} { return p.Field["name"] })
func (p *Field) Search(params ...interface{}) *search.Component {
	return fieldParser(search.New(), params, "").(*search.Component)
}

// 时间范围组件
//
// field.TimeRange("name", "时间范围") 或 field.TimeRange("name", "时间范围", func() interface{} { return p.Field["name"] })
func (p *Field) TimeRange(params ...interface{}) *timerange.Component {
	return fieldParser(timerange.New(), params, "").(*timerange.Component)
}

// 时间组件
//
// field.Time("name", "时间") 或 field.Time("name", "时间", func() interface{} { return p.Field["name"] })
func (p *Field) Time(params ...interface{}) *time.Component {
	return fieldParser(time.New(), params, "").(*time.Component)
}

// 周组件
//
// field.Week("name", "周") 或 field.Week("name", "周", func() interface{} { return p.Field["name"] })
func (p *Field) Week(params ...interface{}) *week.Component {
	return fieldParser(week.New(), params, "").(*week.Component)
}

// 年组件
//
// field.Year("name", "年") 或 field.Year("name", "年", func() interface{} { return p.Field["name"] })
func (p *Field) Year(params ...interface{}) *year.Component {
	return fieldParser(year.New(), params, "").(*year.Component)
}

// Select联动组件
//
//	field.Selects([]interface{}{
//		field.Select("province", "省").
//			SetOptions([]*selectfield.Option{
//				{Value: 1, Label: "北京"},
//				{Value: 2, Label: "天津"},
//				{Value: 3, Label: "河北省"},
//			}).SetLoad("city", "/api/admin/test/cities"),
//		field.Select("city", "市"),
//	}),
func (p *Field) Selects(body interface{}) *selects.Component {
	field := selects.
		New().
		SetBody(body)

	return field
}

// 树选择组件
//
// field.TreeSelect("name", "树选择") 或 field.TreeSelect("name", "树选择", func() interface{} { return p.Field["name"] })
func (p *Field) TreeSelect(params ...interface{}) *treeselect.Component {
	return fieldParser(treeselect.New(), params, "").(*treeselect.Component)
}

// 间距布局组件
//
// field.Space("间距组件",[]interface{}{field.Text("title", "标题"),field.Number("num","奖品数量")})
func (p *Field) Space(label string, items []interface{}) *space.Component {
	field := space.
		New().
		SetLabel(label).
		SetBody(items)

	return field
}

// 紧凑布局组件
//
// field.Compact("紧凑布局",[]interface{}{field.Text("title", "标题"),field.Number("num","奖品数量")})
func (p *Field) Compact(label string, items []interface{}) *compact.Component {
	field := compact.
		New().
		SetLabel(label).
		SetBody(items)

	return field
}

// 录入结构化的一维数组数据
//
// field.FieldSet("name", "组件列表") 或 field.FieldSet("name", "组件列表", func() interface{} { return p.Field["name"] })
func (p *Field) FieldSet(params ...interface{}) *fieldset.Component {
	return fieldParser(fieldset.New(), params, "").(*fieldset.Component)
}

// 数据联动组件
//
// field.Dependency()
func (p *Field) Dependency() *dependency.Component {
	field := dependency.New()

	return field
}
