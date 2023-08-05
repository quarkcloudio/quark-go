package actions

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type DrawerForm struct {
	Action
	Width          int  `json:"width"`
	DestroyOnClose bool `json:"destroyOnClose"`
}

// 初始化
func (p *DrawerForm) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "drawerForm"
	p.Width = 520

	return p
}

// 宽度
func (p *DrawerForm) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *DrawerForm) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 内容
func (p *DrawerForm) GetBody(ctx *builder.Context) interface{} {
	return nil
}

// 弹窗行为
func (p *DrawerForm) GetActions(ctx *builder.Context) []interface{} {
	return []interface{}{}
}
