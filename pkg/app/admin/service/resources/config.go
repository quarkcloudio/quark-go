package resources

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Config struct {
	resource.Template
}

// 初始化
func (p *Config) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

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
				{
					Value: "text",
					Label: "输入框",
				},
				{
					Value: "textarea",
					Label: "文本域",
				},
				{
					Value: "picture",
					Label: "图片",
				},
				{
					Value: "file",
					Label: "文件",
				},
				{
					Value: "switch",
					Label: "开关",
				},
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
		(&searches.Input{}).Init("title", "标题"),
		(&searches.Input{}).Init("name", "名称"),
		(&searches.Status{}).Init(),
	}
}

// 行为
func (p *Config) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&actions.CreateDrawer{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.ChangeStatus{}).Init(),
		(&actions.EditDrawer{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
