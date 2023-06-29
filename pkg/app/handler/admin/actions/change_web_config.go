package actions

import (
	"encoding/json"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

type ChangeWebConfig struct {
	actions.Action
}

// 执行行为句柄
func (p *ChangeWebConfig) Handle(ctx *builder.Context, query *gorm.DB) interface{} {
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
		return ctx.JSON(200, msg.Error("操作失败，请重试！", ""))
	}

	// 刷新网站配置
	(&model.Config{}).Refresh()

	// 返回成功
	return ctx.JSON(200, msg.Success("操作成功", "", ""))
}
