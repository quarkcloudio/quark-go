package resources

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
)

type Permission struct {
	adminresource.Template
}

// 初始化
func (p *Permission) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

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
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "名称必须填写",
				},
			),

		field.Text("guard_name", "GuardName").SetDefault("admin"),
		field.Datetime("created_at", "创建时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}).OnlyOnIndex(),
		field.Datetime("updated_at", "更新时间", func() interface{} {
			if p.Field["updated_at"] == nil {
				return p.Field["updated_at"]
			}

			return p.Field["updated_at"].(time.Time).Format("2006-01-02 15:04:05")
		}).OnlyOnIndex(),
	}
}

// 搜索
func (p *Permission) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
	}
}

// 行为
func (p *Permission) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&actions.SyncPermission{}).Init(),
		(&actions.CreateModal{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.EditModal{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
