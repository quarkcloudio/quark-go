package resources

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/gorm"
)

type ActionLog struct {
	resource.Template
}

// 初始化
func (p *ActionLog) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "操作日志"

	// 模型
	p.Model = &model.ActionLog{}

	// 分页
	p.PerPage = 10

	// 是否具有导出功能
	p.WithExport = true

	return p
}

// 列表查询
func (p *ActionLog) Query(ctx *builder.Context, query *gorm.DB) *gorm.DB {

	return query.
		Select("action_logs.*,admins.username").
		Joins("left join admins on admins.id = action_logs.object_id").
		Where("type = ?", "admin")
}

// 字段
func (p *ActionLog) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("username", "用户"),
		field.Text("url", "行为").SetEllipsis(true),
		field.Text("ip", "IP"),
		field.Datetime("created_at", "发生时间"),
	}
}

// 搜索
func (p *ActionLog) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("url", "行为"),
		searches.Input("ip", "IP"),
	}
}

// 行为
func (p *ActionLog) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
