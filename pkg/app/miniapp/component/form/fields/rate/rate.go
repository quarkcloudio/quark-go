package radio

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
	IconSize int         `json:"iconSize,omitempty"`
	Shape    string      `json:"shape,omitempty"`
}

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

	Value       interface{} `json:"value,omitempty"`
	Count       int         `json:"count,omitempty"`
	ActiveColor string      `json:"activeColor,omitempty"`
	VoidColor   string      `json:"voidColor,omitempty"`
	AllowHalf   bool        `json:"allowHalf,omitempty"`
	Readonly    bool        `json:"readonly,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Spacing     int         `json:"spacing,omitempty"`
	Size        int         `json:"size,omitempty"`
	CustomIcon  string      `json:"customIcon,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "radioField"
	p.SetKey("radio", component.DEFAULT_CRYPT)

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

// star 总数
func (p *Component) SetCount(count int) *Component {
	p.Count = count

	return p
}

// 图标选中颜色
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor

	return p
}

// 最小值
func (p *Component) SetVoidColor(voidColor string) *Component {
	p.VoidColor = voidColor

	return p
}

// 是否半星
func (p *Component) SetAllowHalf(allowHalf bool) *Component {
	p.AllowHalf = allowHalf

	return p
}

// 是否只读
func (p *Component) SetReadonly(readonly bool) *Component {
	p.Readonly = readonly

	return p
}

// 是否禁用
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 间距
func (p *Component) SetSpacing(spacing int) *Component {
	p.Spacing = spacing

	return p
}

// Icon 尺寸大小，如 20px 2em 2rem
func (p *Component) SetSize(size int) *Component {
	p.Size = size

	return p
}

// 自定义 Icon, 传入
func (p *Component) SetCustomIcon(customIcon string) *Component {
	p.CustomIcon = customIcon

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "radioField"

	return p
}
