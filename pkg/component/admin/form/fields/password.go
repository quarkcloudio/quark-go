package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Password struct {
	Item
	AddonAfter       string `json:"addonAfter"`
	AddonBefore      string `json:"addonBefore"`
	MaxLength        int    `json:"maxLength"`
	Prefix           string `json:"prefix"`
	Size             string `json:"size"`
	Suffix           string `json:"suffix"`
	AllowClear       bool   `json:"allowClear"`
	VisibilityToggle bool   `json:"visibilityToggle"`
}

// 初始化
func (p *Password) Init() *Password {
	p.Component = "passwordField"
	p.MaxLength = 200
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)
	p.VisibilityToggle = true

	return p
}

// 带标签的 input，设置后置标签。例如：'http://'
func (p *Password) SetAddonAfter(addonAfter string) *Password {
	p.AddonAfter = addonAfter

	return p
}

// 带标签的 input，设置前置标签。例如：'.com'
func (p *Password) SetAddonBefore(addonBefore string) *Password {
	p.AddonBefore = addonBefore

	return p
}

// 最大长度
func (p *Password) SetMaxLength(maxLength int) *Password {
	p.MaxLength = maxLength

	return p
}

// 带有前缀图标的 input
func (p *Password) SetPrefix(prefix string) *Password {
	p.Prefix = prefix

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Password) SetSize(size string) *Password {
	p.Size = size

	return p
}

// 带有后缀图标的 input
func (p *Password) SetSuffix(suffix string) *Password {
	p.Suffix = suffix

	return p
}

// 可以点击清除图标删除内容
func (p *Password) SetAllowClear(allowClear bool) *Password {
	p.AllowClear = allowClear

	return p
}

// 组件json序列化
func (p *Password) JsonSerialize() *Password {
	p.Component = "passwordField"

	return p
}
