package resource

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

// 列表工具栏
func (p *Template) IndexTableMenus(ctx *builder.Context) interface{} {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	menus := template.Menus(ctx)

	return menus
}
