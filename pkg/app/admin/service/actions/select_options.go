package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type SelectOptions struct {
	actions.Action
}

// 执行行为句柄
func (p *SelectOptions) Handle(ctx *builder.Context, query *gorm.DB) error {
	return ctx.JSON(200, message.Success("操作成功"))
}
