package year

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Year struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`        // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`        // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`  // 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
	Help          string      `json:"help,omitempty"`         // 提示信息，如不设置，则会根据校验规则自动生成
	Hidden        bool        `json:"hidden,omitempty"`       // 是否隐藏字段（依然会收集和校验字段）
	InitialValue  interface{} `json:"initialValue,omitempty"` // 设置子元素默认值，如果与 Form 的 initialValues 冲突则以 Form 为准
	Label         string      `json:"label,omitempty"`        // label 标签的文本
	LabelAlign    string      `json:"labelAlign,omitempty"`   // 标签文本对齐方式
	LabelCol      interface{} `json:"labelCol,omitempty"`     // label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。你可以通过 Form 的 labelCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准
	Name          string      `json:"name,omitempty"`         // 字段名，支持数组
	NoStyle       bool        `json:"noStyle,omitempty"`      // 为 true 时不带样式，作为纯字段控件使用
	Required      bool        `json:"required,omitempty"`     // 必填样式设置。如不设置，则会根据校验规则自动生成
	Tooltip       string      `json:"tooltip,omitempty"`      // 会在 label 旁增加一个 icon，悬浮后展示配置的信息
	ValuePropName string      `json:"valuePropName"`          // 子节点的值的属性，如 Switch 的是 'checked'。该属性为 getValueProps 的封装，自定义 getValueProps 后会失效
	WrapperCol    interface{} `json:"wrapperCol"`             // 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。你可以通过 Form 的 wrapperCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准

	Api            string        `json:"api,omitempty"` // 获取数据接口
	Ignore         bool          `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules          []*rule.Rule  `json:"-"`             // 全局校验规则
	CreationRules  []*rule.Rule  `json:"-"`             // 创建页校验规则
	UpdateRules    []*rule.Rule  `json:"-"`             // 编辑页校验规则
	FrontendRules  []*rule.Rule  `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When           *when.When    `json:"when"`          //
	WhenItem       []*when.Item  `json:"-"`             //
	ShowOnIndex    bool          `json:"-"`             // 在列表页展示
	ShowOnDetail   bool          `json:"-"`             // 在详情页展示
	ShowOnCreation bool          `json:"-"`             // 在创建页面展示
	ShowOnUpdate   bool          `json:"-"`             // 在编辑页面展示
	ShowOnExport   bool          `json:"-"`             // 在导出的Excel上展示
	ShowOnImport   bool          `json:"-"`             // 在导入Excel上展示
	Editable       bool          `json:"-"`             // 表格上是否可编辑
	Column         *table.Column `json:"-"`             // 表格列
	Callback       interface{}   `json:"-"`             // 回调函数

	AllowClear     bool                   `json:"allowClear,omitempty"`     // 是否支持清除，默认true
	AutoFocus      bool                   `json:"autoFocus,omitempty"`      // 自动获取焦点，默认false
	Bordered       bool                   `json:"bordered,omitempty"`       // 是否有边框，默认true
	ClassName      string                 `json:"className,omitempty"`      // 自定义类名
	DefaultValue   interface{}            `json:"defaultValue,omitempty"`   // 默认的选中项
	Disabled       interface{}            `json:"disabled,omitempty"`       // 禁用
	Format         string                 `json:"format,omitempty"`         // 设置日期格式，为数组时支持多格式匹配，展示以第一个为准。
	PopupClassName string                 `json:"popupClassName,omitempty"` // 额外的弹出日历 className
	InputReadOnly  bool                   `json:"inputReadOnly,omitempty"`  // 设置输入框为只读（避免在移动设备上打开虚拟键盘）
	Locale         interface{}            `json:"locale,omitempty"`         // 国际化配置
	Mode           string                 `json:"mode,omitempty"`           // 日期面板的状态 time | date | month | year | decade
	NextIcon       interface{}            `json:"nextIcon,omitempty"`       // 自定义下一个图标
	Open           bool                   `json:"open,omitempty"`           // 控制浮层显隐
	Picker         string                 `json:"picker,omitempty"`         // 设置选择器类型 date | week | month | quarter | year
	Placeholder    string                 `json:"placeholder,omitempty"`    // 输入框占位文本
	Placement      string                 `json:"placement,omitempty"`      // 浮层预设位置，bottomLeft bottomRight topLeft topRight
	PopupStyle     interface{}            `json:"popupStyle,omitempty"`     // 额外的弹出日历样式
	PrevIcon       interface{}            `json:"prevIcon,omitempty"`       // 自定义上一个图标
	Size           string                 `json:"size,omitempty"`           // 输入框大小，large | middle | small
	Status         string                 `json:"status,omitempty"`         // 设置校验状态，'error' | 'warning'
	Style          map[string]interface{} `json:"style,omitempty"`          // 自定义样式
	SuffixIcon     interface{}            `json:"suffixIcon,omitempty"`     // 自定义的选择框后缀图标
	SuperNextIcon  interface{}            `json:"superNextIcon,omitempty"`  // 自定义 << 切换图标
	SuperPrevIcon  interface{}            `json:"superPrevIcon,omitempty"`  // 自定义 >> 切换图标
	Value          interface{}            `json:"value,omitempty"`          // 指定选中项,string[] | number[]

	DefaultPickerValue string      `json:"defaultPickerValue,omitempty"` // 默认面板日期
	ShowNow            bool        `json:"showNow,omitempty"`            // 当设定了 showTime 的时候，面板是否显示“此刻”按钮
	ShowTime           interface{} `json:"showTime,omitempty"`           // 增加时间选择功能
	ShowToday          bool        `json:"showToday,omitempty"`          // 是否展示“今天”按钮
}

