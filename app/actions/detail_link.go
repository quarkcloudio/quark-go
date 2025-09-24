package actions

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
)

type DetailLinkAction struct {
	actions.Link
}

// 跳转详情页，DetailLink() | DetailLink("详情")
func DetailLink(options ...interface{}) *DetailLinkAction {
	action := &DetailLinkAction{}

	// 文字
	action.Name = "详情"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *DetailLinkAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *DetailLinkAction) GetHref(ctx *quark.Context) string {
	return "#${enginePath}?api=" + strings.Replace(ctx.Path(), "/index", "/detail?id=${id}", -1)
}
