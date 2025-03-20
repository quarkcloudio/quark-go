package login

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/page"
)

// 后台登录模板
type Template struct {
	page.Template
	IndexPath  string // 渲染登录页面路由
	HandlePath string // 后台登录执行路由
	FromStyle  string
	Api        string
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/miniapp/login/:resource/index"   // 渲染登录页面路由
	p.HandlePath = "/api/miniapp/login/:resource/handle" // 后台登录执行路由
	return p
}

// 初始化路由映射
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render)   // 渲染登录页面路由
	p.POST(p.HandlePath, p.Handle) // 后台登录执行路由

	return p
}

// 初始化模板
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "登录"

	return p
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
	return p
}

// 内容
func (p *Template) Content(ctx *quark.Context) interface{} {

	return "登录页面"
}

// 执行表单
func (p *Template) Handle(ctx *quark.Context) error {
	return ctx.JSONError("请自行处理表单逻辑")
}
