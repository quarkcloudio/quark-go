package input

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Name              string      `json:"name"`
	Required          bool        `json:"required"`
	Prop              string      `json:"prop"`
	Rules             interface{} `json:"rules,omitempty"`
	Label             string      `json:"label,omitempty"`
	LabelWidth        int         `json:"labelWidth,omitempty"`
	LabelAlign        string      `json:"labelAlign,omitempty"`
	BodyAlign         string      `json:"bodyAlign,omitempty"`
	ErrorMessageAlign string      `json:"errorMessageAlign,omitempty"`
	ShowErrorLine     bool        `json:"showErrorLine,omitempty"`
	ShowErrorMessage  bool        `json:"showErrorMessage,omitempty"`

	Value         interface{} `json:"value"`
	InputWidth    string      `json:"inputWidth"`
	ButtonSize    string      `json:"buttonSize"`
	Min           int         `json:"min"`
	Max           int         `json:"max"`
	Step          int         `json:"step"`
	DecimalPlaces int         `json:"decimalPlaces"`
	Disabled      bool        `json:"disabled"`
	Readonly      bool        `json:"readonly"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "inputNumberField"
	p.SetKey("inputNumber", component.DEFAULT_CRYPT)

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

// 输入框的初始内容
func (p *Component) SetValue(value interface{}) *Component {
	p.Value = value

	return p
}

// 输入框宽度
func (p *Component) SetInputWidth(inputWidth string) *Component {
	p.InputWidth = inputWidth

	return p
}

// 是否是密码类型
func (p *Component) SetButtonSize(buttonSize string) *Component {
	p.ButtonSize = buttonSize

	return p
}

// 最小值限制
func (p *Component) SetMin(min int) *Component {
	p.Min = min

	return p
}

// 最大值限制
func (p *Component) SetMax(max int) *Component {
	p.Max = max

	return p
}

// 步长
func (p *Component) SetStep(step int) *Component {
	p.Step = step

	return p
}

// 设置保留的小数位
func (p *Component) SetDecimalPlaces(decimalPlaces int) *Component {
	p.DecimalPlaces = decimalPlaces

	return p
}

// 禁用所有功能
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 只读状态禁用输入框操作行为
func (p *Component) SetReadonly(readonly bool) *Component {
	p.Readonly = readonly

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "inputNumberField"

	return p
}
