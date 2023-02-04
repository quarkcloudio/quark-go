package adminresource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
)

// 列表页搜索表单
func (p *Template) IndexSearches(request *builder.Request, templateInstance interface{}) interface{} {
	searches := templateInstance.(interface {
		Searches(*builder.Request) []interface{}
	}).Searches(request)
	search := (&table.Search{}).Init()

	withExport := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("WithExport").Bool()

	if withExport {
		search = search.SetShowExportButton(true).SetExportApi(strings.Replace(ExportRoute, ":resource", request.Param("resource"), -1))
	}

	for _, v := range searches {
		component := v.(interface{ GetComponent() string }).GetComponent() // 获取组件名称
		name := v.(interface{ GetName() string }).GetName()                // label 标签的文本
		column := v.(interface {
			GetColumn(search interface{}) string
		}).GetColumn(v) // 字段名，支持数组
		api := v.(interface{ GetApi() string }).GetApi() // 获取接口
		options := v.(interface {
			Options(request *builder.Request) map[interface{}]interface{}
		}).Options(request) // 获取属性
		load := v.(interface {
			Load(request *builder.Request) map[string]string
		}).Load(request) // 获取接口

		// 搜索栏表单项
		item := (&table.SearchItem{}).
			Init().
			SetName(column).
			SetLabel(name).
			SetApi(api)

		switch component {
		case "textField":
			item = item.Input(options)
		case "selectField":
			if load != nil {
				item.SetLoad(load["field"], load["api"])
			}
			item = item.Select(options)
		case "multipleSelectField":
			item = item.MultipleSelect(options)
		case "dateField":
			item = item.Date(options)
		case "datetimeField":
			item = item.Datetime(options)
		case "dateRangeField":
			item = item.DateRange(options)
		case "datetimeRangeField":
			item = item.DatetimeRange(options)
		case "cascaderField":
			item = item.Cascader(options)
		}

		search = search.SetItems(item)
	}

	return search
}
