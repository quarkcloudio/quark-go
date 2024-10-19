package resources

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
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
				field.SelectOption("Any", "Any"),
				field.SelectOption("GET", "GET"),
				field.SelectOption("HEAD", "HEAD"),
				field.SelectOption("OPTIONS", "OPTIONS"),
				field.SelectOption("POST", "POST"),
				field.SelectOption("PUT", "PUT"),
				field.SelectOption("PATCH", "PATCH"),
				field.SelectOption("DELETE", "DELETE"),
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
		searches.Input("path", "路径"),
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
