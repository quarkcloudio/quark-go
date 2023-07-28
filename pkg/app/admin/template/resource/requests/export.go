package requests

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/xuri/excelize/v2"
)

type ExportRequest struct{}

// 执行行为
func (p *ExportRequest) Handle(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	data := p.QueryData(ctx)

	// 获取列表字段
	fields := template.ExportFields(ctx)

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
				FieldByName("Name").
				String()

			component := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Component").
				String()

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

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取模型结构体
	modelInstance := template.GetModel()

	// 创建Gorm对象
	model := db.Client.Model(modelInstance)

	// 搜索项
	searches := template.Searches(ctx)

	// 过滤项
	filters := template.Filters(ctx)

	// 创建查询对象
	query := template.BuildExportQuery(ctx, model, searches, filters, p.columnFilters(ctx), p.orderings(ctx))

	// 查询数据
	query.Find(&lists)

	// 返回解析数据
	return p.performsList(ctx, lists)
}

// Get the column filters for the request.
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

// Get the orderings for the request.
func (p *ExportRequest) orderings(ctx *builder.Context) map[string]interface{} {
	var data map[string]interface{}

	querys := ctx.AllQuerys()
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

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取字段
	exportFields := template.ExportFields(ctx)

	// 解析字段
	for _, v := range lists {

		// 给实例的Field属性赋值
		template.SetField(v)

		fields := make(map[string]interface{})
		for _, field := range exportFields.([]interface{}) {

			// 字段名
			name := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Name").
				String()

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
						FieldByName("Component").
						String()

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
					} else {
						fieldValue = v[name]
					}

					fields[name] = fieldValue
				}
			}
		}

		result = append(result, fields)
	}

	// 导出前回调
	return template.BeforeExporting(ctx, result)
}
