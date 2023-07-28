package actions

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Modal struct {
	Action
	Width          int  `json:"width"`
	DestroyOnClose bool `json:"destroyOnClose"`
}

// 初始化
func (p *Modal) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "modal"
	p.Width = 520

	return p
}

// 宽度
func (p *Modal) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *Modal) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 内容
func (p *Modal) GetBody(ctx *builder.Context) interface{} {
	return nil
}

// 弹窗行为
func (p *Modal) GetActions(ctx *builder.Context) []interface{} {
	return []interface{}{}
}
