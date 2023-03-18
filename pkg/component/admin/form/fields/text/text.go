package text

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Text struct {
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
	ShowOnUpText   bool          `json:"-"`             // 在编辑页面展示
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
	Type         string                 `json:"type,omitempty"`         // 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="textarea")
	Value        interface{}            `json:"value,omitempty"`        // 指定选中项,string[] | number[]
	Placeholder  string                 `json:"placeholder,omitempty"`  // 占位符
	Style        map[string]interface{} `json:"style,omitempty"`        // 自定义样式
}

// 初始化组件
func New() *Text {
	return (&Text{}).Init()
}

// 初始化
func (p *Text) Init() *Text {
	p.Component = "textField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpText = true
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
func (p *Text) SetKey(key string, crypt bool) *Text {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *Text) SetStyle(style map[string]interface{}) *Text {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Text) SetTooltip(tooltip string) *Text {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Text) SetWidth(width interface{}) *Text {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Text) SetColon(colon bool) *Text {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Text) SetExtra(extra string) *Text {
	p.Extra = extra
	return p
}

// 配合 valiTextStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Text) SetHasFeedback(hasFeedback bool) *Text {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Text) SetHelp(help string) *Text {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Text) SetNoStyle() *Text {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Text) SetLabel(label string) *Text {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Text) SetLabelAlign(align string) *Text {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Text) SetLabelCol(col interface{}) *Text {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Text) SetName(name string) *Text {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Text) SetRequired() *Text {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Text) GetFrontendRules(path string) *Text {
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
func (p *Text) SetRules(rules []*rule.Rule) *Text {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Text) SetCreationRules(rules []*rule.Rule) *Text {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Text) SetUpTextRules(rules []*rule.Rule) *Text {
	p.UpTextRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Text) SetValuePropName(valuePropName string) *Text {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Text) SetWrapperCol(col interface{}) *Text {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Text) SetValue(value interface{}) *Text {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Text) SetDefault(value interface{}) *Text {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Text) SetDisabled(disabled bool) *Text {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Text) SetIgnore(ignore bool) *Text {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Text) SetWhen(value ...any) *Text {
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
func (p *Text) HideFromIndex(callback bool) *Text {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Text) HideFromDetail(callback bool) *Text {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Text) HideWhenCreating(callback bool) *Text {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Text) HideWhenUpdating(callback bool) *Text {
	p.ShowOnUpText = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Text) HideWhenExporting(callback bool) *Text {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Text) HideWhenImporting(callback bool) *Text {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Text) OnIndexShowing(callback bool) *Text {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Text) OnDetailShowing(callback bool) *Text {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Text) ShowOnCreating(callback bool) *Text {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Text) ShowOnUpdating(callback bool) *Text {
	p.ShowOnUpText = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Text) ShowOnExporting(callback bool) *Text {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Text) ShowOnImporting(callback bool) *Text {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Text) OnlyOnIndex() *Text {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpText = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Text) OnlyOnDetail() *Text {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpText = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Text) OnlyOnForms() *Text {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpText = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Text) OnlyOnExport() *Text {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpText = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Text) OnlyOnImport() *Text {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpText = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Text) ExceptOnForms() *Text {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpText = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Text) IsShownOnUpText() bool {
	return p.ShowOnUpText
}

// Check showing on index.
func (p *Text) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Text) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Text) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Text) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Text) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Text) SetEditable(editable bool) *Text {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Text) SetColumn(f func(column *table.Column) *table.Column) *Text {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Text) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Text) SetCallback(closure func() interface{}) *Text {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Text) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Text) SetApi(api string) *Text {
	p.Api = api
	return p
}

// 带标签的 input，设置后置标签
func (p *Text) SetAddonAfter(addonAfter interface{}) *Text {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签
func (p *Text) SetAddonBefore(addonBefore interface{}) *Text {
	p.AddonBefore = addonBefore

	return p
}

// 可以点击清除图标删除内容
func (p *Text) SetAllowClear(allowClear bool) *Text {
	p.AllowClear = allowClear

	return p
}

// 是否有边框，默认true
func (p *Text) SetBordered(bordered bool) *Text {
	p.Bordered = bordered

	return p
}

// 输入框默认内容
func (p *Text) SetDefaultValue(defaultValue interface{}) *Text {
	p.DefaultValue = defaultValue

	return p
}

// 输入框的 id
func (p *Text) SetId(id string) *Text {
	p.Id = id

	return p
}

// 最大长度
func (p *Text) SetMaxLength(maxLength int) *Text {
	p.MaxLength = maxLength

	return p
}

// 是否展示字数
func (p *Text) SetShowCount(showCount bool) *Text {
	p.ShowCount = showCount

	return p
}

// 设置校验状态，'error' | 'warning'
func (p *Text) SetStatus(status string) *Text {
	p.Status = status

	return p
}

// 输入框占位文本
func (p *Text) SetPlaceholder(placeholder string) *Text {
	p.Placeholder = placeholder

	return p
}

// 带有前缀图标的 input
func (p *Text) SetPrefix(prefix interface{}) *Text {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Text) SetSize(size string) *Text {
	p.Size = size

	return p
}

// 带有后缀图标的 input
func (p *Text) SetSuffix(suffix interface{}) *Text {
	p.Suffix = suffix

	return p
}

// 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="textarea")
func (p *Text) SetType(Type string) *Text {
	p.Type = Type

	return p
}
