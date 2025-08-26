package types

import "github.com/quarkcloudio/quark-go/v4"

type Linker interface {
	Actioner

	// 获取跳转链接
	GetHref(ctx *quark.Context) string

	// 相当于 a 链接的 target 属性，href 存在时生效
	GetTarget(ctx *quark.Context) string
}
