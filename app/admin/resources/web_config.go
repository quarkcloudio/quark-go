package resources

import (
	"encoding/json"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"gorm.io/gorm"
)

type WebConfig struct {
	resource.Template
}

// 初始化
func (p *WebConfig) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "网站配置"

	// 模型
	p.Model = &model.Config{}

	return p
}

// 字段
func (p *WebConfig) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	groupNames := []string{}

	db.Client.
		Model(p.Model).
		Where("status = ?", 1).
		Distinct("group_name").
		Pluck("group_name", &groupNames)

	tabPanes := []interface{}{}
	for _, groupName := range groupNames {
		configs := []map[string]interface{}{}
		db.Client.
			Model(p.Model).
			Where("status = ?", 1).
			Where("group_name = ?", groupName).
			Order("sort asc").
			Find(&configs)

		fields := []interface{}{}
		for _, config := range configs {
			remark, ok := config["remark"].(string)
			if !ok {
				remark = ""
			}
			switch config["type"] {
			case "text":
				getField := field.
					Text(config["name"], config["title"]).
					SetExtra(remark)
				fields = append(fields, getField)
			case "textarea":
				getField := field.
					TextArea(config["name"], config["title"]).
					SetExtra(remark)
				fields = append(fields, getField)
			case "file":
				getField := field.
					File(config["name"], config["title"]).
					SetButton("上传" + config["title"].(string)).
					SetExtra(remark)
				fields = append(fields, getField)
			case "picture":
				getField := field.
					Image(config["name"], config["title"]).
					SetButton("上传" + config["title"].(string)).
					SetExtra(remark)
				fields = append(fields, getField)
			case "switch":
				getField := field.
					Switch(config["name"].(string), config["title"].(string)).
					SetTrueValue("正常").
					SetFalseValue("禁用").
					SetExtra(remark)
				fields = append(fields, getField)
			}
		}
		tabPane := (&tabs.TabPane{}).
			Init().
			SetTitle(groupName).
			SetBody(fields)
		tabPanes = append(tabPanes, tabPane)
	}

	return tabPanes
}

// 行为
func (p *WebConfig) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 表单显示前回调
func (p *WebConfig) BeforeFormShowing(ctx *quark.Context) map[string]interface{} {
	configs := []map[string]interface{}{}
	data := map[string]interface{}{}

	db.Client.
		Model(p.Model).
		Where("status = ?", 1).
		Find(&configs)

	for _, config := range configs {
		data[config["name"].(string)] = config["value"]
		if config["type"] == "switch" {
			if config["value"] != "0" {
				data[config["name"].(string)] = true
			} else {
				data[config["name"].(string)] = false
			}
		}

		if config["type"] == "picture" || config["type"] == "file" {
			if config["value"] != nil && config["value"] != "" {
				// json字符串
				if strings.Contains(config["value"].(string), "{") {
					var jsonData interface{}
					json.Unmarshal([]byte(config["value"].(string)), &jsonData)

					// 如果为map
					if mapData, ok := jsonData.(map[string]interface{}); ok {
						data[config["name"].(string)] = mapData
					}

					// 如果为数组，返回第一个key的path
					if arrayData, ok := jsonData.([]map[string]interface{}); ok {
						data[config["name"].(string)] = arrayData
					}
				}
			} else {
				data[config["name"].(string)] = nil
			}
		}
	}

	return data
}

func (p *WebConfig) FormHandle(ctx *quark.Context, query *gorm.DB, data map[string]interface{}) error {
	result := true
	for k, v := range data {
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
		return ctx.CJSONError("操作失败，请重试")

	}

	// 刷新网站配置
	service.NewConfigService().Refresh()

	// 返回成功
	return ctx.CJSONOk("操作成功")
}
