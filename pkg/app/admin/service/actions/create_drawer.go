package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type CreateDrawer struct {
	actions.Drawer
}

// 初始化
func (p *CreateDrawer) Init(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 文字
	p.Name = "创建" + template.GetTitle()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *CreateDrawer) GetBody(ctx *builder.Context) interface{} {

	api := ctx.Template.(interface {
		CreationApi(*builder.Context) string
	}).CreationApi(ctx)

	fields := ctx.Template.(interface {
		CreationFieldsWithinComponents(*builder.Context) interface{}
	}).CreationFieldsWithinComponents(ctx)

	// 断言BeforeCreating方法，获取初始数据
	data := ctx.Template.(interface {
		BeforeCreating(*builder.Context) map[string]interface{}
	}).BeforeCreating(ctx)

	return (&form.Component{}).
		Init().
		SetKey("createDrawerForm", false).
		SetApi(api).
		SetBody(fields).
		SetInitialValues(data).
		SetLabelCol(map[string]interface{}{
			"span": 6,
		}).
		SetWrapperCol(map[string]interface{}{
			"span": 18,
		})
}

// 弹窗行为
func (p *CreateDrawer) GetActions(ctx *builder.Context) []interface{} {

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
			SetSubmitForm("createDrawerForm"),
	}
}
