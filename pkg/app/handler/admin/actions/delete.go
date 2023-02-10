package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/actions"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

type Delete struct {
	actions.Action
}

// 初始化
func (p *Delete) Init(name string) *Delete {
	// 初始化父结构
	p.ParentInit()

	// 行为名称，当行为在表格行展示时，支持js表达式
	p.Name = name

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要删除吗？", "删除后数据将无法恢复，请谨慎操作！", "modal")

	if name == "删除" {
		p.SetOnlyOnIndexTableRow(true)
	}

	if name == "批量删除" {
		p.SetOnlyOnIndexTableAlert(true)
	}

	return p
}

/**
 * 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
 *
 * @return array
 */
func (p *Delete) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *Delete) Handle(ctx *builder.Context, model *gorm.DB) interface{} {
	err := model.Delete("").Error
	if err != nil {
		return ctx.JSON(200, msg.Error(err.Error(), ""))
	}

	return ctx.JSON(200, msg.Success("操作成功", "", ""))
}
