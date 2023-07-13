package menu

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
)

type SubMenu struct {
	component.Element
	Disabled       bool        `json:"disabled"`
	Icon           string      `json:"icon"`
	PopupClassName string      `json:"popupClassName"`
	PopupOffset    interface{} `json:"popupOffset"`
	Title          string      `json:"title"`
	Items          interface{} `json:"items"`
}

// 初始化
func (p *SubMenu) Init() *SubMenu {
	p.Component = "menuSubMenu"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 是否禁用
func (p *SubMenu) SetDisabled(disabled bool) *SubMenu {
	p.Disabled = disabled

	return p
}

// 菜单图标
func (p *SubMenu) SetIcon(icon string) *SubMenu {
	p.Icon = icon

	return p
}

// 子菜单样式，mode="inline" 时无效
func (p *SubMenu) SetPopupClassName(popupClassName string) *SubMenu {
	p.PopupClassName = popupClassName

	return p
}

// 子菜单偏移量，mode="inline" 时无效
func (p *SubMenu) SetPopupOffset(popupOffset interface{}) *SubMenu {
	p.PopupOffset = popupOffset

	return p
}

// 子菜单项值
func (p *SubMenu) SetTitle(title string) *SubMenu {
	p.Title = title

	return p
}

// 菜单项
func (p *SubMenu) SetItems(items interface{}) *SubMenu {
	p.Items = items

	return p
}
