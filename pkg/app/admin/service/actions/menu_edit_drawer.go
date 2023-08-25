package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type MenuEditDrawerAction struct {
	actions.Drawer
}

// 编辑-抽屉类型，EditDrawer() | EditDrawer("编辑")
func MenuEditDrawer(options ...interface{}) *MenuEditDrawerAction {
	action := &MenuEditDrawerAction{}

	// 文字
	action.Name = "编辑"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *MenuEditDrawerAction) Init(ctx *builder.Context) interface{} {

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 抽屉弹出层宽度
	p.Width = 750

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *MenuEditDrawerAction) GetBody(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 更新表单的接口
	api := template.UpdateApi(ctx)

	// 编辑页面获取表单数据接口
	initApi := template.EditValueApi(ctx)

	// 包裹在组件内的编辑页字段
	fields := template.UpdateFieldsWithinComponents(ctx)

	// 返回数据
	return (&form.Component{}).
		Init().
		SetKey("editDrawerForm", false).
		SetLayout("vertical").
		SetApi(api).
		SetInitApi(initApi).
		SetBody(fields)
}

// 弹窗行为
func (p *MenuEditDrawerAction) GetActions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&action.Component{}).
			Init().
			SetLabel("取消").
			SetActionType("cancel"),

		(&action.Component{}).
			Init().
			SetLabel("提交").
			SetWithLoading(true).
			SetReload("table").
			SetActionType("submit").
			SetType("primary", false).
			SetSubmitForm("editDrawerForm"),
	}
}
