package resources

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/actions"
	"github.com/quarkcloudio/quark-go/v4/app/searches"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
)

type Permission struct {
	resource.Template
}

// 初始化
func (p *Permission) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "权限"

	// 模型
	p.Model = &model.Permission{}

	// 分页
	p.PageSize = 10

	return p
}

// 字段
func (p *Permission) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称").
			SetRules([]rule.Rule{
				rule.Required("名称必须填写"),
			}),
		field.Text("path", "路径").
			SetRules([]rule.Rule{
				rule.Required("路径必须填写"),
			}),
		field.Select("method", "方法").
			SetOptions([]selectfield.Option{
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
func (p *Permission) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.Input("path", "路径"),
	}
}

// 行为
func (p *Permission) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.CreateModal(),
		actions.SyncPermission(),
		actions.BatchDelete(),
		actions.EditModal(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
