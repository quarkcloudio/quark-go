package layouts

import "github.com/quarkcms/quark-go/pkg/builder/template/adminlayout"

type Index struct {
	adminlayout.Template
}

// 初始化
func (p *Index) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	return p
}
