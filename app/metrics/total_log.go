package metrics

import (
	"github.com/quarkcloudio/quark-go/v4/component/statistic"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/template/dashboard/metrics"
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
