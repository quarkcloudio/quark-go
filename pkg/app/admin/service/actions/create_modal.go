package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type CreateModalAction struct {
	actions.Modal
}

// 创建-弹窗类型
func CreateModal() *CreateModalAction {
	return &CreateModalAction{}
}

// 初始化
func (p *CreateModalAction) Init(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 文字
	p.Name = "创建" + template.GetTitle()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *CreateModalAction) GetBody(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 创建表单的接口
	api := template.CreationApi(ctx)

	// 包裹在组件内的创建页字段
	fields := template.CreationFieldsWithinComponents(ctx)

	// 创建页面显示前回调
	data := template.BeforeCreating(ctx)

	// 返回数据
	return (&form.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"paddingTop": "24px",
		}).
		SetKey("createModalForm", false).
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
func (p *CreateModalAction) GetActions(ctx *builder.Context) []interface{} {

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
			SetSubmitForm("createModalForm"),
	}
}
