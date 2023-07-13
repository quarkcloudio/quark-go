package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Item struct {
	component.Element
	Name            interface{} `json:"name"`
	Rules           interface{} `json:"rules"`
	Required        bool        `json:"required"`
	Label           string      `json:"label"`
	LabelWidth      int         `json:"labelWidth"`
	ErrorMessage    string      `json:"errorMessage"`
	LabelAlign      string      `json:"labelAlign"`
	LabelPosition   string      `json:"labelPosition"`
	ValidateTrigger string      `json:"validateTrigger"`
	LeftIcon        string      `json:"leftIcon"`
	IconColor       string      `json:"iconColor"`
	Body            interface{} `json:"body"`
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "formItem"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.LabelWidth = 70
	p.LabelAlign = "left"
	p.LabelPosition = "left"
	p.ValidateTrigger = "submit"
	p.IconColor = "#606266"

	return p
}

// Set style.
func (p *Item) SetStyle(style interface{}) *Item {
	p.Style = style

	return p
}

// 表单域的属性名，在使用校验规则时必填
func (p *Item) SetName(name interface{}) *Item {
	p.Name = name

	return p
}

// 表单校验规则
func (p *Item) SetRules(rules interface{}) *Item {
	p.Rules = rules

	return p
}

// label 右边显示红色"*"号，样式显示不会对校验规则产生效果
func (p *Item) SetRequired(required bool) *Item {
	p.Required = required

	return p
}

// 输入框左边的文字提示
func (p *Item) SetLabel(label string) *Item {
	p.Label = label

	return p
}

// label的宽度，单位px
func (p *Item) SetLabelWidth(labelWidth int) *Item {
	p.LabelWidth = labelWidth

	return p
}

// 显示的错误提示内容，如果为空字符串或者false，则不显示错误信息
func (p *Item) SetErrorMessage(errorMessage string) *Item {
	p.ErrorMessage = errorMessage

	return p
}

// label的文字对齐方式
func (p *Item) SetLabelAlign(labelAlign string) *Item {
	p.LabelAlign = labelAlign

	return p
}

// label的文字的位置
func (p *Item) SetLabelPosition(labelPosition string) *Item {
	p.LabelPosition = labelPosition

	return p
}

// 表单校验时机
func (p *Item) SetValidateTrigger(validateTrigger string) *Item {
	p.ValidateTrigger = validateTrigger

	return p
}

// label左边的图标，限uni-ui的图标名称
func (p *Item) SetLeftIcon(leftIcon string) *Item {
	p.LeftIcon = leftIcon

	return p
}

// 左边通过icon配置的图标的颜色
func (p *Item) SetIconColor(iconColor string) *Item {
	p.IconColor = iconColor

	return p
}

// 内容
func (p *Item) SetBody(body interface{}) *Item {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Item) JsonSerialize() *Item {
	p.Component = "formItem"

	return p
}
