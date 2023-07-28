package resources

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Permission struct {
	resource.Template
}

// 初始化
func (p *Permission) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "权限"

	// 模型
	p.Model = &model.Permission{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Permission) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
			}),

		field.Text("path", "路径").
			SetRules([]*rule.Rule{
				rule.Required(true, "路径必须填写"),
			}),

		field.Select("method", "方法").
			SetOptions([]*selectfield.Option{
				{
					Value: "Any",
					Label: "Any",
				},
				{
					Value: "GET",
					Label: "GET",
				},
				{
					Value: "HEAD",
					Label: "HEAD",
				},
				{
					Value: "OPTIONS",
					Label: "OPTIONS",
				},
				{
					Value: "POST",
					Label: "POST",
				},
				{
					Value: "PUT",
					Label: "PUT",
				},
				{
					Value: "PATCH",
					Label: "PATCH",
				},
				{
					Value: "DELETE",
					Label: "DELETE",
				},
			}).
			SetFilters(true).
			SetDefault("GET"),
		field.TextArea("remark", "备注"),
	}
}

// 搜索
func (p *Permission) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
	}
}

// 行为
func (p *Permission) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.SyncPermission(),
		actions.CreateModal(),
		actions.BatchDelete(),
		actions.EditModal(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
