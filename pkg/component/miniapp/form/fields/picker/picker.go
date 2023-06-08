package picker

import "github.com/quarkcms/quark-go/pkg/component/miniapp/component"

type Component struct {
	component.Element
	Name              string      `json:"name"`
	Required          bool        `json:"required,omitempty"`
	Prop              string      `json:"prop,omitempty"`
	Rules             interface{} `json:"rules,omitempty"`
	Label             string      `json:"label,omitempty"`
	LabelWidth        int         `json:"labelWidth,omitempty"`
	LabelAlign        string      `json:"labelAlign,omitempty"`
	BodyAlign         string      `json:"bodyAlign,omitempty"`
	ErrorMessageAlign string      `json:"errorMessageAlign,omitempty"`
	ShowErrorLine     bool        `json:"showErrorLine,omitempty"`
	ShowErrorMessage  bool        `json:"showErrorMessage,omitempty"`

	Value        interface{}   `json:"value"`
	Start        string        `json:"start"`
	End          string        `json:"end"`
	Fields       string        `json:"fields"`
	CustomItem   string        `json:"customItem"`
	Mode         string        `json:"mode"`
	Range        []interface{} `json:"range"`
	RangeKey     string        `json:"RangeKey"`
	SelectorType string        `json:"SelectorType"`
	Disabled     bool          `json:"disabled"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "pickerField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Mode = "selector"
	p.Value = 0
	p.SelectorType = "auto"

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetProp(prop string) *Component {
	p.Name = prop
	p.Prop = prop

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetName(name string) *Component {
	p.Name = name
	p.Prop = name

	return p
}

// 定义校验规则
func (p *Component) SetRules(rules interface{}) *Component {
	p.Rules = rules

	return p
}

// 是否显示必填字段的标签旁边的红色星号
func (p *Component) SetRequired(required bool) *Component {
	p.Required = required

	return p
}

// 输入框左边的文字提示
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 表单项 label 宽度，默认单位为px
func (p *Component) SetLabelWidth(labelWidth int) *Component {
	p.LabelWidth = labelWidth

	return p
}

// 表单项 label 对齐方式，可选值为 center right
func (p *Component) SetLabelAlign(labelAlign string) *Component {
	p.LabelAlign = labelAlign

	return p
}

// 右侧插槽对齐方式，可选值为 center right
func (p *Component) SetBodyAlign(bodyAlign string) *Component {
	p.BodyAlign = bodyAlign

	return p
}

// 错误提示文案对齐方式，可选值为 center right
func (p *Component) SetErrorMessageAlign(errorMessageAlign string) *Component {
	p.ErrorMessageAlign = errorMessageAlign

	return p
}

// 是否在校验不通过时标红输入框
func (p *Component) SetShowErrorLine(showErrorLine bool) *Component {
	p.ShowErrorLine = showErrorLine

	return p
}

// 是否在校验不通过时在输入框下方展示错误提示
func (p *Component) SetShowErrorMessage(showErrorMessage bool) *Component {
	p.ShowErrorMessage = showErrorMessage

	return p
}

// value 的值表示选择了 range 中的第几个（下标从 0 开始）
func (p *Component) SetValue(value int) *Component {
	p.Value = value

	return p
}

// 选择器模式：selector,multiSelector,time,date,region
func (p *Component) SetMode(mode string) *Component {
	p.Mode = mode

	return p
}

// mode为 selector 或 multiSelector 时，range 有效
func (p *Component) SetRange(pickerRange []interface{}) *Component {
	p.Range = pickerRange

	return p
}

// 表示有效时间范围的开始
func (p *Component) SetStart(start string) *Component {
	p.Start = start

	return p
}

// 表示有效时间范围的结束
func (p *Component) SetEnd(end string) *Component {
	p.End = end

	return p
}

// 有效值 year、month、day，表示选择器的粒度，默认为 day，App 端未配置此项时使用系统 UI
func (p *Component) SetFields(fields string) *Component {
	p.Fields = fields

	return p
}

// 有效值 year、month、day，表示选择器的粒度，默认为 day，App 端未配置此项时使用系统 UI
func (p *Component) SetCustomItem(customItem string) *Component {
	p.CustomItem = customItem

	return p
}

// 当 range 是一个 Array＜Object＞ 时，通过 range-key 来指定 Object 中 key 的值作为选择器显示内容
func (p *Component) SetRangeKey(rangeKey string) *Component {
	p.RangeKey = rangeKey

	return p
}

// 大屏时UI类型，支持 picker、select、auto，默认在 iPad 以 picker 样式展示而在 PC 以 select 样式展示
func (p *Component) SetSelectorType(selectorType string) *Component {
	p.SelectorType = selectorType

	return p
}

// 是否禁用
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "pickerField"

	return p
}
