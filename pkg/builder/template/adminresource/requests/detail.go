package requests

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
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

		if result[name] != nil {
			fields[name] = result[name]
		}
	}

	return fields
}

// 获取表单初始化数据
func (p *DetailRequest) Values(ctx *builder.Context) interface{} {
	data := p.FillData(ctx)

	// 断言BeforeEditing方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeDetailShowing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeDetailShowing(ctx, data)

	return ctx.JSON(200, msg.Success("获取成功", "", data))
}
