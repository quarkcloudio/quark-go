package resources

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type Picture struct {
	resource.Template
}

// 初始化
func (p *Picture) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "图片"

	// 模型
	p.Model = &model.Picture{}

	// 分页
	p.PageSize = 10

	return p
}

// 字段
func (p *Picture) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("path", "显示", func() interface{} {
			return "<img src='" + service.NewPictureService().GetPath(p.Field["id"]) + "' width=50 height=50 />"
		}),
		field.Text("name", "名称").SetEllipsis(true),
		field.Text("size", "大小").SetSorter(true),
		field.Text("width", "宽度"),
		field.Text("height", "高度"),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间"),
	}
}

// 搜索
func (p *Picture) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.DatetimeRange("created_at", "上传时间"),
	}
}

// 行为
func (p *Picture) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
