package actions

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/gorm"
)

type SelectOptionsAction struct {
	actions.Action
}

// 执行行为句柄
func (p *SelectOptionsAction) Handle(ctx *builder.Context, query *gorm.DB) error {
	return ctx.JSON(200, message.Success("操作成功"))
}
