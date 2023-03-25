package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"github.com/quarkcms/quark-go/pkg/component/admin/action"
	"github.com/quarkcms/quark-go/pkg/component/admin/form"
)

type CreateDrawer struct {
	actions.Drawer
}

// 初始化
func (p *CreateDrawer) Init(name string) *CreateDrawer {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 文字
	p.Name = "创建" + name

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
