package actions

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
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
func (p *ModalFormAction) Init(ctx *quark.Context) interface{} {

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
func (p *ModalFormAction) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Text("id", "ID"),

		field.Text("name", "名称").
			SetRules([]rule.Rule{
				rule.Required("名称必须填写"),
			}),
	}
}

// 表单数据（异步获取）
func (p *ModalFormAction) Data(ctx *quark.Context) map[string]interface{} {
	id := ctx.Query("id")

	return map[string]interface{}{
		"id": id,
	}
}

// 执行行为句柄
func (p *ModalFormAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	return ctx.CJSONError("method not implemented")
}
