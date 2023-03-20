package number

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Number struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`        // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`        // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`  // 配合 valiTextStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
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
	UpTextRules    []*rule.Rule  `json:"-"`             // 编辑页校验规则
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

	AddonAfter       interface{}            `json:"addonAfter,omitempty"`       // 带标签的 input，设置后置标签
	AddonBefore      interface{}            `json:"addonBefore,omitempty"`      // 带标签的 input，设置前置标签
	AllowClear       bool                   `json:"allowClear,omitempty"`       // 可以点击清除图标删除内容
	Bordered         bool                   `json:"bordered,omitempty"`         // 是否有边框，默认true
	Controls         bool                   `json:"controls,omitempty"`         // 是否显示增减按钮，也可设置自定义箭头图标
	DecimalSeparator string                 `json:"decimalSeparator,omitempty"` // 小数点
	DefaultValue     interface{}            `json:"defaultValue,omitempty"`     // 默认的选中项
	Disabled         interface{}            `json:"disabled,omitempty"`         // 禁用
	Keyboard         bool                   `json:"keyboard,omitempty"`         // 是否启用键盘快捷行为
	Max              int                    `json:"max,omitempty"`              // 最大值
	Min              int                    `json:"min,omitempty"`              // 最小值
	Precision        int                    `json:"precision,omitempty"`        // 数值精度，配置 formatter 时会以 formatter 为准
	ReadOnly         bool                   `json:"readOnly,omitempty"`         // 只读
	Status           string                 `json:"status,omitempty"`           // 设置校验状态,'error' | 'warning'
	Prefix           interface{}            `json:"prefix,omitempty"`           // 带有前缀图标的 input
	Size             string                 `json:"size,omitempty"`             // 控件大小。注：标准表单内的输入框大小限制为 middle，large | middle | small
	Step             interface{}            `json:"step,omitempty"`             // 每次改变步数，可以为小数
	StringMode       bool                   `json:"stringMode,omitempty"`       // 字符值模式，开启后支持高精度小数。同时 onChange 将返回 string 类型
	Value            interface{}            `json:"value,omitempty"`            // 指定选中项,string[] | number[]
	Placeholder      string                 `json:"placeholder,omitempty"`      // 占位符
	Style            map[string]interface{} `json:"style,omitempty"`            // 自定义样式
}

// 初始化组件
func New() *Number {
	return (&Number{}).Init()
}

// 初始化
func (p *Number) Init() *Number {
	p.Component = "numberField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()
	p.Placeholder = "请输入"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)

	return p
}

// 设置Key
func (p *Number) SetKey(key string, crypt bool) *Number {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *Number) SetStyle(style map[string]interface{}) *Number {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Number) SetTooltip(tooltip string) *Number {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Number) SetWidth(width interface{}) *Number {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Number) SetColon(colon bool) *Number {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Number) SetExtra(extra string) *Number {
	p.Extra = extra
	return p
}

