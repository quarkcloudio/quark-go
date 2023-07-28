package actions

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Link struct {
	Action
	Href   string `json:"href"`
	Target string `json:"target"`
}

// 初始化
func (p *Link) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "link"
	p.Target = "_self"

	return p
}

// 获取跳转链接
func (p *Link) GetHref(ctx *builder.Context) string {
	return p.Href
}

// 相当于 a 链接的 target 属性，href 存在时生效
func (p *Link) GetTarget(ctx *builder.Context) string {
	return p.Target
}
