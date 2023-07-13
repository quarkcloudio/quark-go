package searchbar

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Value        interface{} `json:"value"`
	Placeholder  string      `json:"placeholdeV"`
	Radius       int         `json:"radius"`
	ClearButton  string      `json:"clearButton"`
	CancelButton string      `json:"cancelButton"`
	CancelText   string      `json:"cancelText"`
	BgColor      string      `json:"bgColor"`
	Maxlength    int         `json:"maxlength"`
	Focus        bool        `json:"focus"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "searchbar"
	p.Placeholder = "搜索"
	p.Radius = 100
	p.ClearButton = "auto"
	p.CancelButton = "auto"
	p.CancelText = "取消"
	p.BgColor = "#FFFFFF"
	p.Maxlength = 100
	p.SetKey("searchbar", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 搜索栏绑定值
func (p *Component) SetValue(value interface{}) *Component {
	p.Value = value

	return p
}

// 搜索栏Placeholder
func (p *Component) SetPlaceholder(placeholder string) *Component {
	p.Placeholder = placeholder
	return p
}

// 搜索栏圆角，单位px
func (p *Component) SetRadius(radius int) *Component {
	p.Radius = radius

	return p
}

// 是否显示清除按钮，可选值always-一直显示、auto-输入框不为空时显示、none-一直不显示
func (p *Component) SetClearButton(clearButton string) *Component {
	p.ClearButton = clearButton

	return p
}

// 是否显示取消按钮，可选值always-一直显示、auto-输入框不为空时显示、none-一直不显示
func (p *Component) SetCancelButton(cancelButton string) *Component {
	p.CancelButton = cancelButton

	return p
}

// 取消按钮的文字
func (p *Component) SetCancelText(cancelText string) *Component {
	p.CancelText = cancelText

	return p
}

// 输入框背景颜色
func (p *Component) SetBgColor(bgColor string) *Component {
	p.BgColor = bgColor

	return p
}

// 输入最大长度
func (p *Component) SetMaxlength(maxlength int) *Component {
	p.Maxlength = maxlength

	return p
}

// 是否包含状态栏
func (p *Component) SetFocus(focus bool) *Component {
	p.Focus = focus

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "searchbar"

	return p
}
