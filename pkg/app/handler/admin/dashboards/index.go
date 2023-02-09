package dashboards

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/metrics"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/admindashboard"
)

type Index struct {
	admindashboard.Template
}

// 初始化
func (p *Index) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Index) Cards(ctx *builder.Context) interface{} {

	return []any{
		&metrics.TotalAdmin{},
		&metrics.TotalLog{},
		&metrics.TotalPicture{},
		&metrics.TotalFile{},
		&metrics.SystemInfo{},
		&metrics.TeamInfo{},
	}
}
