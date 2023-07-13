package form

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"
)

type Component struct {
	component.Element
	Api             string        `json:"api"`
	ApiType         string        `json:"apiType"`
	Model           interface{}   `json:"model"`
	Rules           interface{}   `json:"rules"`
	ValidateTrigger string        `json:"validateTrigger"`
	LabelPosition   string        `json:"labelPosition"`
	LabelWidth      int           `json:"labelWidth"`
	LabelAlign      string        `json:"labelAlign"`
	ErrShowType     string        `json:"errShowType"`
	Border          bool          `json:"border"`
	Actions         []interface{} `json:"actions"`
	Body            interface{}   `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "form"
	p.SetKey("form", component.DEFAULT_CRYPT)
	p.ApiType = "POST"
	p.ValidateTrigger = "submit"
	p.LabelPosition = "left"
	p.LabelWidth = 70
	p.LabelAlign = "left"
	p.ErrShowType = "undertext"

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 表单接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api

	return p
}

// 表单接口提交方式
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

// 表单数据
func (p *Component) SetModel(model interface{}) *Component {
	p.Model = model

	return p
}

// 表单校验规则
func (p *Component) SetRules(rules interface{}) *Component {
	p.Rules = rules
	return p
}

// 表单校验时机,blur仅在 uni-easyinput 中生效
func (p *Component) SetValidateTrigger(validateTrigger string) *Component {
	p.ValidateTrigger = validateTrigger

	return p
}

// label 位置
func (p *Component) SetLabelPosition(labelPosition string) *Component {
	p.LabelPosition = labelPosition

	return p
}

// label 宽度，单位 px
func (p *Component) SetLabelWidth(labelWidth int) *Component {
	p.LabelWidth = labelWidth

	return p
}

// label 居中方式
func (p *Component) SetLabelAlign(labelAlign string) *Component {
	p.LabelAlign = labelAlign

	return p
}

// 表单错误信息提示方式
func (p *Component) SetErrShowType(errShowType string) *Component {
	p.ErrShowType = errShowType

	return p
}

// 是否显示分格线
func (p *Component) SeBorder(border bool) *Component {
	p.Border = border

	return p
}

// 设置表单行为
func (p *Component) SetActions(actions []interface{}) *Component {
	p.Actions = actions

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "form"

	return p
}
