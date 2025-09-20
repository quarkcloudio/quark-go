package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/action"
	"github.com/quarkcloudio/quark-go/v4/component/form"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

type CreateDrawerAction struct {
	actions.Drawer
}

// 创建-抽屉类型
func CreateDrawer() *CreateDrawerAction {
	return &CreateDrawerAction{}
}

// 初始化
func (p *CreateDrawerAction) Init(ctx *quark.Context) interface{} {

	// 文字
	p.Name = "新增"

	// 类型
	p.Type = "primary"

	// 透明
	p.Ghost = true

	// 图标
	p.Icon = "ant-design:plus-outlined"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *CreateDrawerAction) GetBody(ctx *quark.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 包裹在组件内的编辑页字段
	api := template.CreationApi(ctx)

	// 包裹在组件内的创建页字段
	fields := template.CreationFieldsWithinComponents(ctx)

	// 创建页面显示前回调
	data := template.BeforeCreating(ctx)

	// 返回数据
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
func (p *CreateDrawerAction) GetActions(ctx *quark.Context) []interface{} {

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
