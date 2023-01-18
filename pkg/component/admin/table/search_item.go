package table

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type SearchItem struct {
	component.Element
	Value        interface{} `json:"value"`
	DefaultValue interface{} `json:"defaultValue"`
	Label        string      `json:"label"`
	Name         string      `json:"name"`
	Options      interface{} `json:"options"`
	Api          string      `json:"api"`
	Rules        interface{} `json:"rules"`
	RuleMessages interface{} `json:"ruleMessages"`
	Operator     string      `json:"operator"`
	Placeholder  interface{} `json:"placeholder"`
	Load         interface{} `json:"load"`
}

// 初始化
func (p *SearchItem) Init() *SearchItem {
	p.Component = "input"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *SearchItem) SetStyle(style map[string]interface{}) *SearchItem {
	p.Style = style

	return p
}

/**
 * label 标签的文本
 *
 * @param string label
 * @return p
 */
func (p *SearchItem) SetLabel(label string) *SearchItem {
	p.Label = label
	p.Placeholder = "请输入" + p.Label

	return p
}

/**
 * 字段名，支持数组
 *
 * @param string name
 * @return p
 */
func (p *SearchItem) SetName(name string) *SearchItem {
	p.Name = name

	return p
}

/**
 * 校验规则，设置字段的校验逻辑
 *
 * @param array|p rules
 * @return p
 */
func (p *SearchItem) SetRules(rules interface{}, messages interface{}) *SearchItem {
	p.Rules = rules
	p.RuleMessages = messages

	return p
}

/**
 * 设置保存值。
 *
 * @param array|string
 * @return p
 */
func (p *SearchItem) SetValue(value interface{}) *SearchItem {
	p.Value = value

	return p
}

/**
 * 设置默认值。
 *
 * @param array|string
 * @return p
 */
func (p *SearchItem) SetDefault(value interface{}) *SearchItem {
	p.DefaultValue = value

	return p
}

/**
 * 操作符
 *
 * @param string
 * @return p
 */
func (p *SearchItem) SetOperator(operator string) *SearchItem {
	p.Operator = operator

	if p.Operator == "between" {
		p.Placeholder = []string{"开始" + p.Label, "结束" + p.Label}
	}

	return p
}

/**
 * placeholder
 *
 * @param string placeholder
 * @return object
 */
func (p *SearchItem) SetPlaceholder(placeholder interface{}) *SearchItem {

	p.Placeholder = placeholder

	return p
}

/**
 * 控件宽度
 *
 * @param number|string value
 * @return object
 */
func (p *SearchItem) SetWidth(value int) *SearchItem {
	p.Style = map[string]interface{}{
		"width": value,
	}

	return p
}

/**
 * 级联菜单接口
 *
 * @param string api
 * @return object
 */
func (p *SearchItem) SetApi(api string) *SearchItem {
	p.Api = api

	return p
}

/**
 * 单向联动
 *
 * @param  string field
 * @param  string api
 * @return p
 */
func (p *SearchItem) SetLoad(field string, api string) *SearchItem {
	p.Load = map[string]string{
		"field": field,
		"api":   api,
	}

	return p
}

/**
 * 输入框控件
 *
 * @param string options
 * @return object
 */
func (p *SearchItem) Input(options map[interface{}]interface{}) *SearchItem {
	p.Component = "input"

	return p
}

/**
 * 下拉菜单控件
 *
 * @param array options
 * @return object
 */
func (p *SearchItem) Select(options map[interface{}]interface{}) *SearchItem {
	p.Component = "select"

	var data []map[string]interface{}

	for k, v := range options {
		option := map[string]interface{}{
			"label": v,
			"value": k,
		}

		data = append(data, option)
	}

	p.Options = data

	if p.Operator == "between" {
		p.Placeholder = []string{"开始" + p.Label, "结束" + p.Label}
	} else {
		p.Placeholder = "请选择" + p.Label
	}

	return p
}

/**
 * 多选下拉菜单控件
 *
 * @param array options
 * @return object
 */
func (p *SearchItem) MultipleSelect(options map[interface{}]interface{}) *SearchItem {
	p.Component = "multipleSelect"

	var data []map[string]interface{}

	for k, v := range options {
		option := map[string]interface{}{
			"label": v,
			"value": k,
		}

		data = append(data, option)
	}

	p.Options = data
	p.Placeholder = "请选择" + p.Label

	return p
}

/**
 * 时间控件
 *
 * @param string options
 * @return object
 */
func (p *SearchItem) Datetime(options map[interface{}]interface{}) *SearchItem {
	p.Component = "datetime"

	return p
}

/**
 * 日期控件
 *
 * @param string options
 * @return object
 */
func (p *SearchItem) Date(options map[interface{}]interface{}) *SearchItem {
	p.Component = "date"

	return p
}

/**
 * 级联菜单
 *
 * @param array options
 * @return object
 */
func (p *SearchItem) Cascader(options map[interface{}]interface{}) *SearchItem {
	p.Component = "cascader"

	var data []map[string]interface{}

	for k, v := range options {
		option := map[string]interface{}{
			"label": v,
			"value": k,
		}

		data = append(data, option)
	}

	p.Options = data
	p.Placeholder = "请选择" + p.Label

	return p
}

// 组件json序列化
func (p *SearchItem) JsonSerialize() *SearchItem {
	p.Component = "input"

	return p
}
