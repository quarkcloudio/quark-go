package types

import "github.com/quarkcloudio/quark-go/v4"

type ModalFormer interface {
	Actioner

	// 字段
	Fields(ctx *quark.Context) []interface{}

	// 表单数据（异步获取）
	Data(ctx *quark.Context) map[string]interface{}

	// 宽度
	GetWidth() int

	// 关闭时销毁 Modal 里的子元素
	GetDestroyOnClose() bool

	// 获取取消按钮文案
	GetCancelText() string

	// 获取提交按钮文案
	GetSubmitText() string

	// 表单提交接口的类型，GET 或 POST，默认 POST
	GetApiType() string

	// 提交表单的数据是否打开新页面，只有在GET类型的时候有效
	GetTargetBlank() bool
}
