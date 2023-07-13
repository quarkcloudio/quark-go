package metrics

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/statistic"
	"gorm.io/gorm"
)

type Value struct {
	Metrics
	Precision int
}

// 记录条数
func (p *Value) Count(DB *gorm.DB) *statistic.Component {
	var count int64
	DB.Count(&count)

	return p.Result(count)
}

// 包含组件的结果
func (p *Value) Result(value int64) *statistic.Component {
	return (&statistic.Component{}).Init().SetTitle(p.Title).SetValue(value)
}