// 配合 valiTextStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Number) SetHasFeedback(hasFeedback bool) *Number {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Number) SetHelp(help string) *Number {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Number) SetNoStyle() *Number {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Number) SetLabel(label string) *Number {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Number) SetLabelAlign(align string) *Number {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Number) SetLabelCol(col interface{}) *Number {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Number) SetName(name string) *Number {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Number) SetRequired() *Number {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Number) GetFrontendRules(path string) *Number {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		upTextRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "upText")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpTextRules) > 0 {
		upTextRules = rule.ConvertToFrontendRules(p.UpTextRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(upTextRules) > 0 {
		frontendRules = append(frontendRules, upTextRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *Number) SetRules(rules []*rule.Rule) *Number {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Number) SetCreationRules(rules []*rule.Rule) *Number {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Number) SetUpTextRules(rules []*rule.Rule) *Number {
	p.UpTextRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Number) SetValuePropName(valuePropName string) *Number {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Number) SetWrapperCol(col interface{}) *Number {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Number) SetValue(value interface{}) *Number {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Number) SetDefault(value interface{}) *Number {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Number) SetDisabled(disabled bool) *Number {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Number) SetIgnore(ignore bool) *Number {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Number) SetWhen(value ...any) *Number {
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
func (p *Number) HideFromIndex(callback bool) *Number {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Number) HideFromDetail(callback bool) *Number {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Number) HideWhenCreating(callback bool) *Number {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Number) HideWhenUpdating(callback bool) *Number {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Number) HideWhenExporting(callback bool) *Number {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Number) HideWhenImporting(callback bool) *Number {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Number) OnIndexShowing(callback bool) *Number {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Number) OnDetailShowing(callback bool) *Number {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Number) ShowOnCreating(callback bool) *Number {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Number) ShowOnUpdating(callback bool) *Number {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Number) ShowOnExporting(callback bool) *Number {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Number) ShowOnImporting(callback bool) *Number {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Number) OnlyOnIndex() *Number {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Number) OnlyOnDetail() *Number {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Number) OnlyOnForms() *Number {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Number) OnlyOnExport() *Number {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Number) OnlyOnImport() *Number {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Number) ExceptOnForms() *Number {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Number) IsShownOnUpText() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Number) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Number) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Number) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Number) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Number) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Number) SetEditable(editable bool) *Number {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Number) SetColumn(f func(column *table.Column) *table.Column) *Number {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Number) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Number) SetCallback(closure func() interface{}) *Number {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Number) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Number) SetApi(api string) *Number {
	p.Api = api
	return p
}

// 带标签的 input，设置后置标签
func (p *Number) SetAddonAfter(addonAfter interface{}) *Number {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签
func (p *Number) SetAddonBefore(addonBefore interface{}) *Number {
	p.AddonBefore = addonBefore

	return p
}

// 可以点击清除图标删除内容
func (p *Number) SetAllowClear(allowClear bool) *Number {
	p.AllowClear = allowClear

	return p
}

// 是否有边框，默认true
func (p *Number) SetBordered(bordered bool) *Number {
	p.Bordered = bordered

	return p
}

// 是否显示增减按钮，也可设置自定义箭头图标
func (p *Number) SetControls(controls bool) *Number {
	p.Controls = controls

	return p
}

// 小数点
func (p *Number) SetDecimalSeparator(decimalSeparator string) *Number {
	p.DecimalSeparator = decimalSeparator

	return p
}

// 输入框默认内容
func (p *Number) SetDefaultValue(defaultValue interface{}) *Number {
	p.DefaultValue = defaultValue

	return p
}

// 是否启用键盘快捷行为
func (p *Number) SetKeyboard(keyboard bool) *Number {
	p.Keyboard = keyboard

	return p
}

// 最大值
func (p *Number) SetMax(max int) *Number {
	p.Max = max

	return p
}

// 最小值
func (p *Number) SetMin(min int) *Number {
	p.Min = min

	return p
}

// 数值精度，配置 formatter 时会以 formatter 为准
func (p *Number) SetPrecision(precision int) *Number {
	p.Precision = precision

	return p
}

// 只读
func (p *Number) SetReadOnly(readOnly bool) *Number {
	p.ReadOnly = readOnly

	return p
}

// 设置校验状态，'error' | 'warning'
func (p *Number) SetStatus(status string) *Number {
	p.Status = status

	return p
}

// 输入框占位文本
func (p *Number) SetPlaceholder(placeholder string) *Number {
	p.Placeholder = placeholder

	return p
}

// 带有前缀图标的 input
func (p *Number) SetPrefix(prefix interface{}) *Number {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Number) SetSize(size string) *Number {
	p.Size = size

	return p
}

// 每次改变步数，可以为小数
func (p *Number) SetStep(step interface{}) *Number {
	p.Step = step

	return p
}

// 字符值模式，开启后支持高精度小数。同时 onChange 将返回 string 类型
func (p *Number) SetStringMode(stringMode bool) *Number {
	p.StringMode = stringMode

	return p
}
