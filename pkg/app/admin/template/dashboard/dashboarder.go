package dashboard

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Dashboarder interface {

	// 模版接口
	builder.Templater

	// 获取页面标题
	GetTitle() string

	// 获取页面子标题
	GetSubTitle() string

	// 内容
	Cards(ctx *builder.Context) []interface{}

	// 页面组件渲染
	PageComponentRender(ctx *builder.Context, body interface{}) interface{}

	// 页面容器组件渲染
	PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{}

	// 组件渲染
	Render(ctx *builder.Context) error
}
