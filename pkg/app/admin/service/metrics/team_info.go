package metrics

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/descriptions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/dashboard/metrics"
)

type TeamInfo struct {
	metrics.Descriptions
}

// 初始化
func (p *TeamInfo) Init() *TeamInfo {
	p.Title = "团队信息"
	p.Col = 12

	return p
}

// 计算数值
func (p *TeamInfo) Calculate() *descriptions.Component {

	field := &descriptions.Field{}

	return p.Init().Result([]interface{}{
		field.Text("作者").SetValue("tangtanglove"),
		field.Text("联系方式").SetValue("dai_hang_love@126.com"),
		field.Text("官方网址").SetValue("<a href='https://www.quarkcms.com' target='_blank'>www.quarkcms.com</a>"),
		field.Text("文档地址").SetValue("<a href='https://www.quarkcms.com' target='_blank'>查看文档</a>"),
		field.Text("BUG反馈").SetValue("<a href='https://github.com/quarkcms/quark-go/v2/issues' target='_blank'>提交BUG</a>"),
		field.Text("代码仓储").SetValue("<a href='https://github.com/quarkcms/quark-go' target='_blank'>Github</a>"),
	})
}
