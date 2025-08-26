package requests

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

type IndexRequest struct{}

// 列表查询
func (p *IndexRequest) QueryData(ctx *quark.Context) interface{} {
	var lists []map[string]interface{}

	template := ctx.Template.(types.Resourcer)

	modelInstance := template.GetModel()

	model := db.Client.Model(modelInstance)

	// 搜索项
	searches := template.Searches(ctx)

	// 过滤项，预留
	filters := template.Filters(ctx)

	query := template.BuildIndexQuery(ctx, model, searches, filters, p.columnFilters(ctx), p.orderings(ctx))

	// 获取分页
	pageSize := template.GetPageSize()

	// 指定每页可以显示多少条，[10, 20, 50, 100]
	pageSizeOptions := template.GetPageSizeOptions()
	if pageSize == nil {
		query.Find(&lists)

		// 返回解析列表
		return p.performsList(ctx, lists)
	}

	// 不分页，直接返回lists
	if reflect.TypeOf(pageSize).String() != "int" {
		query.Find(&lists)

		// 返回解析列表
		return p.performsList(ctx, lists)
	}

	var total int64
	var data map[string]interface{}
	page := 1
	querys := ctx.AllQuerys()
	if querys["search"] != nil {
		err := json.Unmarshal([]byte(querys["search"].(string)), &data)
		if err == nil {
			if data["current"] != nil {
				page = int(data["current"].(float64))
			}
			if data["pageSize"] != nil {
				pageSize = int(data["pageSize"].(float64))
			}
		}
	}

	// 获取总数量
	query.Count(&total)

	// 获取列表
	query.Limit(pageSize.(int)).Offset((page - 1) * pageSize.(int)).Find(&lists)

	// 解析列表
	result := p.performsList(ctx, lists)

	return map[string]interface{}{
		"page":            page,
		"pageSize":        pageSize,
		"pageSizeOptions": pageSizeOptions,
		"total":           total,
		"items":           result,
	}
}

// Get the column filters for the request.
func (p *IndexRequest) columnFilters(ctx *quark.Context) map[string]interface{} {
	querys := ctx.AllQuerys()
	var data map[string]interface{}
	if querys["filter"] == nil {
		return data
	}
	err := json.Unmarshal([]byte(querys["filter"].(string)), &data)
	if err != nil {
		return data
	}

	return data
}

// Get the orderings for the request.
func (p *IndexRequest) orderings(ctx *quark.Context) map[string]interface{} {
	querys := ctx.AllQuerys()
	var data map[string]interface{}
	if querys["sorter"] == nil {
		return data
	}
	err := json.Unmarshal([]byte(querys["sorter"].(string)), &data)
	if err != nil {
		return data
	}

	return data
}

// 处理列表
func (p *IndexRequest) performsList(ctx *quark.Context, lists []map[string]interface{}) []interface{} {
	result := []map[string]interface{}{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取列表字段
	indexFields := template.IndexFields(ctx)

	// 解析字段回调函数
	for _, v := range lists {
		fields := make(map[string]interface{})
		for _, field := range indexFields.([]interface{}) {

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
					actionItems = callback.(func(map[string]interface{}) interface{})(v)
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
					fields[name] = callback.(func(map[string]interface{}) interface{})(v)
				} else {
					if v[name] != nil {
						var fieldValue interface{}
						fieldValue = v[name]
						getV, ok := v[name].(string)
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

						// 单独解析时间和日期组件
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
							fieldValue = v[name].(time.Time).Format(format)
						}

						// 单独解析图片、图片选择器组件
						if component == "imageField" || component == "imagePickerField" {
							fieldValue = service.NewAttachmentService().GetImageUrl(v[name].(string))
						}

						fields[name] = fieldValue
					}
				}
			}
		}

		result = append(result, fields)
	}

	// 列表显示前回调
	return template.BeforeIndexShowing(ctx, result)
}
