package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type DeleteAction struct {
	actions.Action
}

// 删除
func Delete() *DeleteAction {
	return &DeleteAction{}
}

// 初始化
func (p *DeleteAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "删除"

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
func (p *DeleteAction) Handle(ctx *builder.Context, query *gorm.DB) error {
	err := query.Delete("").Error
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
