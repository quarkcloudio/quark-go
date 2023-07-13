package picker

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

	Value            interface{} `json:"value"`
	Columns          []*Column   `json:"columns"`
	Title            string      `json:"title"`
	CancelText       string      `json:"cancelText"`
	OkText           string      `json:"okText"`
	ThreeDimensional bool        `json:"threeDimensional"`
	SwipeDuration    int         `json:"swipeDuration"`
	VisibleOptionNum int         `json:"visibleOptionNum"`
	OptionHeight     int         `json:"optionHeight"`
	ShowToolbar      bool        `json:"showToolbar"`
}

type Column struct {
	Text     string    `json:"text,omitempty"`
	Value    string    `json:"value,omitempty"`
	Children []*Column `json:"children,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "pickerField"
	p.SetKey("picker", component.DEFAULT_CRYPT)
	p.Value = 0

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

// 对象数组，配置每一列显示的数据
func (p *Component) SetColumns(columns []*Column) *Component {
	p.Columns = columns

	return p
}

// 设置标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 取消按钮文案
func (p *Component) SetCancelText(cancelText string) *Component {
	p.CancelText = cancelText

	return p
}

// 确定按钮文案
func (p *Component) SetOkText(okText string) *Component {
	p.OkText = okText

	return p
}

// 是否开启3D效果
func (p *Component) SetThreeDimensional(threeDimensional bool) *Component {
	p.ThreeDimensional = threeDimensional

	return p
}

// 惯性滚动时长
func (p *Component) SetSwipeDuration(swipeDuration int) *Component {
	p.SwipeDuration = swipeDuration

	return p
}

// 可见的选项个数
func (p *Component) SetVisibleOptionNum(visibleOptionNum int) *Component {
	p.VisibleOptionNum = visibleOptionNum

	return p
}

// 选项高度
func (p *Component) SetOptionHeight(optionHeight int) *Component {
	p.OptionHeight = optionHeight

	return p
}

// 是否显示顶部导航
func (p *Component) SetShowToolbar(showToolbar bool) *Component {
	p.ShowToolbar = showToolbar

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "pickerField"

	return p
}
