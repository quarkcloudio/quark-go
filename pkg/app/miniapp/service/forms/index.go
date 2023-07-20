package forms

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/cascader"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/form/fields/checkbox"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/template/form"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Index struct {
	form.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {

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
		p.Field().Checkbox("checkbox", "多选").
			SetOptions([]*checkbox.Option{
				{
					Value: "1",
					Label: "测试1",
				},
				{
					Value: "2",
					Label: "测试2",
				},
			}),
	}
}
