package actions

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
)

type CreateLinkAction struct {
	actions.Link
}

// 创建-跳转类型
func CreateLink() *CreateLinkAction {
	return &CreateLinkAction{}
}

// 初始化
func (p *CreateLinkAction) Init(ctx *quark.Context) interface{} {

	// 文字
	p.Name = "新增"

	// 类型
	p.Type = "primary"

	// 透明
	p.Ghost = true

	// 图标
	p.Icon = "ant-design:plus-outlined"

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 跳转链接
func (p *CreateLinkAction) GetHref(ctx *quark.Context) string {
	return "#/layout/index?api=" + strings.Replace(ctx.Path(), "/index", "/create", -1)
}
