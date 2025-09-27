package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type BatchDisableAction struct {
	actions.Batch
}

// 批量禁用，BatchDisable() | BatchDisable("批量禁用")
func BatchDisable(options ...interface{}) *BatchDisableAction {
	action := &BatchDisableAction{}

	action.Name = "批量禁用"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *BatchDisableAction) Init(ctx *quark.Context) interface{} {

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要禁用吗？", "禁用后数据将无法使用，请谨慎操作！", "modal")

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *BatchDisableAction) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *BatchDisableAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	err := query.Update("status", 0).Error
	if err != nil {
		return ctx.CJSONError(err.Error())
	}
	return ctx.CJSONOk("操作成功")
}
