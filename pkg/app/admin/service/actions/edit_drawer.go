package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type EditDrawer struct {
	actions.Drawer
}

// 初始化
func (p *EditDrawer) Init(name string) *EditDrawer {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 文字
	p.Name = name

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *EditDrawer) GetBody(ctx *builder.Context) interface{} {

	api := ctx.Template.(interface {
		UpdateApi(*builder.Context) string
	}).UpdateApi(ctx)

	initApi := ctx.Template.(interface {
		EditValueApi(*builder.Context) string
	}).EditValueApi(ctx)

	fields := ctx.Template.(interface {
		UpdateFieldsWithinComponents(*builder.Context) interface{}
	}).UpdateFieldsWithinComponents(ctx)

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
func (p *EditDrawer) GetActions(ctx *builder.Context) []interface{} {

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
