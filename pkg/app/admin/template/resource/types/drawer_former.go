package types

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type DrawerFormer interface {
	Actioner

	// 字段
	Fields(ctx *builder.Context) []interface{}

	// 表单数据（异步获取）
	Data(ctx *builder.Context) map[string]interface{}

	// 宽度
	GetWidth() int

	// 关闭时销毁 Modal 里的子元素
	GetDestroyOnClose() bool

	// 获取取消按钮文案
	GetCancelText() string

	// 获取提交按钮文案
	GetSubmitText() string
}
