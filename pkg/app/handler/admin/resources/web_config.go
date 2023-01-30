package resources

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type WebConfig struct {
	adminresource.Template
}

// 初始化
func (p *WebConfig) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "网站配置"

	// 模型
	p.Model = &models.Config{}

	return p
}

// 表单接口
func (p *WebConfig) FormApi(request *builder.Request) string {

	return "/api/admin/webConfig/action/change-web-config"
}

// 字段
func (p *WebConfig) Fields(request *builder.Request) []interface{} {
	field := &builder.AdminField{}
	groupNames := []string{}

	db.Client.Model(p.Model).Where("status = ?", 1).Distinct("group_name").Pluck("group_name", &groupNames)

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

			if ok == false {
				remark = ""
			}

			switch config["type"] {
			case "text":
				getField := field.
					Text(config["name"], config["title"]).SetExtra(remark)

				fields = append(fields, getField)

			case "textarea":
				getField := field.
					TextArea(config["name"], config["title"]).SetExtra(remark)

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
func (p *WebConfig) Actions(request *builder.Request) []interface{} {
	return []interface{}{
		(&actions.ChangeWebConfig{}),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}

// 创建页面显示前回调
func (p *WebConfig) BeforeCreating(request *builder.Request) map[string]interface{} {
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
		}
	}

	return data
}
