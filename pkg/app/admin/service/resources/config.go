package resources

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

type Config struct {
	resource.Template
}

// 初始化
func (p *Config) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "配置"

	// 模型
	p.Model = &model.Config{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Config) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("title", "标题").
			SetRules([]*rule.Rule{
				rule.Required(true, "标题必须填写"),
			}),

		field.Text("name", "名称").
			SetEditable(true).
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("configs", "name", "名称已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("configs", "name", "{id}", "名称已存在"),
			}),
		field.Radio("type", "表单类型").
			SetOptions([]*radio.Option{
				field.RadioOption("text", "文本"),
				field.RadioOption("textarea", "文本域"),
				field.RadioOption("picture", "图片"),
				field.RadioOption("file", "文件"),
				field.RadioOption("switch", "开关"),
			}).
			SetDefault("text").
			OnlyOnForms(),

		field.Text("sort", "排序").
			SetEditable(true).
			SetDefault(0).
			SetHelp("值越小越靠前").
			OnlyOnForms(),

		field.Text("group_name", "分组名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "分组名称必须填写"),
			}).OnlyOnForms(),

		field.Text("remark", "备注").
			OnlyOnForms(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Config) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("title", "标题"),
		searches.Input("name", "名称"),
		searches.Status(),
	}
}

// 行为
func (p *Config) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateDrawer(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.ChangeStatus(),
		actions.EditDrawer(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
