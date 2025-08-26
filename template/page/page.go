package page

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/action"
	"github.com/quarkcloudio/quark-go/v4/component/card"
	"github.com/quarkcloudio/quark-go/v4/component/pagecontainer"
)

// 增删改查模板
type Template struct {
	quark.Template
	IndexPath string // 列表路径
	Title     string // 页面标题
	SubTitle  string // 页面子标题
	BackIcon  bool   // 页面是否携带返回Icon
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/admin/page/:resource/index" // 列表路径
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.IndexRender) // 列表
	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 页面是否携带返回Icon
	p.BackIcon = true

	return p
}

// 模版初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
	return p
}

// 自定义路由
func (p *Template) Route() interface{} {
	// p.GET("/api/admin/:resource/demo", p.Demo)
	return p
}

// 获取标题
func (p *Template) GetTitle() string {
	return p.Title
}

// 获取子标题
func (p *Template) GetSubTitle() string {
	return p.SubTitle
}

// 页面是否携带返回Icon
func (p *Template) GetBackIcon() bool {
	return p.BackIcon
}

// 行为
func (p *Template) ExtraActions(ctx *quark.Context) []interface{} {
	action := action.New()
	return []interface{}{
		action.
			SetLabel("返回上一页").
			SetType("link", false).
			SetActionType("back"),
	}
}

// 列表页渲染
func (p *Template) IndexRender(ctx *quark.Context) error {
	template := ctx.Template.(Pager)

	// 组件渲染
	body := template.ContentComponentRender(ctx)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 自定义内容页内容
func (p *Template) Content(ctx *quark.Context) interface{} {
	return "please implement the component content."
}

// 内容页标题
func (p *Template) ContentTitle(ctx *quark.Context) string {
	template := ctx.Template.(Pager)
	title := template.GetTitle()

	return title
}

// 渲染自定义内容页组件
func (p *Template) ContentComponentRender(ctx *quark.Context) interface{} {
	template := ctx.Template.(Pager)

	// 内容页标题
	title := template.ContentTitle(ctx)

	// 内容页右上角自定义区域行为
	extraActions := template.ExtraActions(ctx)

	// 包裹在组件内的内容页字段
	body := template.Content(ctx)

	return p.ContentWithinCard(
		ctx,
		title,
		extraActions,
		body,
	)
}

// 在卡片内的内容页组件
func (p *Template) ContentWithinCard(
	ctx *quark.Context,
	title string,
	extra interface{},
	content interface{}) interface{} {

	return (&card.Component{}).
		Init().
		SetTitle(title).
		SetHeaderBordered(true).
		SetExtra(extra).
		SetBody(content)
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(Pager)

	// 页面容器组件渲染
	return template.PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(Pager)

	// 页面标题
	title := template.GetTitle()

	// 页面子标题
	subTitle := template.GetSubTitle()

	// 页面是否携带返回Icon
	backIcon := template.GetBackIcon()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	if !backIcon {
		header.SetBackIcon(false)
	}

	return (&pagecontainer.Component{}).
		Init().
		SetHeader(header).
		SetBody(body)
}
