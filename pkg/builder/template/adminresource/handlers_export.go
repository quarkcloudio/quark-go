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
func (p *ExportRequest) Handle(ctx *builder.Context) interface{} {
	data := p.QueryData(ctx)

	// 获取列表字段
	fields := ctx.Template.(interface {
		ExportFields(ctx *builder.Context) interface{}
	}).ExportFields(ctx)

	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
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
func (p *ExportRequest) QueryData(ctx *builder.Context) interface{} {
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
		BuildExportQuery(ctx *builder.Context, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB
	}).BuildExportQuery(ctx, model, searches, filters, p.columnFilters(ctx), p.orderings(ctx))

	query.Find(&lists)

	// 返回解析列表
	return p.performsList(ctx, lists)
}

/**
 * Get the column filters for the request.
 *
 * @return array
 */
func (p *ExportRequest) columnFilters(ctx *builder.Context) map[string]interface{} {
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
func (p *ExportRequest) orderings(ctx *builder.Context) map[string]interface{} {
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
func (p *ExportRequest) performsList(ctx *builder.Context, lists []map[string]interface{}) []interface{} {
	result := []map[string]interface{}{}

	// 获取列表字段
	exportFields := ctx.Template.(interface {
		ExportFields(ctx *builder.Context) interface{}
	}).ExportFields(ctx)

	// 解析字段回调函数
	for _, v := range lists {

		// 给实例的Field属性赋值
		ctx.Template.(interface {
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
	return ctx.Template.(interface {
		BeforeExporting(ctx *builder.Context, result []map[string]interface{}) []interface{}
	}).BeforeExporting(ctx, result)
}
