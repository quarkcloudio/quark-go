package pages

import (
	"github.com/quarkcms/quark-go/pkg/builder/template/adminpage"
)

type Component struct {
	adminpage.Template
}

// 初始化
func (p *Component) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	return p
}
