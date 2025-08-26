package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

type DetailDrawerAction struct {
	actions.Drawer
}

// 详情-抽屉类型，DetailDrawer() | DetailDrawer("详情")
func DetailDrawer(options ...interface{}) *DetailDrawerAction {
	action := &DetailDrawerAction{}

	// 文字
	action.Name = "详情"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *DetailDrawerAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 宽度
	p.Width = 750

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *DetailDrawerAction) GetBody(ctx *quark.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 详情页面获取表单数据接口
	initApi := template.DetailValueApi(ctx)

	// 包裹在组件内的编辑页字段
	component := template.DetailFieldsWithinComponents(ctx, initApi, nil)

	// 返回数据
	return component
}
