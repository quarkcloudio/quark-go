package actions

import (
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
)

type CreateLink struct {
	actions.Link
}

// 初始化
func (p *CreateLink) Init(name string) *CreateLink {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 文字
	p.Name = "创建" + name

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 跳转链接
func (p *CreateLink) GetHref(ctx *builder.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/index", "/create", -1)
}
