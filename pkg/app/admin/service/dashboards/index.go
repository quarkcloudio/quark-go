package dashboards

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/metrics"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/dashboard"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Index struct {
	dashboard.Template
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
