package checkbox

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

	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled"`
	Checked  bool        `json:"checked"`
	Color    string      `json:"color"`
	Options  interface{} `json:"options"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "checkboxField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)

	return p
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
func (p *Component) SetValue(value []interface{}) *Component {
	p.Value = value

	return p
}

// 本地渲染数据
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 本地渲染数据
func (p *Component) SetChecked(checked bool) *Component {
	p.Checked = checked

	return p
}

// list 列表模式下 icon 显示的位置
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 设置单选属性，[]map[string]interface{}{{"text": "Title1","value": "value1"},{"text": "Title2","value": "value2"}}
// 或者 map[interface{}]interface{}{"value1":"Title1","value2":"Title2"}
func (p *Component) SetOptions(options interface{}) *Component {
	var data []map[string]interface{}

	if mapOptions, ok := options.(map[interface{}]interface{}); ok {
		for k, v := range mapOptions {
			option := map[string]interface{}{
				"text":  v,
				"value": k,
			}

			data = append(data, option)
		}
	} else if sliceOptions, ok := options.([]map[string]interface{}); ok {
		data = sliceOptions
	}

	p.Options = data

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "checkboxField"

	return p
}
