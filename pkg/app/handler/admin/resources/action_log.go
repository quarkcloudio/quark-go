package resources

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"gorm.io/gorm"
)

type ActionLog struct {
	adminresource.Template
}

// 初始化
func (p *ActionLog) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "操作日志"

	// 模型
	p.Model = &model.ActionLog{}

	// 分页
	p.PerPage = 10

	p.WithExport = true

	return p
}

// 列表查询
func (p *ActionLog) Query(request *builder.Request, query *gorm.DB) *gorm.DB {

	return query.
		Select("action_logs.*,admins.username").
		Joins("left join admins on admins.id = action_logs.object_id").
		Where("type = ?", "admin")
}

// 字段
func (p *ActionLog) Fields(request *builder.Request) []interface{} {
	field := &builder.AdminField{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("username", "用户"),
		field.Text("url", "行为").
			SetColumn(func(column *table.Column) *table.Column {
				return column.SetEllipsis(true)
			}),
		field.Text("ip", "IP"),
		field.Datetime("created_at", "发生时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}),
	}
}

// 搜索
func (p *ActionLog) Searches(request *builder.Request) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("url", "行为"),
		(&searches.Input{}).Init("ip", "IP"),
	}
}

// 行为
func (p *ActionLog) Actions(request *builder.Request) []interface{} {
	return []interface{}{
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Delete{}).Init("删除"),
	}
}