// 初始化组件
func New() *Year {
	return (&Year{}).Init()
}

// 初始化
func (p *Year) Init() *Year {
	p.Component = "yearField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()
	p.Placeholder = "请选择"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Year) SetKey(key string, crypt bool) *Year {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *Year) SetStyle(style map[string]interface{}) *Year {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Year) SetTooltip(tooltip string) *Year {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Year) SetWidth(width interface{}) *Year {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Year) SetColon(colon bool) *Year {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Year) SetExtra(extra string) *Year {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Year) SetHasFeedback(hasFeedback bool) *Year {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Year) SetHelp(help string) *Year {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Year) SetNoStyle() *Year {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Year) SetLabel(label string) *Year {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Year) SetLabelAlign(align string) *Year {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Year) SetLabelCol(col interface{}) *Year {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Year) SetName(name string) *Year {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Year) SetRequired() *Year {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Year) GetFrontendRules(path string) *Year {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		updateRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(updateRules) > 0 {
		frontendRules = append(frontendRules, updateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *Year) SetRules(rules []*rule.Rule) *Year {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Year) SetCreationRules(rules []*rule.Rule) *Year {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Year) SetUpdateRules(rules []*rule.Rule) *Year {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Year) SetValuePropName(valuePropName string) *Year {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Year) SetWrapperCol(col interface{}) *Year {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Year) SetValue(value interface{}) *Year {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Year) SetDefault(value interface{}) *Year {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Year) SetDisabled(disabled bool) *Year {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Year) SetIgnore(ignore bool) *Year {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Year) SetWhen(value ...any) *Year {
	w := when.New()
	i := when.NewItem()
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		i.Body = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		i.Body = callback()
	}

	getOption := untils.InterfaceToString(option)

	switch operator {
	case "=":
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	case ">":
		i.Condition = "<%=String(" + p.Name + ") > '" + getOption + "' %>"
		break
	case "<":
		i.Condition = "<%=String(" + p.Name + ") < '" + getOption + "' %>"
		break
	case "<=":
		i.Condition = "<%=String(" + p.Name + ") <= '" + getOption + "' %>"
		break
	case ">=":
		i.Condition = "<%=String(" + p.Name + ") => '" + getOption + "' %>"
		break
	case "has":
		i.Condition = "<%=(String(" + p.Name + ").indexOf('" + getOption + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		i.Condition = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	}

	i.ConditionName = p.Name
	i.ConditionOperator = operator
	i.Option = option
	p.WhenItem = append(p.WhenItem, i)
	p.When = w.SetItems(p.WhenItem)

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Year) HideFromIndex(callback bool) *Year {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Year) HideFromDetail(callback bool) *Year {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Year) HideWhenCreating(callback bool) *Year {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Year) HideWhenUpdating(callback bool) *Year {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Year) HideWhenExporting(callback bool) *Year {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Year) HideWhenImporting(callback bool) *Year {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Year) OnIndexShowing(callback bool) *Year {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Year) OnDetailShowing(callback bool) *Year {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Year) ShowOnCreating(callback bool) *Year {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Year) ShowOnUpdating(callback bool) *Year {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Year) ShowOnExporting(callback bool) *Year {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Year) ShowOnImporting(callback bool) *Year {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Year) OnlyOnIndex() *Year {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Year) OnlyOnDetail() *Year {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Year) OnlyOnForms() *Year {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Year) OnlyOnExport() *Year {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Year) OnlyOnImport() *Year {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Year) ExceptOnForms() *Year {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Year) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Year) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Year) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Year) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Year) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Year) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Year) SetEditable(editable bool) *Year {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Year) SetColumn(f func(column *table.Column) *table.Column) *Year {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Year) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Year) SetCallback(closure func() interface{}) *Year {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Year) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Year) SetApi(api string) *Year {
	p.Api = api
	return p
}

