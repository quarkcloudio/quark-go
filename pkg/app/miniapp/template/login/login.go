package login

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/template/page"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	page.Template
	FromStyle string
	Api       string
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "登录"

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET("/api/miniapp/login/:resource/index", p.Render)   // 渲染登录页面路由
	p.POST("/api/miniapp/login/:resource/handle", p.Handle) // 后台登录执行路由

	return p
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {

	return "登录页面"
}

// 执行表单
func (p *Template) Handle(ctx *builder.Context) error {
	return ctx.JSONError("请自行处理表单逻辑")
}
