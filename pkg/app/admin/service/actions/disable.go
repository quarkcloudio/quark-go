package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Disable struct {
	actions.Action
}

// 初始化
func (p *Disable) Init(name string) *Disable {
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

	// 设置展示位置
	p.SetOnlyOnIndexTableAlert(true)

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要禁用吗？", "禁用后数据将无法使用，请谨慎操作！", "modal")

	return p
}

/**
 * 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
 *
 * @return array
 */
func (p *Disable) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *Disable) Handle(ctx *builder.Context, query *gorm.DB) error {
	err := query.Update("status", 0).Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("操作成功")
}
