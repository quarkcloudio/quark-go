package adminresource

import (
	"reflect"
	"strconv"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type IndexRequest struct{}

// 列表查询
func (p *IndexRequest) QueryData(request *builder.Request, templateInstance interface{}) interface{} {
	var lists []map[string]interface{}
	modelInstance := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	// 搜索项
	searches := templateInstance.(interface {
		Searches(request *builder.Request) []interface{}
	}).Searches(request)

	// 过滤项，预留
	filters := templateInstance.(interface {
		Filters(request *builder.Request) []interface{}
	}).Filters(request)

	query := templateInstance.(interface {
		BuildIndexQuery(request *builder.Request, templateInstance interface{}, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB
	}).BuildIndexQuery(request, templateInstance, model, searches, filters, p.columnFilters(request), p.orderings(request))

	// 获取分页
	perPage := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("PerPage").Interface()

	// 不分页，直接返回lists
	if reflect.TypeOf(perPage).String() != "int" {
		query.Find(&lists)

		// 返回解析列表
		return p.performsList(request, templateInstance, lists)
	}

	var total int64
	page := request.Query("page", "1")
	pageSize := request.Query("pageSize", "10")
	if pageSize != "" {
		perPage, _ = strconv.Atoi(pageSize.(string))
	}
	getPage, _ := strconv.Atoi(page.(string))

	// 获取总数量
	query.Count(&total)

	// 获取列表
	query.Limit(perPage.(int)).Offset((getPage - 1) * perPage.(int)).Find(&lists)

	// 解析列表
	result := p.performsList(request, templateInstance, lists)

	return map[string]interface{}{
		"currentPage": getPage,
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
func (p *IndexRequest) columnFilters(request *builder.Request) map[string]interface{} {
	data := request.AllQuerys()
	result, ok := data["filter"].(map[string]interface{})
	if ok == false {
		return map[string]interface{}{}
	}

	return result
}

/**
 * Get the orderings for the request.
 *
 * @return array
 */
func (p *IndexRequest) orderings(request *builder.Request) map[string]interface{} {
	data := request.AllQuerys()
	result, ok := data["sorter"].(map[string]interface{})
	if ok == false {
		return map[string]interface{}{}
	}

	return result
}

// 处理列表
func (p *IndexRequest) performsList(request *builder.Request, templateInstance interface{}, lists []map[string]interface{}) []interface{} {
	result := []map[string]interface{}{}

	// 获取列表字段
	indexFields := templateInstance.(interface {
		IndexFields(request *builder.Request, templateInstance interface{}) interface{}
	}).IndexFields(request, templateInstance)

	// 解析字段回调函数
	for _, v := range lists {

		// 给实例的Field属性赋值
		templateInstance.(interface {
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
					fields[name] = v[name]
				}
			}
		}

		result = append(result, fields)
	}

	// 回调处理列表字段值
	return templateInstance.(interface {
		BeforeIndexShowing(request *builder.Request, result []map[string]interface{}) []interface{}
	}).BeforeIndexShowing(request, result)
}
