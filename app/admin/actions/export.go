package actions

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
)

type ExportAction struct {
	actions.Link
}

// 导出-跳转类型，Export() | Export("导出")
func Export(options ...interface{}) *ExportAction {
	action := &ExportAction{}

	// 文字
	action.Name = "导出"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *ExportAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 设置按钮跳转链接
	p.Target = "_blank"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *ExportAction) GetHref(ctx *quark.Context) string {
	return strings.Replace(ctx.Path(), "/index", "/export?id=${id}&token="+ctx.Token(), -1)
}
