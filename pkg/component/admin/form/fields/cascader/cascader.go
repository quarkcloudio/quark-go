package cascader

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type FieldNames struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Children string `json:"children"`
}

type Option struct {
	Label    string      `json:"label,omitempty"`
	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled,omitempty"`
	Children []*Option   `json:"children,omitempty"`
	// 标记是否为叶子节点，设置了 `loadData` 时有效
	// 设为 `false` 时会强制标记为父节点，即使当前节点没有 children，也会显示展开图标
	IsLeaf bool `json:"isLeaf,omitempty"`
}

type Cascader struct {
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

	Api                  string            `json:"api,omitempty"` // 获取数据接口
	Ignore               bool              `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules                []*rule.Rule      `json:"-"`             // 全局校验规则
	RuleMessages         map[string]string `json:"-"`             // 全局校验提示信息
	CreationRules        []string          `json:"-"`             // 创建页校验规则
	CreationRuleMessages map[string]string `json:"-"`             // 创建页校验提示信息
	UpdateRules          []string          `json:"-"`             // 编辑页校验规则
	UpdateRuleMessages   map[string]string `json:"-"`             // 编辑页校验提示信息
	FrontendRules        interface{}       `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When                 *when.When        `json:"when"`          //
	WhenItem             []*when.Item      `json:"-"`             //
	ShowOnIndex          bool              `json:"-"`             // 在列表页展示
	ShowOnDetail         bool              `json:"-"`             // 在详情页展示
	ShowOnCreation       bool              `json:"-"`             // 在创建页面展示
	ShowOnUpdate         bool              `json:"-"`             // 在编辑页面展示
	ShowOnExport         bool              `json:"-"`             // 在导出的Excel上展示
	ShowOnImport         bool              `json:"-"`             // 在导入Excel上展示
	Editable             bool              `json:"-"`             // 表格上是否可编辑
	Column               *table.Column     `json:"-"`             // 表格列
	Callback             interface{}       `json:"-"`             // 回调函数

	AllowClear              bool                   `json:"allowClear,omitempty"`              // 是否支持清除，默认true
	AutoFocus               bool                   `json:"autoFocus,omitempty"`               // 自动获取焦点，默认false
	Bordered                bool                   `json:"bordered,omitempty"`                // 是否有边框，默认true
	ClearIcon               interface{}            `json:"clearIcon,omitempty"`               // 自定义的选择框清空图标
	ChangeOnSelect          bool                   `json:"changeOnSelect,omitempty"`          // （单选时生效）当此项为 true 时，点选每级菜单选项值都会发生变化，默认false
	ClassName               string                 `json:"className,omitempty"`               // 自定义类名
	DefaultValue            interface{}            `json:"defaultValue,omitempty"`            // 默认的选中项
	Disabled                interface{}            `json:"disabled,omitempty"`                // 禁用
	PopupClassName          string                 `json:"popupClassName,omitempty"`          // 自定义类名
	ExpandIcon              interface{}            `json:"expandIcon,omitempty"`              // 自定义次级菜单展开图标
	ExpandTrigger           string                 `json:"expandTrigger,omitempty"`           // 次级菜单的展开方式，可选 'click' 和 'hover'
	FieldNames              *FieldNames            `json:"fieldNames,omitempty"`              // 自定义 options 中 label value children 的字段
	MaxTagCount             int                    `json:"maxTagCount,omitempty"`             // 最多显示多少个 tag，响应式模式会对性能产生损耗
	MaxTagPlaceholder       string                 `json:"maxTagPlaceholder,omitempty"`       // 隐藏 tag 时显示的内容
	MaxTagTextLength        int                    `json:"maxTagTextLength,omitempty"`        // 最大显示的 tag 文本长度
	NotFoundContent         string                 `json:"notFoundContent,omitempty"`         // 当下拉列表为空时显示的内容
	Open                    bool                   `json:"open,omitempty"`                    // 控制浮层显隐
	Options                 []*Option              `json:"options,omitempty"`                 // 可选项数据源
	Placeholder             string                 `json:"placeholder,omitempty"`             // 输入框占位文本
	Placement               string                 `json:"placement,omitempty"`               // 浮层预设位置，bottomLeft bottomRight topLeft topRight
	ShowSearch              bool                   `json:"showSearch,omitempty"`              // 在选择框中显示搜索框
	Size                    string                 `json:"size,omitempty"`                    // 输入框大小，large | middle | small
	Status                  string                 `json:"status,omitempty"`                  // 设置校验状态，'error' | 'warning'
	Style                   map[string]interface{} `json:"style,omitempty"`                   // 自定义样式
	SuffixIcon              interface{}            `json:"suffixIcon,omitempty"`              // 自定义的选择框后缀图标
	Value                   interface{}            `json:"value,omitempty"`                   // 指定选中项,string[] | number[]
	Multiple                bool                   `json:"multiple,omitempty"`                // 支持多选节点
	ShowCheckedStrategy     string                 `json:"showCheckedStrategy,omitempty"`     // 定义选中项回填的方式。Cascader.SHOW_CHILD: 只显示选中的子节点。Cascader.SHOW_PARENT: 只显示父节点（当父节点下所有子节点都选中时）。Cascader.SHOW_PARENT | Cascader.SHOW_CHILD
	RemoveIcon              interface{}            `json:"removeIcon,omitempty"`              // 自定义的多选框清除图标
	SearchValue             string                 `json:"searchValue,omitempty"`             // 设置搜索的值，需要与 showSearch 配合使用
	DropdownMenuColumnStyle interface{}            `json:"dropdownMenuColumnStyle,omitempty"` // 下拉菜单列的样式
}

