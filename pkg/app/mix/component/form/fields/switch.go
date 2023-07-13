package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Switch struct {
	Item
	Value    interface{} `json:"value"`
	Checked  bool        `json:"checked"`
	Disabled bool        `json:"disabled"`
	Type     string      `json:"type"`
	Color    string      `json:"color"`
}

// 初始化
func (p *Switch) Init() *Switch {
	p.Component = "switchField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Type = "switch"

	return p
}

// 默认值
func (p *Switch) SetValue(value interface{}) *Switch {
	p.Value = value

	return p
}

// 是否选中
func (p *Switch) SetChecked(checked bool) *Switch {
	p.Checked = checked

	return p
}

// 是否禁用
func (p *Switch) SetDisabled(disabled bool) *Switch {
	p.Disabled = disabled

	return p
}

// 样式，有效值：switch, checkbox
func (p *Switch) SetType(switchType string) *Switch {
	p.Type = switchType

	return p
}

// switch 的颜色，同 css 的 color
func (p *Switch) SetColor(color string) *Switch {
	p.Color = color

	return p
}

// 组件json序列化
func (p *Switch) JsonSerialize() *Switch {
	p.Component = "switchField"

	return p
}
