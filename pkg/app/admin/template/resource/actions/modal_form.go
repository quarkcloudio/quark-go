package actions

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type ModalForm struct {
	Action
	Width          int  `json:"width"`
	DestroyOnClose bool `json:"destroyOnClose"`
}

// 初始化
func (p *ModalForm) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "modal"
	p.Width = 520

	return p
}

// 宽度
func (p *ModalForm) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *ModalForm) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 字段
func (p *ModalForm) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 内容
func (p *ModalForm) GetBody(ctx *builder.Context) interface{} {
	return nil
}

// 弹窗行为
func (p *ModalForm) GetActions(ctx *builder.Context) []interface{} {
	return []interface{}{}
}
