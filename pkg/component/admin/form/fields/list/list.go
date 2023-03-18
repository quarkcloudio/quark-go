package list

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type List struct {
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

	DefaultValue        interface{} `json:"defaultValue,omitempty"` // 默认选中的选项
	Disabled            bool        `json:"disabled,omitempty"`     // 整组失效
	Value               interface{} `json:"value,omitempty"`        // 指定选中项,string[] | number[]
	Items               interface{} `json:"items"`                  // 列表项
	ButtonText          string      `json:"buttonText"`             // 按钮文字
	ButtonPosition      string      `json:"buttonPosition"`         // 按钮位置
	AlwaysShowItemLabel bool        `json:"alwaysShowItemLabel"`    // Item 中总是展示 label
}

// 初始化组件
func New() *List {
	return (&List{}).Init()
}

// 初始化
func (p *List) Init() *List {
	p.Component = "listField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()

	p.ButtonText = "添加一行数据"
	p.ButtonPosition = "top"
	p.AlwaysShowItemLabel = true

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *List) SetKey(key string, crypt bool) *List {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *List) SetTooltip(tooltip string) *List {
	p.Tooltip = tooltip

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *List) SetColon(colon bool) *List {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *List) SetExtra(extra string) *List {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *List) SetHasFeedback(hasFeedback bool) *List {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *List) SetHelp(help string) *List {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *List) SetNoStyle() *List {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *List) SetLabel(label string) *List {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *List) SetLabelAlign(align string) *List {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *List) SetLabelCol(col interface{}) *List {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *List) SetName(name string) *List {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *List) SetRequired() *List {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *List) GetFrontendRules(path string) *List {
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
func (p *List) SetRules(rules []*rule.Rule) *List {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *List) SetCreationRules(rules []*rule.Rule) *List {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *List) SetUpdateRules(rules []*rule.Rule) *List {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *List) SetValuePropName(valuePropName string) *List {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *List) SetWrapperCol(col interface{}) *List {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *List) SetValue(value interface{}) *List {
	p.Value = value
	return p
}

// 设置默认值。
func (p *List) SetDefault(value interface{}) *List {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *List) SetDisabled(disabled bool) *List {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *List) SetIgnore(ignore bool) *List {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *List) SetWhen(value ...any) *List {
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
func (p *List) HideFromIndex(callback bool) *List {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *List) HideFromDetail(callback bool) *List {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *List) HideWhenCreating(callback bool) *List {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *List) HideWhenUpdating(callback bool) *List {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export List.
func (p *List) HideWhenExporting(callback bool) *List {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import List.
func (p *List) HideWhenImporting(callback bool) *List {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *List) OnIndexShowing(callback bool) *List {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *List) OnDetailShowing(callback bool) *List {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *List) ShowOnCreating(callback bool) *List {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *List) ShowOnUpdating(callback bool) *List {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export List.
func (p *List) ShowOnExporting(callback bool) *List {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import List.
func (p *List) ShowOnImporting(callback bool) *List {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *List) OnlyOnIndex() *List {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *List) OnlyOnDetail() *List {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *List) OnlyOnForms() *List {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export List.
func (p *List) OnlyOnExport() *List {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import List.
func (p *List) OnlyOnImport() *List {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *List) ExceptOnForms() *List {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *List) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *List) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *List) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *List) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *List) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *List) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *List) SetEditable(editable bool) *List {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *List) SetColumn(f func(column *table.Column) *table.Column) *List {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *List) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *List) SetCallback(closure func() interface{}) *List {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *List) GetCallback() interface{} {
	return p.Callback
}

// 按钮名称
func (p *List) SetButton(text string, position string) *List {
	p.ButtonText = text
	p.ButtonPosition = position

	return p
}

// 表单项
func (p *List) SetItem(callback interface{}) *List {
	getCallback := callback.(func() interface{})

	p.Items = getCallback()

	return p
}

// Item 中总是展示 label
func (p *List) SetAlwaysShowItemLabel(alwaysShowItemLabel bool) *List {
	p.AlwaysShowItemLabel = alwaysShowItemLabel

	return p
}

// 上传的api接口
func (p *List) SetApi(api string) *List {
	p.Api = api
	return p
}
