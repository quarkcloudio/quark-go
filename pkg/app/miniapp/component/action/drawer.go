package action

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Drawer struct {
	component.Element
	Mask      bool        `json:"mask"`
	MaskClick bool        `json:"maskClick"`
	Mode      string      `json:"mode"`
	Width     int         `json:"width"`
	Body      interface{} `json:"body"`
}

// 初始化
func (p *Drawer) Init() *Drawer {
	p.Component = "drawer"
	p.SetKey("drawer", component.DEFAULT_CRYPT)
	p.Mask = true
	p.MaskClick = true
	p.Mode = "left"
	p.Width = 220

	return p
}

// Set style.
func (p *Drawer) SetStyle(style map[string]interface{}) *Drawer {
	p.Style = style

	return p
}

// 是否显示遮罩
func (p *Drawer) SetMask(mask bool) *Drawer {
	p.Mask = mask

	return p
}

// 点击遮罩是否可以关闭抽屉
func (p *Drawer) SetMaskClick(maskClick bool) *Drawer {
	p.MaskClick = maskClick

	return p
}

// Drawer滑出位置，可选值：left（从左侧滑出）， right（从右侧滑出）
func (p *Drawer) SetMode(mode string) *Drawer {
	p.Mode = mode

	return p
}

// Drawer 宽度，仅vue页面设置生效
func (p *Drawer) SetWidth(width int) *Drawer {
	p.Width = width

	return p
}

// 容器控件里面的内容
func (p *Drawer) SetBody(body interface{}) *Drawer {
	p.Body = body

	return p
}
