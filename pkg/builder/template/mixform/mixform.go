package mixform

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/mixpage"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	mixpage.Template
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
	p.GET("/api/mix/page/:resource/index", "Render") // 渲染页面路由

	// 标题
	p.Title = "QuarkGo"

	return p
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {
	return nil
}
