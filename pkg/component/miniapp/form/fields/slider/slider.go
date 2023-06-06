package slider

import "github.com/quarkcms/quark-go/pkg/component/miniapp/component"

type Component struct {
	component.Element
	Required          bool        `json:"required"`
	Prop              string      `json:"prop"`
	Rules             interface{} `json:"rules"`
	Label             string      `json:"label"`
	LabelWidth        int         `json:"labelWidth"`
	LabelAlign        string      `json:"labelAlign"`
	BodyAlign         string      `json:"bodyAlign"`
	ErrorMessageAlign string      `json:"errorMessageAlign"`
	ShowErrorLine     bool        `json:"showErrorLine"`
	ShowErrorMessage  bool        `json:"showErrorMessage"`

	Value           int    `json:"value"`
	Min             int    `json:"min"`
	Max             int    `json:"max"`
	Step            int    `json:"step"`
	Disabled        bool   `json:"disabled"`
	ActiveColor     string `json:"activeColor"`
	BackgroundColor string `json:"backgroundColor"`
	BlockSize       int    `json:"blockSize"`
	BlockColor      string `json:"blockColor"`
	ShowValue       bool   `json:"showValue"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "sliderField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Min = 0
	p.Max = 100
	p.Step = 1
	p.Value = 0
	p.BackgroundColor = "#e9e9e9"
	p.BlockSize = 28
	p.BlockColor = "#ffffff"

	return p
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetProp(prop string) *Component {
	p.Prop = prop

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetName(name string) *Component {
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

// 默认值
func (p *Component) SetValue(value int) *Component {
	p.Value = value

	return p
}

// 最小值
func (p *Component) SetMin(min int) *Component {
	p.Min = min

	return p
}

// 最大值
func (p *Component) SetMax(max int) *Component {
	p.Max = max

	return p
}

// 步长，取值必须大于 0，并且可被(max - min)整除
func (p *Component) SetStep(step int) *Component {
	p.Step = step

	return p
}

// 是否禁用
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 滑块左侧已选择部分的线条颜色
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor

	return p
}

// 滑块右侧背景条的颜色
func (p *Component) SetBackgroundColor(backgroundColor string) *Component {
	p.BackgroundColor = backgroundColor

	return p
}

// 滑块的大小，取值范围为 12 - 28
func (p *Component) SetBlockSize(blockSize int) *Component {
	p.BlockSize = blockSize

	return p
}

// 滑块的颜色
func (p *Component) SetBlockColor(blockColor string) *Component {
	p.BlockColor = blockColor

	return p
}

// 是否显示当前 value
func (p *Component) SetShowValue(showValue bool) *Component {
	p.ShowValue = showValue

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "sliderField"

	return p
}
