package resources

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type Position struct {
	resource.Template
}

// 初始化
func (p *Position) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "职位"

	// 模型
	p.Model = &model.Position{}

	// 分页
	p.PageSize = 10

	// 默认排序
	p.IndexQueryOrder = "sort asc, id asc"

	return p
}

// 字段
func (p *Position) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称").
			SetRules([]rule.Rule{
				rule.Required("名称必须填写"),
				rule.Min(2, "名称不能少于2个字符"),
				rule.Max(100, "名称不能超过100个字符"),
			}),
		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),
		field.TextArea("remark", "备注"),
		field.Switch("status", "状态").
			SetRules([]rule.Rule{
				rule.Required("请选择状态"),
			}).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Position) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.Status(),
	}
}

// 行为
func (p *Position) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.CreateModal(),
		actions.ChangeStatus(),
		actions.EditModal(),
		actions.Delete(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
	}
}
