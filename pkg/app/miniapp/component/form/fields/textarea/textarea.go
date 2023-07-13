package textarea

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
	Placeholder string      `json:"placeholder"`
	MaxLength   int         `json:"maxLength"`
	Rows        int         `json:"rows"`
	LimitShow   bool        `json:"limitShow"`
	Autosize    interface{} `json:"autosize"`
	TextAlign   string      `json:"textAlign"`
	Readonly    bool        `json:"readonly"`
	Disabled    bool        `json:"disabled"`
	Autofocus   bool        `json:"autofocus"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "textareaField"
	p.SetKey("textarea", component.DEFAULT_CRYPT)

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

// 设置占位提示文字
func (p *Component) SetPlaceholder(placeholder string) *Component {
	p.Placeholder = placeholder

	return p
}

// 限制最长输入字符
func (p *Component) SetMaxLength(maxLength int) *Component {
	p.MaxLength = maxLength

	return p
}

// textarea的高度，优先级高于autosize属性 仅支持 H5
func (p *Component) SetRows(rows int) *Component {
	p.Rows = rows

	return p
}

// textarea是否展示输入字符。须配合max-length使用
func (p *Component) SetLimitShow(limitShow bool) *Component {
	p.LimitShow = limitShow

	return p
}

// 是否自适应内容高度，也可传入对象,如 { maxHeight: 200, minHeight: 100 }，单位为px
func (p *Component) SetAutosize(autosize interface{}) *Component {
	p.Autosize = autosize

	return p
}

// 文本位置,可选值left,center,right
func (p *Component) SetTextAlign(textAlign string) *Component {
	p.TextAlign = textAlign

	return p
}

// 只读属性
func (p *Component) SetReadonly(readonly bool) *Component {
	p.Readonly = readonly

	return p
}

// 禁用属性
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 自动获取焦点
func (p *Component) SetAutofocus(autofocus bool) *Component {
	p.Autofocus = autofocus

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "textareaField"

	return p
}
