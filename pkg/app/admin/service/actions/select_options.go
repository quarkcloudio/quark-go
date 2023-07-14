package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type SelectOptions struct {
	actions.Action
}

// 执行行为句柄
func (p *SelectOptions) Handle(ctx *builder.Context, query *gorm.DB) error {
	resource := ctx.Param("resource")
	search := ctx.Query("search")
	lists := []map[string]interface{}{}
	results := []map[string]interface{}{}

	switch resource {
	case "Some Field":
		query.Where("Some Field = ?", search).Find(&lists)
		for _, v := range lists {
			item := map[string]interface{}{
				"label": v["name"],
				"value": v["id"],
			}

			results = append(results, item)
		}
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
