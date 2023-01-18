package adminresource

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExportRequest struct{}

// 执行行为
func (p *ExportRequest) Handle(request *builder.Request, templateInstance interface{}) interface{} {
	data := p.QueryData(request, templateInstance)

	// 获取列表字段
	fields := templateInstance.(interface {
		ExportFields(request *builder.Request, templateInstance interface{}) interface{}
	}).ExportFields(request, templateInstance)

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	rowData := map[string]interface{}{}

	var a = 'a'
	for _, fieldValue := range fields.([]interface{}) {
		Label := reflect.
			ValueOf(fieldValue).
			Elem().
			FieldByName("Label").
			String()
		f.SetCellValue("Sheet1", string(a)+"1", Label)
		a++
	}

	for dataKey, dataValue := range data.([]interface{}) {
		var a = 'a'
		for _, fieldValue := range fields.([]interface{}) {

			name := reflect.
				ValueOf(fieldValue).
				Elem().
				FieldByName("Name").String()

			component := reflect.
				ValueOf(fieldValue).
				Elem().
				FieldByName("Component").String()

			switch component {
			case "inputNumberField":
				rowData[name] = dataValue.(map[string]interface{})[name]
			case "textField":
				rowData[name] = dataValue.(map[string]interface{})[name]
			case "selectField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()
				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])
			case "cascaderField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()
				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])
			case "checkboxField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()
				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])
			case "radioField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()
				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])
			case "switchField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()
				rowData[name] = p.getSwitchValue(options, dataValue.(map[string]interface{})[name].(int))
			default:
				rowData[name] = dataValue.(map[string]interface{})[name]
			}

			f.SetCellValue("Sheet1", string(a)+strconv.Itoa(dataKey+2), rowData[name])
			a++
		}
	}

	f.SetActiveSheet(index)
	buf, _ := f.WriteToBuffer()

	return buf.Bytes()
}

// 获取属性值
func (p *ExportRequest) getOptionValue(options interface{}, value interface{}) string {
	result := ""
	arr := []interface{}{}

	if value, ok := value.(string); ok {
		if strings.Contains(value, "[") || strings.Contains(value, "{") {
			json.Unmarshal([]byte(value), &arr)
		}
	}

	if len(arr) > 0 {
		if getOptions, ok := options.([]interface{}); ok {
			for _, option := range getOptions {
				for _, v := range arr {
					if v == option.(map[string]interface{})["value"] {
						result = result + option.(map[string]interface{})["label"].(string)
					}
				}
			}
		}
		if getOptions, ok := options.([]map[string]interface{}); ok {
			for _, option := range getOptions {
				for _, v := range arr {
					if v == option["value"] {
						result = result + option["label"].(string)
					}
				}
			}
		}
	} else {
		if getOptions, ok := options.([]interface{}); ok {
			for _, option := range getOptions {
				if value == option.(map[string]interface{})["value"] {
					result = option.(map[string]interface{})["label"].(string)
				}
			}
		}
		if getOptions, ok := options.([]map[string]interface{}); ok {
			for _, option := range getOptions {
				if value == option["value"] {
					result = option["label"].(string)
				}
			}
		}
	}

	return result
}

// 获取开关组件值
func (p *ExportRequest) getSwitchValue(options interface{}, value int) string {
	if value == 1 {
		return options.(map[string]interface{})["on"].(string)
	} else {
		return options.(map[string]interface{})["off"].(string)
	}
}

// 列表查询
func (p *ExportRequest) QueryData(request *builder.Request, templateInstance interface{}) interface{} {
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
		BuildExportQuery(request *builder.Request, templateInstance interface{}, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB
	}).BuildExportQuery(request, templateInstance, model, searches, filters, p.columnFilters(request), p.orderings(request))

	query.Find(&lists)

	// 返回解析列表
	return p.performsList(request, templateInstance, lists)
}

/**
 * Get the column filters for the request.
 *
 * @return array
 */
func (p *ExportRequest) columnFilters(request *builder.Request) map[string]interface{} {
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
func (p *ExportRequest) orderings(request *builder.Request) map[string]interface{} {
	data := request.AllQuerys()
	result, ok := data["sorter"].(map[string]interface{})
	if ok == false {
		return map[string]interface{}{}
	}

	return result
}

// 处理列表
func (p *ExportRequest) performsList(request *builder.Request, templateInstance interface{}, lists []map[string]interface{}) []interface{} {
	result := []map[string]interface{}{}

	// 获取列表字段
	exportFields := templateInstance.(interface {
		ExportFields(request *builder.Request, templateInstance interface{}) interface{}
	}).ExportFields(request, templateInstance)

	// 解析字段回调函数
	for _, v := range lists {

		// 给实例的Field属性赋值
		templateInstance.(interface {
			SetField(fieldData map[string]interface{}) interface{}
		}).SetField(v)

		fields := make(map[string]interface{})
		for _, field := range exportFields.([]interface{}) {

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
		BeforeExporting(request *builder.Request, result []map[string]interface{}) []interface{}
	}).BeforeExporting(request, result)
}
