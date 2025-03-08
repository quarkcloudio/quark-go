package actions

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
)

type BatchExportAction struct {
	actions.Link
}

// 批量导出，BatchExport() | BatchExport("批量导出")
func BatchExport(options ...interface{}) *BatchExportAction {
	action := &BatchExportAction{}

	action.Name = "批量导出"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *BatchExportAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 设置按钮跳转链接
	p.Target = "_blank"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要导出数据吗？", "导出数据可能会等待时间较长！", "modal")

	// 在表格多选弹出层展示
	p.SetOnlyOnIndexTableAlert(true)

	return p
}

// 跳转链接
func (p *BatchExportAction) GetHref(ctx *quark.Context) string {
	return strings.Replace(ctx.Path(), "/index", "/export?id=${id}&token="+ctx.Token(), -1)
}
