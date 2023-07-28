package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type BatchEnableAction struct {
	actions.Action
}

// 批量启用
func BatchEnable() *BatchEnableAction {
	return &BatchEnableAction{}
}

// 初始化
func (p *BatchEnableAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "批量启用"

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndexTableAlert(true)

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要启用吗？", "启用后数据将正常使用！", "modal")

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *BatchEnableAction) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *BatchEnableAction) Handle(ctx *builder.Context, model *gorm.DB) error {
	err := model.Update("status", 1).Error
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
