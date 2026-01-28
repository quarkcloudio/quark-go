package dashboard

import (
	"github.com/quarkcloudio/quark-go/v4"
)

type Dashboarder interface {

	// 模版接口
	quark.Templater

	// 内容
	Cards(ctx *quark.Context) []interface{}

	// 组件渲染
	Render(ctx *quark.Context) error
}
