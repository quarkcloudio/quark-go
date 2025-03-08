package resources

import (
	"encoding/json"
	"fmt"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"gorm.io/gorm"
)

type Image struct {
	resource.Template
}

// 初始化
func (p *Image) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "图片"

	// 模型
	p.Model = &model.Attachment{}

	// 分页
	p.PageSize = 10

	return p
}

// 列表查询
func (p *Image) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("type = ?", "IMAGE")
}

// 字段
func (p *Image) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("path", "显示", func(row map[string]interface{}) interface{} {
			return "<img src='" + service.NewAttachmentService().GetImageUrl(row["id"]) + "' width=50 height=50 />"
		}),
		field.Text("name", "名称").SetEllipsis(true),
		field.Text("size", "大小").SetSorter(true),
		field.Text("extra", "尺寸", func(row map[string]interface{}) interface{} {
			var extra map[string]interface{}
			var extraInfo string
			if row["extra"] != "" {
				err := json.Unmarshal([]byte(row["extra"].(string)), &extra)
				if err == nil && extra["width"] != nil && extra["height"] != nil {
					extraInfo = fmt.Sprintf("%d*%d", int(extra["width"].(float64)), int(extra["height"].(float64)))
				}
			}
			return extraInfo
		}),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间"),
	}
}

// 搜索
func (p *Image) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.DatetimeRange("created_at", "上传时间"),
	}
}

// 行为
func (p *Image) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
