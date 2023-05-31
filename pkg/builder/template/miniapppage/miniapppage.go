package miniapppage

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/navbar"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/page"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	template.Template
	Title string
	Style string
}

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 注册路由映射
	p.GET("/api/miniapp/page/:resource/index", p.Render) // 渲染页面路由

	// 标题
	p.Title = "QuarkGo"

	return p
}

// 头部导航
func (p *Template) Navbar(ctx *builder.Context, navbar *navbar.Component) interface{} {
	return nil
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {
	return nil
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	var (
		components []interface{}
	)

	// 标题
	title := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Title").
		String()

	// 导航
	navbar := ctx.Template.(interface {
		Navbar(ctx *builder.Context, navbar *navbar.Component) interface{}
	}).Navbar(ctx, navbar.New())

	// 样式
	style := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Style").
		String()

	// 内容
	content := ctx.Template.(interface {
		Content(ctx *builder.Context) interface{}
	}).Content(ctx)
	components = append(components, content)

	// 组件
	component := (&page.Component{}).
		Init().
		SetTitle(title).
		SetNavbar(navbar).
		SetStyle(style).
		SetContent(components).
		JsonSerialize()

	return ctx.JSON(200, component)
}
