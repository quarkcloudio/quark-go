package page

import (
	"github.com/quarkcloudio/quark-go/v3"
)

type Pager interface {

	// 模版接口
	quark.Templater

	// 获取标题
	GetTitle() string

	// 获取子标题
	GetSubTitle() string

	// 页面是否携带返回Icon
	GetBackIcon() bool

	//页面扩展区域行为
	ExtraActions(ctx *quark.Context) []interface{}

	// 页面渲染
	IndexRender(ctx *quark.Context) error

	// 页面组件渲染
	PageComponentRender(ctx *quark.Context, body interface{}) interface{}

	// 页面容器组件渲染
	PageContainerComponentRender(ctx *quark.Context, body interface{}) interface{}

	// 自定义内容页标题
	ContentTitle(ctx *quark.Context) string

	// 自定义内容页内容
	Content(ctx *quark.Context) interface{}

	// 渲染自定义内容页组件
	ContentComponentRender(ctx *quark.Context) interface{}
}
