package resources

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/table"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/lister"
)

type Department struct {
	resource.Template
}

// 初始化
func (p *Department) Init(ctx *builder.Context) interface{} {

	p.Table.
		SetExpandable(&table.Expandable{
			DefaultExpandedRowKeys: []interface{}{1},
		})

	// 标题
	p.Title = "部门"

	// 模型
	p.Model = &model.Department{}

	// 分页
	p.PerPage = false

	// 默认排序
	p.IndexQueryOrder = "sort asc, id asc"

	return p
}

// 字段
func (p *Department) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	// 列表
	departments, _ := (&model.Department{}).TreeSelect()

	return []interface{}{
		field.Hidden("id", "ID"),                 // 列表读取且不展示的字段
		field.Hidden("pid", "PID").OnlyOnIndex(), // 列表读取且不展示的字段
		field.Text("name", "名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
				rule.Min(2, "名称不能少于2个字符"),
				rule.Max(100, "名称不能超过100个字符"),
			}),
		field.Dependency().SetWhen("id", ">", 1, func() interface{} {
			return field.TreeSelect("pid", "父节点").
				SetData(departments).
				SetRules([]*rule.Rule{
					rule.Required(true, "请选择父节点"),
				}).
				SetDefault(1).
				OnlyOnForms()
		}),
		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),
		field.Switch("status", "状态").
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择状态"),
			}).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Department) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.Status(),
	}
}

// 行为
func (p *Department) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateModal(),
		actions.ChangeStatus(),
		actions.EditModal(),
		actions.DeleteSpecial(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
	}
}

// 列表页面显示前回调
func (p *Department) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok && search != nil {
		result := []interface{}{}
		for _, v := range list {
			result = append(result, v)
		}

		return result
	}
	// 转换成树形表格
	tree, _ := lister.ListToTree(list, "id", "pid", "children", 0)
	return tree
}
