package metrics

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/statistic"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/dashboard/metrics"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type TotalAdmin struct {
	metrics.Value
}

// 初始化
func (p *TotalAdmin) Init() *TotalAdmin {
	p.Title = "管理员数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalAdmin) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.Admin{})).
		SetValueStyle(map[string]string{"color": "#3f8600"})
}
