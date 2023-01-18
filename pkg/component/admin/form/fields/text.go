package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Text struct {
	Item
	AddonAfter  string `json:"addonAfter"`
	AddonBefore string `json:"addonBefore"`
	MaxLength   int    `json:"maxLength"`
	Prefix      string `json:"prefix"`
	Size        string `json:"size"`
	Suffix      string `json:"suffix"`
	AllowClear  bool   `json:"allowClear"`
}

// 初始化
func (p *Text) Init() *Text {
	p.Component = "textField"
	p.MaxLength = 200
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)

	return p
}

// 带标签的 input，设置后置标签。例如：'http://'
func (p *Text) SetAddonAfter(addonAfter string) *Text {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签。例如：'.com'
func (p *Text) SetAddonBefore(addonBefore string) *Text {
	p.AddonBefore = addonBefore

	return p
}

// 最大长度
func (p *Text) SetMaxLength(maxLength int) *Text {
	p.MaxLength = maxLength

	return p
}

// 带有前缀图标的 input
func (p *Text) SetPrefix(prefix string) *Text {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Text) SetSize(size string) *Text {
	p.Size = size

	return p
}

// 带有后缀图标的 input
func (p *Text) SetSuffix(suffix string) *Text {
	p.Suffix = suffix

	return p
}

// 可以点击清除图标删除内容
func (p *Text) SetAllowClear(allowClear bool) *Text {
	p.AllowClear = allowClear

	return p
}

// 组件json序列化
func (p *Text) JsonSerialize() *Text {
	p.Component = "textField"

	return p
}
