package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type EditDrawerAction struct {
	actions.Drawer
}

// 编辑-抽屉类型
func EditDrawer() *EditDrawerAction {
	return &EditDrawerAction{}
}

// 初始化
func (p *EditDrawerAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "编辑"

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *EditDrawerAction) GetBody(ctx *builder.Context) interface{} {
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
		SetApi(api).
		SetInitApi(initApi).
		SetBody(fields).
		SetLabelCol(map[string]interface{}{
			"span": 6,
		}).
		SetWrapperCol(map[string]interface{}{
			"span": 18,
		})
}

// 弹窗行为
func (p *EditDrawerAction) GetActions(ctx *builder.Context) []interface{} {

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
