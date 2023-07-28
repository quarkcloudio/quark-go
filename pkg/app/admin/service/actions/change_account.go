package actions

import (
	"encoding/json"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/utils/hash"
	"gorm.io/gorm"
)

type ChangeAccountAction struct {
	actions.Action
}

// 更改账户信息
func ChangeAccount() *ChangeAccountAction {
	return &ChangeAccountAction{}
}

// 执行行为句柄
func (p *ChangeAccountAction) Handle(ctx *builder.Context, query *gorm.DB) error {
	data := map[string]interface{}{}
	ctx.Bind(&data)

	if data["avatar"] != "" && data["avatar"] != nil {
		data["avatar"], _ = json.Marshal(data["avatar"])
	}

	// 加密密码
	if data["password"] != nil {
		data["password"] = hash.Make(data["password"].(string))
	}

	// 获取登录管理员信息
	adminInfo, err := (&model.Admin{}).GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	err = query.Where("id", adminInfo.Id).Updates(data).Error
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
