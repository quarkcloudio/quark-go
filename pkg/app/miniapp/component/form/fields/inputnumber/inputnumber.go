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

	Value       interface{} `json:"value"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
	RandomKeys  bool        `json:"randomKeys"`
	CustomKey   []string    `json:"customKey"`
	Overlay     bool        `json:"overlay"`
	Maxlength   int         `json:"maxlength"`
	ConfirmText string      `json:"confirmText"`
	PopClass    string      `json:"popClass"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "numberKeyboardField"
	p.SetKey("numberKeyboard", component.DEFAULT_CRYPT)

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

// 键盘标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 键盘模式
func (p *Component) SetType(numberkeyboardType string) *Component {
	p.Type = numberkeyboardType

	return p
}

// 随机数
func (p *Component) SetRandomKeys(randomKeys bool) *Component {
	p.RandomKeys = randomKeys

	return p
}

// 自定义键盘额外的键
func (p *Component) SetCustomKey(customKey []string) *Component {
	p.CustomKey = customKey

	return p
}

// 是否显示遮罩
func (p *Component) SetOverlay(overlay bool) *Component {
	p.Overlay = overlay

	return p
}

// 输入值最大长度，结合 v-model 使用
func (p *Component) SetMaxlength(maxlength int) *Component {
	p.Maxlength = maxlength

	return p
}

// 自定义完成按钮文字，如"支付"，"下一步"，"提交"等
func (p *Component) SetConfirmText(confirmText string) *Component {
	p.ConfirmText = confirmText

	return p
}

// 自定义弹框类名
func (p *Component) SetPopClass(popClass string) *Component {
	p.PopClass = popClass

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "numberKeyboardField"

	return p
}
