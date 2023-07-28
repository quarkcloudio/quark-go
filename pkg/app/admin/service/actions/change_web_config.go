package actions

import (
	"encoding/json"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

type ChangeWebConfigAction struct {
	actions.Action
}

// 更改网站配置
func ChangeWebConfig() *ChangeWebConfigAction {
	return &ChangeWebConfigAction{}
}

// 执行行为句柄
func (p *ChangeWebConfigAction) Handle(ctx *builder.Context, query *gorm.DB) error {
	data := map[string]interface{}{}
	json.Unmarshal(ctx.Body(), &data)
	result := true

	for k, v := range data {
		config := map[string]interface{}{}
		db.Client.Model(&model.Config{}).Where("name =?", k).First(&config)
		if getValue, ok := v.([]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		if getValue, ok := v.([]map[string]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		if getValue, ok := v.(map[string]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		updateResult := db.Client.Model(&model.Config{}).Where("name", k).Update("value", v)
		if updateResult.Error != nil {
			result = false
		}
	}

	if !result {
		return ctx.JSON(200, message.Error("操作失败，请重试！"))
	}

	// 刷新网站配置
	(&model.Config{}).Refresh()

	// 返回成功
	return ctx.JSON(200, message.Success("操作成功"))
}
