package requests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type DetailRequest struct{}

// 表单数据
func (p *DetailRequest) FillData(ctx *builder.Context) map[string]interface{} {
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
	detailFields := ctx.Template.(interface {
		DetailFields(ctx *builder.Context) interface{}
	}).DetailFields(ctx)

	// 给实例的Field属性赋值
	ctx.Template.(interface {
		SetField(fieldData map[string]interface{}) interface{}
	}).SetField(result)

	fields := make(map[string]interface{})
	for _, field := range detailFields.([]interface{}) {

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
			if result[name] != nil {
				var fieldValue interface{}
				fieldValue = result[name]
				getV, ok := result[name].(string)
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

					fieldValue = result[name].(time.Time).Format(format)
				}

				fields[name] = fieldValue
			}
		}
	}

	return fields
}

// 获取表单初始化数据
func (p *DetailRequest) Values(ctx *builder.Context) error {
	data := p.FillData(ctx)

	// 获取初始数据
	data = ctx.Template.(interface {
		BeforeDetailShowing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeDetailShowing(ctx, data)

	return ctx.JSONOk("获取成功", "", data)
}
