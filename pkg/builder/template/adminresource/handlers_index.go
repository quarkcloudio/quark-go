package adminresource

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type IndexRequest struct{}

// 列表查询
func (p *IndexRequest) QueryData(ctx *builder.Context) interface{} {
	var lists []map[string]interface{}
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	// 搜索项
	searches := ctx.Template.(interface {
		Searches(ctx *builder.Context) []interface{}
	}).Searches(ctx)

	// 过滤项，预留
	filters := ctx.Template.(interface {
		Filters(ctx *builder.Context) []interface{}
	}).Filters(ctx)

	query := ctx.Template.(interface {
		BuildIndexQuery(ctx *builder.Context, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB
	}).BuildIndexQuery(ctx, model, searches, filters, p.columnFilters(ctx), p.orderings(ctx))

	// 获取分页
	perPage := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("PerPage").Interface()

	// 不分页，直接返回lists
	if reflect.TypeOf(perPage).String() != "int" {
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
				perPage = int(data["pageSize"].(float64))
			}
		}
	}

	// 获取总数量
	query.Count(&total)

	// 获取列表
	query.Limit(perPage.(int)).Offset((page - 1) * perPage.(int)).Find(&lists)

	// 解析列表
	result := p.performsList(ctx, lists)

	return map[string]interface{}{
		"currentPage": page,
		"perPage":     perPage,
		"total":       total,
		"items":       result,
	}
}

/**
 * Get the column filters for the request.
 *
 * @return array
 */
func (p *IndexRequest) columnFilters(ctx *builder.Context) map[string]interface{} {
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

/**
 * Get the orderings for the request.
 *
 * @return array
 */
func (p *IndexRequest) orderings(ctx *builder.Context) map[string]interface{} {
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
func (p *IndexRequest) performsList(ctx *builder.Context, lists []map[string]interface{}) []interface{} {
	result := []map[string]interface{}{}

	// 获取列表字段
	indexFields := ctx.Template.(interface {
		IndexFields(ctx *builder.Context) interface{}
	}).IndexFields(ctx)

	// 解析字段回调函数
	for _, v := range lists {

		// 给实例的Field属性赋值
		ctx.Template.(interface {
			SetField(fieldData map[string]interface{}) interface{}
		}).SetField(v)

		fields := make(map[string]interface{})
		for _, field := range indexFields.([]interface{}) {

			// 字段名
			name := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Name").String()

			// 获取实例的回调函数
			callback := field.(interface{ GetCallback() interface{} }).GetCallback()

			if callback != nil {
				getCallback := callback.(func() interface{})
				fields[name] = getCallback()
			} else {
				if v[name] != nil {
					var fieldValue interface{}
					getV, ok := v[name].(string)
					if ok {
						if strings.Contains(getV, "{") {
							var m map[string]interface{}
							err := json.Unmarshal([]byte(getV), &m)
							if err != nil {
								fmt.Printf("Unmarshal with error: %+v\n", err)
							}
							fieldValue = m
						}
						if strings.Contains(getV, "[") {
							var m []interface{}
							err := json.Unmarshal([]byte(getV), &m)
							if err != nil {
								fmt.Printf("Unmarshal with error: %+v\n", err)
							}
							fieldValue = m
						}
					}

					// 组件名称
					component := reflect.
						ValueOf(field).
						Elem().
						FieldByName("Component").String()

					if component == "datetimeField" || component == "dateField" {
						format := reflect.
							ValueOf(field).
							Elem().
							FieldByName("Format").String()

						format = strings.Replace(format, "YYYY", "2006", -1)
						format = strings.Replace(format, "MM", "01", -1)
						format = strings.Replace(format, "DD", "02", -1)
						format = strings.Replace(format, "HH", "15", -1)
						format = strings.Replace(format, "mm", "04", -1)
						format = strings.Replace(format, "ss", "05", -1)

						fieldValue = v[name].(time.Time).Format(format)
					} else {
						fieldValue = v[name]
					}

					fields[name] = fieldValue
				}
			}
		}

		result = append(result, fields)
	}

	// 回调处理列表字段值
	return ctx.Template.(interface {
		BeforeIndexShowing(ctx *builder.Context, result []map[string]interface{}) []interface{}
	}).BeforeIndexShowing(ctx, result)
}
