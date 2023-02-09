package actions

import (
	"encoding/json"

	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/actions"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

type ChangeWebConfig struct {
	actions.Action
}

// 执行行为句柄
func (p *ChangeWebConfig) Handle(ctx *builder.Context, model *gorm.DB) interface{} {
	data := map[string]interface{}{}
	json.Unmarshal(ctx.Body(), &data)
	result := true

	for k, v := range data {
		config := map[string]interface{}{}
		db.Client.Model(&models.Config{}).Where("name =?", k).First(&config)
		if getValue, ok := v.([]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		if getValue, ok := v.([]map[string]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		if getValue, ok := v.(map[string]interface{}); ok {
			v, _ = json.Marshal(getValue)
		}
		updateResult := db.Client.Model(&models.Config{}).Where("name", k).Update("value", v)
		if updateResult.Error != nil {
			result = false
		}
	}

	if !result {
		return msg.Error("操作失败，请重试！", "")
	}

	// 刷新网站配置
	(&models.Config{}).Refresh()

	// 返回成功
	return msg.Success("操作成功", "", "")
}
