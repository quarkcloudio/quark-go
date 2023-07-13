package layouts

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/template/layout"

type Index struct {
	layout.Template
}

// 初始化
func (p *Index) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	return p
}
