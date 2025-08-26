package dashboards

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/metrics"
	"github.com/quarkcloudio/quark-go/v4/template/dashboard"
)

type Index struct {
	dashboard.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {
	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Index) Cards(ctx *quark.Context) []interface{} {
	return []interface{}{
		&metrics.TotalAdmin{},
		&metrics.TotalLog{},
		&metrics.TotalImage{},
		&metrics.TotalFile{},
		&metrics.SystemInfo{},
		&metrics.TeamInfo{},
	}
}
