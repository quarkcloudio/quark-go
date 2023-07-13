package searchbar

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

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

	Value           interface{} `json:"value,omitempty"`
	Shape           string      `json:"shape,omitempty"`
	MaxLength       int         `json:"maxLength,omitempty"`
	InputType       string      `json:"inputType,omitempty"`
	Placeholder     string      `json:"placeholder,omitempty"`
	Clearable       bool        `json:"clearable,omitempty"`
	ClearIcon       interface{} `json:"clearIcon,omitempty"`
	Background      string      `json:"background,omitempty"`
	InputBackground string      `json:"inputBackground,omitempty"`
	ConfirmType     string      `json:"confirmType,omitempty"`
	Autofocus       bool        `json:"autofocus,omitempty"`
	FocusStyle      interface{} `json:"focusStyle,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	Readonly        bool        `json:"readonly,omitempty"`
	InputAlign      string      `json:"inputAlign,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "searchbarField"
	p.SetKey("searchbar", component.DEFAULT_CRYPT)

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

// 默认值
func (p *Component) SetValue(value []interface{}) *Component {
	p.Value = value

	return p
}

// 搜索框形状，可选值为 square round
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 最大输入长度
func (p *Component) SetMaxLength(maxLength int) *Component {
	p.MaxLength = maxLength

	return p
}

// 输入框类型
func (p *Component) SetInputType(inputType string) *Component {
	p.InputType = inputType

	return p
}

// 输入框默认暗纹
func (p *Component) SetPlaceholder(placeholder string) *Component {
	p.Placeholder = placeholder

	return p
}

// 是否展示清除按钮
func (p *Component) SetClearable(clearable bool) *Component {
	p.Clearable = clearable

	return p
}

// 自定义清除按钮图标（默认使用 @nutui/nutui-icons）
func (p *Component) SetClearIcon(clearIcon interface{}) *Component {
	p.ClearIcon = clearIcon

	return p
}

// 间距
func (p *Component) SetBackground(background string) *Component {
	p.Background = background

	return p
}

// 输入框内部背景
func (p *Component) SetInputBackground(inputBackground string) *Component {
	p.InputBackground = inputBackground

	return p
}

// 键盘右下角按钮的文字，仅在type='text'时生效，可选值 send：发送、search：搜索、next：下一个、go：前往、done：完成
func (p *Component) SetConfirmType(confirmType string) *Component {
	p.ConfirmType = confirmType

	return p
}

// 是否自动聚焦
func (p *Component) SetAutofocus(autofocus bool) *Component {
	p.Autofocus = autofocus

	return p
}

// 聚焦时搜索框样式
func (p *Component) SetFocusStyle(focusStyle interface{}) *Component {
	p.FocusStyle = focusStyle

	return p
}

// 是否禁用输入框
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 输入框只读
func (p *Component) SetReadonly(readonly bool) *Component {
	p.Readonly = readonly

	return p
}

// 对齐方式，可选 left center right
func (p *Component) SetInputAlign(inputAlign string) *Component {
	p.InputAlign = inputAlign

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "searchbarField"

	return p
}
