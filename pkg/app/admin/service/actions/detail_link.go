package actions

import (
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type DetailLinkAction struct {
	actions.Link
}

// 跳转详情页
func DetailLink() *DetailLinkAction {
	return &DetailLinkAction{}
}

// 初始化
func (p *DetailLinkAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "详情"

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *DetailLinkAction) GetHref(ctx *builder.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/index", "/detail&id=${id}", -1)
}
