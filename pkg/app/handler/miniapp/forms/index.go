package forms

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/miniappform"
	"github.com/quarkcms/quark-go/pkg/component/miniapp/form/fields/cascader"
)

type Index struct {
	miniappform.Template
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 字段
func (p *Index) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{
		p.Field().Input("username", "姓名"),
		p.Field().Calendar("date", "日历"),
		p.Field().Cascader("areas", "地域").
			SetOptions([]*cascader.Option{
				{
					Value: "1",
					Text:  "测试1",
				},
				{
					Value: "2",
					Text:  "测试2",
				},
			}),
	}
}