// 可以点击清除图标删除内容
func (p *Year) SetAllowClear(allowClear bool) *Year {
	p.AllowClear = allowClear

	return p
}

// 自动获取焦点，默认false
func (p *Year) SetAutoFocus(autoFocus bool) *Year {
	p.AutoFocus = autoFocus

	return p
}

// 是否有边框，默认true
func (p *Year) SetBordered(bordered bool) *Year {
	p.Bordered = bordered

	return p
}

// 自定义类名
func (p *Year) SetClassName(className string) *Year {
	p.ClassName = className

	return p
}

// 默认的选中项
func (p *Year) SetDefaultValue(defaultValue interface{}) *Year {
	p.DefaultValue = defaultValue

	return p
}

// 设置日期格式，为数组时支持多格式匹配，展示以第一个为准。
func (p *Year) SetFormat(format string) *Year {
	p.Format = format

	return p
}

// 自定义类名
func (p *Year) SetPopupClassName(popupClassName string) *Year {
	p.PopupClassName = popupClassName

	return p
}

// 设置输入框为只读（避免在移动设备上打开虚拟键盘）
func (p *Year) SetInputReadOnly(inputReadOnly bool) *Year {
	p.InputReadOnly = inputReadOnly

	return p
}

// 国际化配置
func (p *Year) SetLocale(locale interface{}) *Year {
	p.Locale = locale

	return p
}

// 日期面板的状态 time | date | month | year | decade
func (p *Year) SetMode(mode string) *Year {
	p.Mode = mode

	return p
}

// 自定义下一个图标
func (p *Year) SetNextIcon(nextIcon interface{}) *Year {
	p.NextIcon = nextIcon

	return p
}

// 控制浮层显隐
func (p *Year) SetOpen(open bool) *Year {
	p.Open = open

	return p
}

// 设置选择器类型 date | week | month | quarter | year
func (p *Year) SetPicker(picker string) *Year {
	p.Picker = picker

	return p
}

// 输入框占位文本
func (p *Year) SetPlaceholder(placeholder string) *Year {
	p.Placeholder = placeholder

	return p
}

// 浮层预设位置，bottomLeft bottomRight topLeft topRight
func (p *Year) SetPlacement(placement string) *Year {
	p.Placement = placement

	return p
}

// 额外的弹出日历样式
func (p *Year) SetPopupStyle(popupStyle interface{}) *Year {
	p.PopupStyle = popupStyle

	return p
}

// 自定义上一个图标
func (p *Year) SetPrevIcon(prevIcon interface{}) *Year {
	p.PrevIcon = prevIcon

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Year) SetSize(size string) *Year {
	p.Size = size

	return p
}

// 设置校验状态，'error' | 'warning'
func (p *Year) SetStatus(status string) *Year {
	p.Status = status

	return p
}

// 自定义的选择框后缀图标
func (p *Year) SetSuffixIcon(suffixIcon interface{}) *Year {
	p.SuffixIcon = suffixIcon

	return p
}

// 自定义 << 切换图标
func (p *Year) SetSuperNextIcon(superNextIcon interface{}) *Year {
	p.SuperNextIcon = superNextIcon

	return p
}

// 自定义 >> 切换图标
func (p *Year) SetSuperPrevIcon(superPrevIcon interface{}) *Year {
	p.SuperPrevIcon = superPrevIcon

	return p
}

// 默认面板日期
func (p *Year) SetDefaultPickerValue(defaultPickerValue string) *Year {
	p.DefaultPickerValue = defaultPickerValue

	return p
}

// 当设定了 showTime 的时候，面板是否显示“此刻”按钮
func (p *Year) SetShowNow(showNow bool) *Year {
	p.ShowNow = showNow

	return p
}

// 增加时间选择功能
func (p *Year) SetShowTime(showTime interface{}) *Year {
	p.ShowTime = showTime

	return p
}

// 是否展示“今天”按钮
func (p *Year) SetShowToday(showToday bool) *Year {
	p.ShowToday = showToday

	return p
}
