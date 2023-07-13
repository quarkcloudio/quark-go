package metrics

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/statistic"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/dashboard/metrics"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type TotalLog struct {
	metrics.Value
}

// 初始化
func (p *TotalLog) Init() *TotalLog {
	p.Title = "日志数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalLog) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.ActionLog{})).
		SetValueStyle(map[string]string{"color": "#999999"})
}
