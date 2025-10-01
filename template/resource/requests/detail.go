package requests

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

type DetailRequest struct{}

// 表单数据
func (p *DetailRequest) FillData(ctx *quark.Context) map[string]interface{} {
	result := map[string]interface{}{}
	id := ctx.Query("id", "")
	if id == "" {
		return result
	}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// Gorm对象
	model := db.Client.Model(&modelInstance)

	// 创建详情页查询
	query := template.BuildDetailQuery(ctx, model)

	// 查询数据
	query.First(&result)

	// 获取字段
	detailFields := template.DetailFields(ctx)

	// 解析字段值
	fields := make(map[string]interface{})
	for _, field := range detailFields.([]interface{}) {

		// 组件名称
		component := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Component").
			String()

		// 字段名
		name := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Name").
			String()

		if component == "actionField" {
			// 行为项
			actionItems := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Items").
				Interface()

			callback := field.(interface{ GetCallback() interface{} }).GetCallback()
			if callback != nil {
				actionItems = callback.(func(map[string]interface{}) interface{})(result)
			}

			var items []interface{}

			// 解析行为
			for _, action := range actionItems.([]interface{}) {

				actionInstance := action.(types.Actioner)

				// 初始化模版
				actionInstance.New(ctx)

				// 初始化
				actionInstance.Init(ctx)

				items = append(items, template.BuildAction(ctx, actionInstance))
			}

			fields[name] = items
		} else {
			callback := field.(interface{ GetCallback() interface{} }).GetCallback()
			if callback != nil {
				fields[name] = callback.(func(map[string]interface{}) interface{})(result)
			} else {
				if result[name] != nil {
					var fieldValue interface{}
					fieldValue = result[name]
					getV, ok := result[name].(string)
					if ok {
						if strings.Contains(getV, "[") {
							var m []interface{}
							err := json.Unmarshal([]byte(getV), &m)
							if err == nil {
								fieldValue = m
							} else {
								if strings.Contains(getV, "{") {
									var m map[string]interface{}
									err := json.Unmarshal([]byte(getV), &m)
									if err == nil {
										fieldValue = m
									}
								}
							}
						} else {
							if strings.Contains(getV, "{") {
								var m map[string]interface{}
								err := json.Unmarshal([]byte(getV), &m)
								if err == nil {
									fieldValue = m
								}
							}
						}
					}

					if component == "datetimeField" || component == "dateField" {
						format := reflect.
							ValueOf(field).
							Elem().
							FieldByName("Format").
							String()

						format = strings.Replace(format, "YYYY", "2006", -1)
						format = strings.Replace(format, "MM", "01", -1)
						format = strings.Replace(format, "DD", "02", -1)
						format = strings.Replace(format, "HH", "15", -1)
						format = strings.Replace(format, "mm", "04", -1)
						format = strings.Replace(format, "ss", "05", -1)

						fieldValue = result[name].(time.Time).Format(format)
					}

					fields[name] = fieldValue
				}
			}
		}
	}

	return fields
}

// 获取表单初始化数据
func (p *DetailRequest) Values(ctx *quark.Context) error {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取赋值数据
	data := p.FillData(ctx)

	// 显示前回调
	data = template.BeforeDetailShowing(ctx, data)

	return ctx.JSONOk("获取成功", data)
}
