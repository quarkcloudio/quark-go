package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type ModalFormAction struct {
	actions.ModalForm
}

// 弹窗表单
func ModalForm() *ModalFormAction {
	return &ModalFormAction{}
}

// 初始化
func (p *ModalFormAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "Test"

	// 类型
	p.Type = "link"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数
	p.SetApiParams([]string{
		"id",
	})

	return p
}

// 字段
func (p *ModalFormAction) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Text("id", "ID"),

		field.Text("name", "名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
			}),
	}
}

// 表单数据（异步获取）
func (p *ModalFormAction) Data(ctx *builder.Context) map[string]interface{} {
	id := ctx.Query("id")

	return map[string]interface{}{
		"id": id,
	}
}

// 执行行为句柄
func (p *ModalFormAction) Handle(ctx *builder.Context, query *gorm.DB) error {

	return ctx.JSON(200, message.Error("Method not implemented"))
}
