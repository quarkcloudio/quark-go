package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
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

	return ctx.JSONOk("操作成功")
}
