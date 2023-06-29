package requests

import (
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

type EditRequest struct{}

// 表单数据
func (p *EditRequest) FillData(ctx *builder.Context) map[string]interface{} {
	result := map[string]interface{}{}
	id := ctx.Query("id", "")
	if id == "" {
		return result
	}

	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)
	model.Where("id = ?", id).First(&result)

	// 获取列表字段
	updateFields := ctx.Template.(interface {
		UpdateFields(ctx *builder.Context) interface{}
	}).UpdateFields(ctx)

	// 给实例的Field属性赋值
	ctx.Template.(interface {
		SetField(fieldData map[string]interface{}) interface{}
	}).SetField(result)

	fields := make(map[string]interface{})
	for _, field := range updateFields.([]interface{}) {

		// 字段名
		name := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Name").String()

		if result[name] != nil {
			var fieldValue interface{}
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
func (p *EditRequest) Values(ctx *builder.Context) interface{} {
	data := p.FillData(ctx)

	// 断言BeforeEditing方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeEditing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeEditing(ctx, data)

	return ctx.JSON(200, msg.Success("获取成功", "", data))
}
