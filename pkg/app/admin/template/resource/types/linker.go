package types

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Linker interface {
	Actioner

	// 获取跳转链接
	GetHref(ctx *builder.Context) string

	// 相当于 a 链接的 target 属性，href 存在时生效
	GetTarget(ctx *builder.Context) string
}
