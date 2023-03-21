package TextArea

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type TextArea struct {
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

	AddonAfter   interface{}            `json:"addonAfter,omitempty"`   // 带标签的 input，设置后置标签
	AddonBefore  interface{}            `json:"addonBefore,omitempty"`  // 带标签的 input，设置前置标签
	AllowClear   bool                   `json:"allowClear,omitempty"`   // 可以点击清除图标删除内容
	Bordered     bool                   `json:"bordered,omitempty"`     // 是否有边框，默认true
	DefaultValue interface{}            `json:"defaultValue,omitempty"` // 默认的选中项
	Disabled     interface{}            `json:"disabled,omitempty"`     // 禁用
	Id           string                 `json:"id,omitempty"`           // 输入框的 id
	MaxLength    int                    `json:"maxLength,omitempty"`    //最大长度
	ShowCount    bool                   `json:"showCount,omitempty"`    // 是否展示字数
	Status       string                 `json:"status,omitempty"`       // 设置校验状态,'error' | 'warning'
	Prefix       interface{}            `json:"prefix,omitempty"`       // 带有前缀图标的 input
	Size         string                 `json:"size,omitempty"`         // 控件大小。注：标准表单内的输入框大小限制为 middle，large | middle | small
	Suffix       interface{}            `json:"suffix,omitempty"`       // 带有后缀图标的 input
	Type         string                 `json:"type,omitempty"`         // 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="TextArea")
	Value        interface{}            `json:"value,omitempty"`        // 指定选中项,string[] | number[]
	Placeholder  string                 `json:"placeholder,omitempty"`  // 占位符
	Style        map[string]interface{} `json:"style,omitempty"`        // 自定义样式

}

// 初始化组件
func New() *TextArea {
	return (&TextArea{}).Init()
}

// 初始化
func (p *TextArea) Init() *TextArea {
	p.Component = "textAreaField"
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
	p.MaxLength = 200

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)

	return p
}

// 设置Key
func (p *TextArea) SetKey(key string, crypt bool) *TextArea {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *TextArea) SetStyle(style map[string]interface{}) *TextArea {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *TextArea) SetTooltip(tooltip string) *TextArea {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *TextArea) SetWidth(width interface{}) *TextArea {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *TextArea) SetColon(colon bool) *TextArea {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *TextArea) SetExtra(extra string) *TextArea {
	p.Extra = extra
	return p
}

// 配合 valiTextStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *TextArea) SetHasFeedback(hasFeedback bool) *TextArea {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *TextArea) SetHelp(help string) *TextArea {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *TextArea) SetNoStyle() *TextArea {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *TextArea) SetLabel(label string) *TextArea {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *TextArea) SetLabelAlign(align string) *TextArea {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *TextArea) SetLabelCol(col interface{}) *TextArea {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *TextArea) SetName(name string) *TextArea {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *TextArea) SetRequired() *TextArea {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *TextArea) GetFrontendRules(path string) *TextArea {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		UpdateRules   []*rule.Rule
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
	if isEditing && len(p.UpdateRules) > 0 {
		UpdateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(UpdateRules) > 0 {
		frontendRules = append(frontendRules, UpdateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *TextArea) SetRules(rules []*rule.Rule) *TextArea {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *TextArea) SetCreationRules(rules []*rule.Rule) *TextArea {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *TextArea) SetUpdateRules(rules []*rule.Rule) *TextArea {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *TextArea) SetValuePropName(valuePropName string) *TextArea {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *TextArea) SetWrapperCol(col interface{}) *TextArea {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *TextArea) SetValue(value interface{}) *TextArea {
	p.Value = value
	return p
}

// 设置默认值。
func (p *TextArea) SetDefault(value interface{}) *TextArea {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *TextArea) SetDisabled(disabled bool) *TextArea {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *TextArea) SetIgnore(ignore bool) *TextArea {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *TextArea) SetWhen(value ...any) *TextArea {
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
func (p *TextArea) HideFromIndex(callback bool) *TextArea {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *TextArea) HideFromDetail(callback bool) *TextArea {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *TextArea) HideWhenCreating(callback bool) *TextArea {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *TextArea) HideWhenUpdating(callback bool) *TextArea {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *TextArea) HideWhenExporting(callback bool) *TextArea {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *TextArea) HideWhenImporting(callback bool) *TextArea {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *TextArea) OnIndexShowing(callback bool) *TextArea {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *TextArea) OnDetailShowing(callback bool) *TextArea {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *TextArea) ShowOnCreating(callback bool) *TextArea {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *TextArea) ShowOnUpdating(callback bool) *TextArea {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *TextArea) ShowOnExporting(callback bool) *TextArea {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *TextArea) ShowOnImporting(callback bool) *TextArea {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *TextArea) OnlyOnIndex() *TextArea {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *TextArea) OnlyOnDetail() *TextArea {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *TextArea) OnlyOnForms() *TextArea {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *TextArea) OnlyOnExport() *TextArea {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *TextArea) OnlyOnImport() *TextArea {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *TextArea) ExceptOnForms() *TextArea {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *TextArea) IsShownOnUpText() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *TextArea) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *TextArea) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *TextArea) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *TextArea) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *TextArea) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *TextArea) SetEditable(editable bool) *TextArea {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *TextArea) SetColumn(f func(column *table.Column) *table.Column) *TextArea {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *TextArea) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *TextArea) SetCallback(closure func() interface{}) *TextArea {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *TextArea) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *TextArea) SetApi(api string) *TextArea {
	p.Api = api
	return p
}

// 带标签的 input，设置后置标签
func (p *TextArea) SetAddonAfter(addonAfter interface{}) *TextArea {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签
func (p *TextArea) SetAddonBefore(addonBefore interface{}) *TextArea {
	p.AddonBefore = addonBefore

	return p
}

// 可以点击清除图标删除内容
func (p *TextArea) SetAllowClear(allowClear bool) *TextArea {
	p.AllowClear = allowClear

	return p
}

// 是否有边框，默认true
func (p *TextArea) SetBordered(bordered bool) *TextArea {
	p.Bordered = bordered

	return p
}

// 输入框默认内容
func (p *TextArea) SetDefaultValue(defaultValue interface{}) *TextArea {
	p.DefaultValue = defaultValue

	return p
}

// 输入框的 id
func (p *TextArea) SetId(id string) *TextArea {
	p.Id = id

	return p
}

// 最大长度
func (p *TextArea) SetMaxLength(maxLength int) *TextArea {
	p.MaxLength = maxLength

	return p
}

// 是否展示字数
func (p *TextArea) SetShowCount(showCount bool) *TextArea {
	p.ShowCount = showCount

	return p
}

// 设置校验状态，'error' | 'warning'
func (p *TextArea) SetStatus(status string) *TextArea {
	p.Status = status

	return p
}

// 输入框占位文本
func (p *TextArea) SetPlaceholder(placeholder string) *TextArea {
	p.Placeholder = placeholder

	return p
}

// 带有前缀图标的 input
func (p *TextArea) SetPrefix(prefix interface{}) *TextArea {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *TextArea) SetSize(size string) *TextArea {
	p.Size = size

	return p
}

// 带有后缀图标的 input
func (p *TextArea) SetSuffix(suffix interface{}) *TextArea {
	p.Suffix = suffix

	return p
}

// 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="TextArea")
func (p *TextArea) SetType(Type string) *TextArea {
	p.Type = Type

	return p
}
