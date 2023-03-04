package pages

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/mixpage"
)

type Index struct {
	mixpage.Template
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {

	return "Hello World"
}
