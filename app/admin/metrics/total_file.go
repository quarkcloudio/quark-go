package metrics

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/statistic"
	"github.com/quarkcloudio/quark-go/v3/template/admin/dashboard/metrics"
)

type TotalFile struct {
	metrics.Value
}

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
		Count(db.Client.Model(&model.Attachment{}).Where("type = ?", "FILE")).
		SetValueStyle(map[string]string{"color": "#cf1322"})
}
