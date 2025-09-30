package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type StatusAction struct {
	actions.Switch
}

// 更改状态，Status()
func Status() *StatusAction {
	action := &StatusAction{}
	return action
}

// 初始化
func (p *StatusAction) Init(ctx *quark.Context) interface{} {

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 选中时的内容
	p.CheckedChildren = "正常"

	// 未选中时的内容
	p.UnCheckedChildren = "禁用"

	// 设置字段名
	p.FieldName = "status"

	// 设置字段值
	p.FieldValue = "1"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *StatusAction) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *StatusAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	status := ctx.Query("status")
	if status == "" {
		return ctx.JSONError("参数错误")
	}

	var fieldStatus int
	if status == "0" || status == "false" {
		fieldStatus = 0
	} else {
		fieldStatus = 1
	}

	err := query.Update("status", fieldStatus).Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("操作成功")
}
