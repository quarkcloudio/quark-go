package actions

import (
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/actions"
)

type EditLink struct {
	actions.Link
}

// 初始化
func (p *EditLink) Init(name string) *EditLink {
	// 初始化父结构
	p.ParentInit()

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 文字
	p.Name = name

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *EditLink) GetHref(request *builder.Request) string {
	return "#/index?api=" + strings.Replace(request.Path(), "/index", "/edit&id=${id}", -1)
}
