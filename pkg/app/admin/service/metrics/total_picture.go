package metrics

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/statistic"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/dashboard/metrics"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type TotalPicture struct {
	metrics.Value
}

// 初始化
func (p *TotalPicture) Init() *TotalPicture {
	p.Title = "图片数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalPicture) Calculate() *statistic.Component {

	return p.
		Init().
		Count(db.Client.Model(&model.Picture{})).
		SetValueStyle(map[string]string{"color": "#cf1322"})
}
