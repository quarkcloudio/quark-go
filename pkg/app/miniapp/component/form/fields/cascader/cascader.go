package cascader

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

	Value             interface{} `json:"value,omitempty"`
	Options           []*Option   `json:"options,omitempty"`
	Lazy              bool        `json:"lazy,omitempty"`
	ValueKey          string      `json:"valueKey,omitempty"`
	TextKey           string      `json:"textKey,omitempty"`
	ChildrenKey       string      `json:"childrenKey,omitempty"`
	ConvertConfig     interface{} `json:"convertConfig,omitempty"`
	Title             string      `json:"title,omitempty"`
	CloseIconPosition string      `json:"closeIconPosition,omitempty"`
	Closeable         bool        `json:"closeable,omitempty"`
	Poppable          bool        `json:"poppable,omitempty"`
}

type Option struct {
	Value    string    `json:"value,omitempty"`
	Text     string    `json:"text,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
	Children []*Option `json:"children,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "cascaderField"
	p.SetKey("cascader", component.DEFAULT_CRYPT)

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

// 级联数据
func (p *Component) SetOptions(options []*Option) *Component {
	p.Options = options

	return p
}

// 是否开启动态加载
func (p *Component) SetLazy(lazy bool) *Component {
	p.Lazy = lazy

	return p
}

// 自定义 options 结构中 value 的字段
func (p *Component) SetValueKey(valueKey string) *Component {
	p.ValueKey = valueKey

	return p
}

// 自定义 options 结构中 text 的字段
func (p *Component) SetTextKey(textKey string) *Component {
	p.TextKey = textKey

	return p
}

// 自定义 options 结构中 children 的字段
func (p *Component) SetChildrenKey(childrenKey string) *Component {
	p.ChildrenKey = childrenKey

	return p
}

// 当 options 为可转换为树形结构的扁平结构时，配置转换规则
func (p *Component) SetConvertConfig(convertConfig interface{}) *Component {
	p.ConvertConfig = convertConfig

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 取消按钮位置，继承 Popup 组件
func (p *Component) SetCloseIconPosition(closeIconPosition string) *Component {
	p.CloseIconPosition = closeIconPosition

	return p
}

// 是否显示关闭按钮，继承 Popup 组件
func (p *Component) SetCloseable(closeable bool) *Component {
	p.Closeable = closeable

	return p
}

// 是否需要弹层展示（设置为 false 后，title 失效）
func (p *Component) SetPoppable(poppable bool) *Component {
	p.Poppable = poppable

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "cascaderField"

	return p
}
