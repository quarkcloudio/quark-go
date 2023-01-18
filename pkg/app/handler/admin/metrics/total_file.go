package metrics

import (
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder/metrics"
	"github.com/quarkcms/quark-go/pkg/component/admin/statistic"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type TotalFile struct {
	metrics.AdminValue
}

// 初始化

// 初始化
func (p *TotalFile) Init() *TotalFile {
	p.Title = "文件数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalFile) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.File{})).
		SetValueStyle(map[string]string{"color": "#cf1322"})
}