// 初始化组件
func New() *Cascader {
	return (&Cascader{}).Init()
}

// 初始化
func (p *Cascader) Init() *Cascader {
	p.Component = "cascaderField"
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
	p.SetWidth(400)

	return p
}

// 设置Key
func (p *Cascader) SetKey(key string, crypt bool) *Cascader {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// Set style.
func (p *Cascader) SetStyle(style map[string]interface{}) *Cascader {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Cascader) SetTooltip(tooltip string) *Cascader {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Cascader) SetWidth(width interface{}) *Cascader {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Cascader) SetColon(colon bool) *Cascader {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Cascader) SetExtra(extra string) *Cascader {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Cascader) SetHasFeedback(hasFeedback bool) *Cascader {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Cascader) SetHelp(help string) *Cascader {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Cascader) SetNoStyle() *Cascader {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Cascader) SetLabel(label string) *Cascader {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Cascader) SetLabelAlign(align string) *Cascader {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Cascader) SetLabelCol(col interface{}) *Cascader {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Cascader) SetName(name string) *Cascader {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Cascader) SetRequired() *Cascader {
	p.Required = true
	return p
}

// 解析成前端验证规则
func (p *Cascader) parseFrontendRules(rules []string, messages map[string]string) []map[string]interface{} {
	result := []map[string]interface{}{}
	values := []string{}
	rule := ""

	for _, v := range rules {
		if strings.Contains(v, ":") {
			values = strings.Split(v, ":")
			rule = values[0]
		} else {
			rule = v
		}

		data := map[string]interface{}{}

		switch rule {
		case "required":
			data = map[string]interface{}{
				"required": true,
				"message":  messages["required"],
			}
		case "min":
			min, _ := strconv.Atoi(values[1])

			data = map[string]interface{}{
				"min":     min,
				"message": messages["min"],
			}
		case "max":
			max, _ := strconv.Atoi(values[1])

			data = map[string]interface{}{
				"max":     max,
				"message": messages["max"],
			}

		case "email":
			data = map[string]interface{}{
				"type":    "email",
				"message": messages["email"],
			}

		case "numeric":
			data = map[string]interface{}{
				"type":    "number",
				"message": messages["numeric"],
			}

		case "url":
			data = map[string]interface{}{
				"type":    "url",
				"message": messages["url"],
			}

		case "integer":
			data = map[string]interface{}{
				"type":    "integer",
				"message": messages["integer"],
			}

		case "date":
			data = map[string]interface{}{
				"type":    "date",
				"message": messages["date"],
			}
		case "boolean":
			data = map[string]interface{}{
				"type":    "boolean",
				"message": messages["boolean"],
			}
		}

		if len(data) > 0 {
			result = append(result, data)
		}
	}

	return result
}

// 校验规则，设置字段的校验逻辑
func (p *Cascader) SetRules(rules []*rule.Rule) *Cascader {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Cascader) SetCreationRules(rules []string, messages map[string]string) *Cascader {
	p.CreationRules = rules
	p.CreationRuleMessages = messages

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Cascader) SetUpdateRules(rules []string, messages map[string]string) *Cascader {
	p.UpdateRules = rules
	p.UpdateRuleMessages = messages

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Cascader) SetValuePropName(valuePropName string) *Cascader {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Cascader) SetWrapperCol(col interface{}) *Cascader {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Cascader) SetValue(value interface{}) *Cascader {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Cascader) SetDefault(value interface{}) *Cascader {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Cascader) SetDisabled(disabled bool) *Cascader {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Cascader) SetIgnore(ignore bool) *Cascader {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Cascader) SetWhen(value ...any) *Cascader {
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
func (p *Cascader) HideFromIndex(callback bool) *Cascader {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Cascader) HideFromDetail(callback bool) *Cascader {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Cascader) HideWhenCreating(callback bool) *Cascader {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Cascader) HideWhenUpdating(callback bool) *Cascader {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Cascader) HideWhenExporting(callback bool) *Cascader {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Cascader) HideWhenImporting(callback bool) *Cascader {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Cascader) OnIndexShowing(callback bool) *Cascader {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Cascader) OnDetailShowing(callback bool) *Cascader {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Cascader) ShowOnCreating(callback bool) *Cascader {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Cascader) ShowOnUpdating(callback bool) *Cascader {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Cascader) ShowOnExporting(callback bool) *Cascader {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Cascader) ShowOnImporting(callback bool) *Cascader {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Cascader) OnlyOnIndex() *Cascader {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Cascader) OnlyOnDetail() *Cascader {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Cascader) OnlyOnForms() *Cascader {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Cascader) OnlyOnExport() *Cascader {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Cascader) OnlyOnImport() *Cascader {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Cascader) ExceptOnForms() *Cascader {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Cascader) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Cascader) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Cascader) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Cascader) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Cascader) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Cascader) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Cascader) SetEditable(editable bool) *Cascader {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Cascader) SetColumn(f func(column *table.Column) *table.Column) *Cascader {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Cascader) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Cascader) SetCallback(closure func() interface{}) *Cascader {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Cascader) GetCallback() interface{} {
	return p.Callback
}

// 设置属性
func (p *Cascader) SetOptions(options []*Option) *Cascader {
	p.Options = options

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Cascader) SetSize(size string) *Cascader {
	p.Size = size

	return p
}

// 可以点击清除图标删除内容
func (p *Cascader) SetAllowClear(allowClear bool) *Cascader {
	p.AllowClear = allowClear

	return p
}

// 获取数据接口
func (p *Cascader) SetApi(api string) *Cascader {
	p.Api = api
	return p
}
