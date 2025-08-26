package resource

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

// 列表工具栏菜单项，示例如下：
//
//	[]map[string]string{
//		{
//			"key":   "day",
//			"label": "日账单",
//		},
//		{
//			"key":   "week",
//			"label": "周账单",
//		},
//	}
func (p *Template) IndexTableMenuItems(ctx *quark.Context) []map[string]string {
	template := ctx.Template.(types.Resourcer)
	return template.MenuItems(ctx)
}

// 列表工具栏菜单，示例如下：
//
//	map[string]interface{}{
//		"type": "tab",
//		"items": []map[string]string{
//			{
//				"key":   "day",
//				"label": "日账单",
//			},
//			{
//				"key":   "week",
//				"label": "周账单",
//			},
//		},
//	}
func (p *Template) IndexTableMenu(ctx *quark.Context) interface{} {
	template := ctx.Template.(types.Resourcer)
	items := template.IndexTableMenuItems(ctx)
	if items == nil {
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"type":  "tab",
		"items": items,
	}
}
