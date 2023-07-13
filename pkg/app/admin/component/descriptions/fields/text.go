package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Text struct {
	component.Element
	Label     string      `json:"label"`
	Tooltip   string      `json:"tooltip"`
	Ellipsis  bool        `json:"ellipsis"`
	Copyable  bool        `json:"copyable"`
	Span      int         `json:"span"`
	ValueType string      `json:"valueType"`
	ValueEnum string      `json:"valueEnum"`
	DataIndex string      `json:"dataIndex"`
	Value     interface{} `json:"value"`
}

// 初始化组件
func New() *Text {
	return (&Text{}).Init()
}

// 初始化
func (p *Text) Init() *Text {
	p.Component = "text"
	p.ValueType = "text"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Text) SetStyle(style map[string]interface{}) *Text {
	p.Style = style

	return p
}

// 内容的描述
func (p *Text) SetLabel(label string) *Text {
	p.Label = label

	return p
}

// 内容的补充描述，hover 后显示
func (p *Text) SetTooltip(tooltip string) *Text {
	p.Tooltip = tooltip

	return p
}

// 是否自动缩略
func (p *Text) SetEllipsis(ellipsis bool) *Text {
	p.Ellipsis = ellipsis

	return p
}

// 是否支持复制
func (p *Text) SetCopyable(copyable bool) *Text {
	p.Copyable = copyable

	return p
}

// 列数
func (p *Text) SetSpan(span int) *Text {
	p.Span = span

	return p
}

// 值类型
func (p *Text) SetValueType(valueType string) *Text {
	p.ValueType = valueType

	return p
}

// 值枚举
func (p *Text) SetValueEnum(valueEnum string) *Text {
	p.ValueEnum = valueEnum

	return p
}

// 索引
func (p *Text) SetDataIndex(dataIndex string) *Text {
	p.DataIndex = dataIndex

	return p
}

// 设置保存值。
func (p *Text) SetValue(value interface{}) *Text {
	p.Value = value

	return p
}
