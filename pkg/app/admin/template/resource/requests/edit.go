package requests

import (
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type EditRequest struct{}

// 表单数据
func (p *EditRequest) FillData(ctx *builder.Context) map[string]interface{} {
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

	// 查询数据
	model.Where("id = ?", id).First(&result)

	// 获取字段
	updateFields := template.UpdateFields(ctx)

	// 给实例的Field属性赋值
	template.SetField(result)

	// 解析字段值
	fields := make(map[string]interface{})
	for _, field := range updateFields.([]interface{}) {

		// 字段名
		name := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Name").
			String()

		if result[name] != nil {
			var fieldValue interface{}

			// 组件名称
			component := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Component").
				String()

			if component == "datetimeField" {
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
			} else if component == "dateField" {
				format := reflect.
					ValueOf(field).
					Elem().
					FieldByName("Format").
					String()

				format = strings.Replace(format, "YYYY", "2006", -1)
				format = strings.Replace(format, "MM", "01", -1)
				format = strings.Replace(format, "DD", "02", -1)

				fieldValue = result[name].(time.Time).Format(format)
			} else {
				fieldValue = result[name]
			}

			fields[name] = fieldValue
		}
	}

	return fields
}

// 获取表单初始化数据
func (p *EditRequest) Values(ctx *builder.Context) error {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取赋值数据
	data := p.FillData(ctx)

	// 获取初始数据
	data = template.BeforeEditing(ctx, data)

	return ctx.JSON(200, message.Success("获取成功", "", data))
}
