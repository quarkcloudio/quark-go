package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type DeleteSpecialAction struct {
	actions.Action
}

// 删除，DeleteSpecial() | DeleteSpecial("删除")
func DeleteSpecial(options ...interface{}) *DeleteSpecialAction {
	action := &DeleteSpecialAction{}

	action.Name = "<%= (id==1 ? '' : '删除') %>"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *DeleteSpecialAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要删除吗？", "删除后数据将无法恢复，请谨慎操作！", "modal")

	// 在表格行内展示
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	p.SetApiParams([]string{
		"id",
	})

	return p
}

// 执行行为句柄
func (p *DeleteSpecialAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	err := query.Delete("").Error
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	return ctx.CJSONOk("操作成功")
}
