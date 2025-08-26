package types

import "github.com/quarkcloudio/quark-go/v4"

type Drawer interface {
	Actioner

	// 宽度
	GetWidth() int

	// 关闭时销毁 Drawer 里的子元素
	GetDestroyOnClose() bool

	// 内容
	GetBody(ctx *quark.Context) interface{}

	// 弹窗行为
	GetActions(ctx *quark.Context) []interface{}
}
