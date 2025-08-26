package metrics

import (
	"github.com/quarkcloudio/quark-go/v4/component/statistic"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/template/dashboard/metrics"
)

type TotalAdmin struct {
	metrics.Value
}

// 初始化
func (p *TotalAdmin) Init() *TotalAdmin {
	p.Title = "用户数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalAdmin) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.User{})).
		SetValueStyle(map[string]string{"color": "#3f8600"})
}
