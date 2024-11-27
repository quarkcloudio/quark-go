package resources

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/table"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type Department struct {
	resource.Template
}

// 初始化
func (p *Department) Init(ctx *quark.Context) interface{} {

	p.Table.
		SetExpandable(&table.Expandable{
			DefaultExpandedRowKeys: []interface{}{1},
		})

	// 标题
	p.Title = "部门"

	// 模型
	p.Model = &model.Department{}

	// 列表页数据转换成树
	p.TableListToTree = true

	// 默认排序
	p.IndexQueryOrder = "sort asc, id asc"

	// 分页
	p.PageSize = false

	return p
}

// 字段
func (p *Department) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 列表
	departments, _ := service.NewDepartmentService().GetList()

	return []interface{}{
		field.Hidden("id", "ID"),                 // 列表读取且不展示的字段
		field.Hidden("pid", "PID").OnlyOnIndex(), // 列表读取且不展示的字段
		field.Text("name", "名称").
			SetRules([]rule.Rule{
				rule.Required("名称必须填写"),
				rule.Min(2, "名称不能少于2个字符"),
				rule.Max(100, "名称不能超过100个字符"),
			}),
		field.TreeSelect("pid", "父节点").
			SetTreeData(departments, "pid", "name", "id").
			SetRules([]rule.Rule{
				rule.Required("请选择父节点"),
			}).
			SetDefault(1).
			OnlyOnCreating(),
		field.Dependency().SetWhen("id", "!=", 1, func() interface{} {
			return field.TreeSelect("pid", "父节点").
				SetTreeData(departments, "pid", "name", "id").
				SetRules([]rule.Rule{
					rule.Required("请选择父节点"),
				}).
				SetDefault(1).
				OnlyOnUpdating()
		}),
		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),
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
func (p *Department) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.Status(),
	}
}

// 行为
func (p *Department) Actions(ctx *quark.Context) []interface{} {
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
