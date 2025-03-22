package actions

import "github.com/quarkcloudio/quark-go/v3"

type Drawer struct {
	Action
	Width          int  `json:"width"`          // 抽屉弹出层宽度
	DestroyOnClose bool `json:"destroyOnClose"` // 关闭时销毁弹出层里的子元素
}

// 初始化
func (p *Drawer) New(ctx *quark.Context) interface{} {
	p.ActionType = "drawer"
	p.Width = 520

	return p
}

// 内容
func (p *Drawer) Content(ctx *quark.Context) interface{} {
	return nil
}

// 数据（异步获取）
func (p *Drawer) Data(ctx *quark.Context) map[string]interface{} {
	return map[string]interface{}{}
}

// 宽度
func (p *Drawer) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *Drawer) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 内容
func (p *Drawer) GetBody(ctx *quark.Context) interface{} {
	return p.Content(ctx)
}

// 弹窗行为
func (p *Drawer) GetActions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}
