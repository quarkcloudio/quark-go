package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type ModalForm struct {
	Action
	Width          int    `json:"width"`
	DestroyOnClose bool   `json:"destroyOnClose"`
	CancelText     string `json:"cancelText"`
	SubmitText     string `json:"submitText"`
}

// 初始化
func (p *ModalForm) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "modalForm"
	p.Width = 520
	p.Reload = "table"
	p.CancelText = "取消"
	p.SubmitText = "提交"

	return p
}

// 表单字段
func (p *ModalForm) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 表单数据（异步获取）
func (p *ModalForm) Data(ctx *builder.Context) map[string]interface{} {
	return map[string]interface{}{}
}

// 宽度
func (p *ModalForm) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *ModalForm) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 获取取消按钮文案
func (p *ModalForm) GetCancelText() string {
	return p.CancelText
}

// 获取提交按钮文案
func (p *ModalForm) GetSubmitText() string {
	return p.SubmitText
}
