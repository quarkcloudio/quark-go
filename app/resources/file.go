package resources

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/actions"
	"github.com/quarkcloudio/quark-go/v4/app/searches"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
	"gorm.io/gorm"
)

type File struct {
	resource.Template
}

// 初始化
func (p *File) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "文件"

	// 模型
	p.Model = &model.Attachment{}

	// 分页
	p.PageSize = 10

	return p
}

// 列表查询
func (p *File) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type = ?", "FILE")
}

// 字段
func (p *File) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称"),
		field.Text("size", "大小").SetSorter(true),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间"),
	}
}

// 搜索
func (p *File) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.DatetimeRange("created_at", "上传时间"),
	}
}

// 行为
func (p *File) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
