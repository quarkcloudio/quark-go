package menu

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
)

type ItemGroup struct {
	component.Element
	Title string      `json:"title"`
	Items interface{} `json:"items"`
}

// 初始化
func (p *ItemGroup) Init() *ItemGroup {
	p.Component = "menuItemGroup"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置收缩时展示的悬浮标题
func (p *ItemGroup) SetTitle(title string) *ItemGroup {
	p.Title = title

	return p
}

// 设置按钮文字
func (p *ItemGroup) SetItems(items interface{}) *ItemGroup {
	p.Items = items

	return p
}
