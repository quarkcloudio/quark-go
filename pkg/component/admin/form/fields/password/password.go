package password

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Password struct {
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
	DefaultValue     interface{}            `json:"defaultValue,omitempty"`     // 默认的选中项
	Disabled         interface{}            `json:"disabled,omitempty"`         // 禁用
	Id               string                 `json:"id,omitempty"`               // 输入框的 id
	MaxLength        int                    `json:"maxLength,omitempty"`        //最大长度
	ShowCount        bool                   `json:"showCount,omitempty"`        // 是否展示字数
	Status           string                 `json:"status,omitempty"`           // 设置校验状态,'error' | 'warning'
	Prefix           interface{}            `json:"prefix,omitempty"`           // 带有前缀图标的 input
	Size             string                 `json:"size,omitempty"`             // 控件大小。注：标准表单内的输入框大小限制为 middle，large | middle | small
	Suffix           interface{}            `json:"suffix,omitempty"`           // 带有后缀图标的 input
	Type             string                 `json:"type,omitempty"`             // 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="textarea")
	Value            interface{}            `json:"value,omitempty"`            // 指定选中项,string[] | number[]
	Placeholder      string                 `json:"placeholder,omitempty"`      // 占位符
	Style            map[string]interface{} `json:"style,omitempty"`            // 自定义样式
	VisibilityToggle bool                   `json:"visibilityToggle,omitempty"` // 是否显示切换按钮或者控制密码显隐
}

// 初始化组件
func New() *Password {
	return (&Password{}).Init()
}

// 初始化
func (p *Password) Init() *Password {
	p.Component = "passwordField"
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
func (p *Password) SetKey(key string, crypt bool) *Password {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *Password) SetStyle(style map[string]interface{}) *Password {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Password) SetTooltip(tooltip string) *Password {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Password) SetWidth(width interface{}) *Password {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Password) SetColon(colon bool) *Password {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Password) SetExtra(extra string) *Password {
	p.Extra = extra
	return p
}

// 配合 valiTextStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Password) SetHasFeedback(hasFeedback bool) *Password {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Password) SetHelp(help string) *Password {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Password) SetNoStyle() *Password {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Password) SetLabel(label string) *Password {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Password) SetLabelAlign(align string) *Password {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Password) SetLabelCol(col interface{}) *Password {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Password) SetName(name string) *Password {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Password) SetRequired() *Password {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Password) GetFrontendRules(path string) *Password {
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
func (p *Password) SetRules(rules []*rule.Rule) *Password {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Password) SetCreationRules(rules []*rule.Rule) *Password {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Password) SetUpTextRules(rules []*rule.Rule) *Password {
	p.UpTextRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Password) SetValuePropName(valuePropName string) *Password {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Password) SetWrapperCol(col interface{}) *Password {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Password) SetValue(value interface{}) *Password {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Password) SetDefault(value interface{}) *Password {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Password) SetDisabled(disabled bool) *Password {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Password) SetIgnore(ignore bool) *Password {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Password) SetWhen(value ...any) *Password {
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
func (p *Password) HideFromIndex(callback bool) *Password {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Password) HideFromDetail(callback bool) *Password {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Password) HideWhenCreating(callback bool) *Password {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Password) HideWhenUpdating(callback bool) *Password {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Password) HideWhenExporting(callback bool) *Password {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Password) HideWhenImporting(callback bool) *Password {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Password) OnIndexShowing(callback bool) *Password {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Password) OnDetailShowing(callback bool) *Password {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Password) ShowOnCreating(callback bool) *Password {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the upText view.
func (p *Password) ShowOnUpdating(callback bool) *Password {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Password) ShowOnExporting(callback bool) *Password {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Password) ShowOnImporting(callback bool) *Password {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Password) OnlyOnIndex() *Password {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Password) OnlyOnDetail() *Password {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Password) OnlyOnForms() *Password {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Password) OnlyOnExport() *Password {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Password) OnlyOnImport() *Password {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Password) ExceptOnForms() *Password {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Password) IsShownOnUpText() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Password) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Password) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Password) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Password) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Password) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Password) SetEditable(editable bool) *Password {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Password) SetColumn(f func(column *table.Column) *table.Column) *Password {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Password) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Password) SetCallback(closure func() interface{}) *Password {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Password) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Password) SetApi(api string) *Password {
	p.Api = api
	return p
}

// 带标签的 input，设置后置标签
func (p *Password) SetAddonAfter(addonAfter interface{}) *Password {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签
func (p *Password) SetAddonBefore(addonBefore interface{}) *Password {
	p.AddonBefore = addonBefore

	return p
}

// 可以点击清除图标删除内容
func (p *Password) SetAllowClear(allowClear bool) *Password {
	p.AllowClear = allowClear

	return p
}

// 是否有边框，默认true
func (p *Password) SetBordered(bordered bool) *Password {
	p.Bordered = bordered

	return p
}

// 输入框默认内容
func (p *Password) SetDefaultValue(defaultValue interface{}) *Password {
	p.DefaultValue = defaultValue

	return p
}

// 输入框的 id
func (p *Password) SetId(id string) *Password {
	p.Id = id

	return p
}

// 最大长度
func (p *Password) SetMaxLength(maxLength int) *Password {
	p.MaxLength = maxLength

	return p
}

// 是否展示字数
func (p *Password) SetShowCount(showCount bool) *Password {
	p.ShowCount = showCount

	return p
}

// 设置校验状态，'error' | 'warning'
func (p *Password) SetStatus(status string) *Password {
	p.Status = status

	return p
}

// 输入框占位文本
func (p *Password) SetPlaceholder(placeholder string) *Password {
	p.Placeholder = placeholder

	return p
}

// 带有前缀图标的 input
func (p *Password) SetPrefix(prefix interface{}) *Password {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Password) SetSize(size string) *Password {
	p.Size = size

	return p
}

// 带有后缀图标的 input
func (p *Password) SetSuffix(suffix interface{}) *Password {
	p.Suffix = suffix

	return p
}

// 声明 input 类型，同原生 input 标签的 type 属性，见：MDN(请直接使用 Input.TextArea 代替 type="textarea")
func (p *Password) SetType(Type string) *Password {
	p.Type = Type

	return p
}

// 是否显示切换按钮或者控制密码显隐
func (p *Password) SetVisibilityToggle(visibilityToggle bool) *Password {
	p.VisibilityToggle = visibilityToggle

	return p
}
