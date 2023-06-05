package form

import (
	"github.com/quarkcms/quark-go/pkg/component/miniapp/component"
)

type Component struct {
	component.Element
	Api        string                 `json:"api"`
	ApiType    string                 `json:"apiType"`
	ModelValue map[string]interface{} `json:"modelValue"`
	Rules      interface{}            `json:"rules"`
	Actions    []interface{}          `json:"actions"`
	Body       interface{}            `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "form"
	p.SetKey("form", component.DEFAULT_CRYPT)
	p.ApiType = "POST"

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
func (p *Component) SetModelValue(modelValue map[string]interface{}) *Component {
	p.ModelValue = modelValue

	return p
}

// 表单校验规则
func (p *Component) SetRules(rules interface{}) *Component {
	p.Rules = rules
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
