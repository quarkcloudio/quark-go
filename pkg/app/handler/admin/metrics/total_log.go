package metrics

import (
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder/metrics"
	"github.com/quarkcms/quark-go/pkg/component/admin/statistic"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type TotalLog struct {
	metrics.AdminValue
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
