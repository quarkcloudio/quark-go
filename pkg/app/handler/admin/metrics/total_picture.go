package metrics

import (
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder/metrics"
	"github.com/quarkcms/quark-go/pkg/component/admin/statistic"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type TotalPicture struct {
	metrics.AdminValue
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
