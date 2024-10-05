package metrics

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/descriptions"
)

type Descriptions struct {
	Metrics
}

// 包含组件的结果
func (p *Descriptions) Result(value interface{}) *descriptions.Component {
	return (&descriptions.Component{}).Init().SetTitle(p.Title).SetItems(value)
}
