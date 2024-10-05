package actions

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
)

type EditLinkAction struct {
	actions.Link
}

// 编辑-跳转类型，EditLink() | EditLink("编辑")
func EditLink(options ...interface{}) *EditLinkAction {
	action := &EditLinkAction{}

	// 文字
	action.Name = "编辑"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *EditLinkAction) Init(ctx *builder.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *EditLinkAction) GetHref(ctx *builder.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/index", "/edit&id=${id}", -1)
}
