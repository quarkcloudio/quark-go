package resources

import (
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Picture struct {
	resource.Template
}

// 初始化
func (p *Picture) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "图片"

	// 模型
	p.Model = &model.Picture{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Picture) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("path", "显示", func() interface{} {

			return "<img src='" + (&model.Picture{}).GetPath(p.Field["id"]) + "' width=50 height=50 />"
		}),
		field.Text("name", "名称").SetEllipsis(true),
		field.Text("size", "大小").SetSorter(true),
		field.Text("width", "宽度"),
		field.Text("height", "高度"),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}),
	}
}

// 搜索
func (p *Picture) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.DatetimeRange("created_at", "上传时间"),
	}
}

// 行为
func (p *Picture) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
