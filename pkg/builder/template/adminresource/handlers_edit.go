package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

type EditRequest struct{}

// 表单数据
func (p *EditRequest) FillData(request *builder.Request, templateInstance interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	id := request.Query("id", "")
	if id == "" {
		return result
	}

	modelInstance := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)
	model.Where("id = ?", id).First(&result)

	// 获取列表字段
	updateFields := templateInstance.(interface {
		UpdateFields(request *builder.Request, templateInstance interface{}) interface{}
	}).UpdateFields(request, templateInstance)

	// 给实例的Field属性赋值
	templateInstance.(interface {
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
			fields[name] = result[name]
		}
	}

	return fields
}

// 获取表单初始化数据
func (p *EditRequest) Values(request *builder.Request, templateInstance interface{}) interface{} {
	data := p.FillData(request, templateInstance)

	// 断言BeforeEditing方法，获取初始数据
	data = templateInstance.(interface {
		BeforeEditing(*builder.Request, map[string]interface{}) map[string]interface{}
	}).BeforeEditing(request, data)

	return msg.Success("获取成功", "", data)
}
