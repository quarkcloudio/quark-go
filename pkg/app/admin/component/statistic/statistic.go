package statistic

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	DecimalSeparator string            `json:"decimalSeparator"`
	GroupSeparator   string            `json:"groupSeparator"`
	Precision        int               `json:"precision"`
	Prefix           string            `json:"prefix"`
	Suffix           string            `json:"suffix"`
	Title            string            `json:"title"`
	Value            int64             `json:"value"`
	ValueStyle       map[string]string `json:"valueStyle"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "statistic"
	p.DecimalSeparator = "."
	p.GroupSeparator = ","

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置小数点
func (p *Component) SetDecimalSeparator(decimalSeparator string) *Component {
	p.DecimalSeparator = decimalSeparator
	return p
}

// 设置千分位标识符
func (p *Component) SetGroupSeparator(groupSeparator string) *Component {
	p.GroupSeparator = groupSeparator
	return p
}

// 数值精度
func (p *Component) SetPrecision(precision int) *Component {
	p.Precision = precision
	return p
}

// 设置数值的前缀
func (p *Component) SetPrefix(prefix string) *Component {
	p.Prefix = prefix
	return p
}

// 设置数值的后缀
func (p *Component) SetSuffix(suffix string) *Component {
	p.Suffix = suffix
	return p
}

// 设置标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 数值内容
func (p *Component) SetValue(value int64) *Component {
	p.Value = value
	return p
}

// 设置数值的样式
func (p *Component) SetValueStyle(valueStyle map[string]string) *Component {
	p.ValueStyle = valueStyle
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "statistic"
	p.DecimalSeparator = "."
	p.GroupSeparator = ","

	return p
}
