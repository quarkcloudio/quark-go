package frontpages

import (
	"github.com/quarkcms/quark-go/pkg/builder/template/adminfrontpage"
)

type Component struct {
	adminfrontpage.Template
}

// 初始化
func (p *Component) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	return p
}
