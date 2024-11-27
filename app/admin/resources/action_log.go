package resources

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"gorm.io/gorm"
)

type ActionLog struct {
	resource.Template
}

// 初始化
func (p *ActionLog) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "操作日志"

	// 模型
	p.Model = &model.ActionLog{}

	// 分页
	p.PageSize = 10

	// 是否具有导出功能
	p.WithExport = true

	return p
}

// 列表查询
func (p *ActionLog) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {

	return query.
		Select("action_logs.*,users.username").
		Joins("left join users on users.id = action_logs.object_id").
		Where("type = ?", "admin")
}

// 字段
func (p *ActionLog) Fields(ctx *quark.Context) []interface{} {
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
func (p *ActionLog) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("url", "行为"),
		searches.Input("ip", "IP"),
	}
}

// 行为
func (p *ActionLog) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
