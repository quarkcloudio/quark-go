package requests

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExportRequest struct{}

// 执行行为
func (p *ExportRequest) Handle(ctx *builder.Context) error {
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
		for _, field := range fields.([]interface{}) {

			name := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Name").String()

			component := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Component").String()

			switch component {
			case "inputNumberField":
				rowData[name] = dataValue.(map[string]interface{})[name]
			case "textField":
				rowData[name] = dataValue.(map[string]interface{})[name]
			case "selectField":
				optionLabel := field.(interface {
					GetOptionLabel(interface{}) string
				}).GetOptionLabel(dataValue.(map[string]interface{})[name])

				rowData[name] = optionLabel
			case "checkboxField":
				optionLabel := field.(interface {
					GetOptionLabel(interface{}) string
				}).GetOptionLabel(dataValue.(map[string]interface{})[name])

				rowData[name] = optionLabel
			case "radioField":
				optionLabel := field.(interface {
					GetOptionLabel(interface{}) string
				}).GetOptionLabel(dataValue.(map[string]interface{})[name])

				rowData[name] = optionLabel
			case "switchField":
				optionLabel := field.(interface {
					GetOptionLabel(interface{}) interface{}
				}).GetOptionLabel(dataValue.(map[string]interface{})[name])

				rowData[name] = optionLabel
			default:
				rowData[name] = dataValue.(map[string]interface{})[name]
			}

			f.SetCellValue("Sheet1", string(a)+strconv.Itoa(dataKey+2), rowData[name])
			a++
		}
	}

	f.SetActiveSheet(index)
	buf, _ := f.WriteToBuffer()

	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.Write(buf.Bytes())

	return nil
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
					var fieldValue interface{}

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
		BeforeExporting(ctx *builder.Context, result []map[string]interface{}) []interface{}
	}).BeforeExporting(ctx, result)
}
