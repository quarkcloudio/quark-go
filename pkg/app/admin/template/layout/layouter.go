package layout

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Layouter interface {

	// 模版接口
	builder.Templater

	// 组件渲染
	Render(ctx *builder.Context) error
}
