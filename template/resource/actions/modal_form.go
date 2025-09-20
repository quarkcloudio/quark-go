package actions

import "github.com/quarkcloudio/quark-go/v4"

type ModalForm struct {
	Action
	Width          int    `json:"width"`          // 弹出层宽度
	DestroyOnClose bool   `json:"destroyOnClose"` // 关闭时销毁弹出层里的子元素
	CancelText     string `json:"cancelText"`     // 获取取消按钮文案
	SubmitText     string `json:"submitText"`     // 获取提交按钮文案
	ApiType        string `json:"apiType"`        // 表单提交接口的类型，GET 或 POST，默认 POST
	TargetBlank    bool   `json:"targetBlank"`    // 提交表单的数据是否打开新页面，只有在GET类型的时候有效
}

// 初始化
func (p *ModalForm) New(ctx *quark.Context) interface{} {
	p.ActionType = "modalForm"
	p.Width = 520
	p.Reload = "table"
	p.CancelText = "取消"
	p.SubmitText = "提交"
	p.ApiType = "POST"
	p.TargetBlank = false
	p.Size = "small"
	p.Type = "default"

	return p
}

// 表单字段
func (p *ModalForm) Fields(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

// 表单数据（异步获取）
func (p *ModalForm) Data(ctx *quark.Context) map[string]interface{} {
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

// 表单提交接口的类型，GET 或 POST，默认 POST
func (p *ModalForm) GetApiType() string {
	return p.ApiType
}

// 提交表单的数据是否打开新页面，只有在GET类型的时候有效
func (p *ModalForm) GetTargetBlank() bool {
	return p.TargetBlank
}
