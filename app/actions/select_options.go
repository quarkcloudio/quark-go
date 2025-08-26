package actions

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type SelectOptionsAction struct {
	actions.Action
}

// 执行行为句柄
func (p *SelectOptionsAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	return ctx.CJSONOk("操作成功")
}
