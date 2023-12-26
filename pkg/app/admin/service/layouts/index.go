package layouts

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/layout"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Index struct {
	layout.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {
	return p
}
