package requests

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type EditableRequest struct{}

// 执行行为
func (p *EditableRequest) Handle(ctx *builder.Context) interface{} {
	var (
		id    interface{}
		field string
		value interface{}
	)

	// 获取模型结构体
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").
		Interface()

	// 创建Gorm对象
	model := db.Client.Model(&modelInstance)

	// 获取所有Query数据
	data := ctx.AllQuerys()
	if data == nil {
		return ctx.SimpleError("参数错误！")
	}

	id = data["id"]
	if id == nil {
		return ctx.SimpleError("id不能为空！")
	}

	// 解析数据
	for k, v := range data {
		if v == "true" {
			v = 1
		} else if v == "false" {
			v = 0
		}

		if k != "id" {
			field = k
			value = v
		}
	}

	if field == "" {
		return ctx.SimpleError("参数错误！")
	}

	if value == nil {
		return ctx.SimpleError("参数错误！")
	}

	// 更新数据
	err := model.Where("id = ?", id).Update(field, value).Error
	if err != nil {
		return ctx.SimpleError(err.Error())
	}

	result := ctx.Template.(interface {
		AfterEditable(ctx *builder.Context, id interface{}, field string, value interface{}) interface{}
	}).AfterEditable(ctx, id, field, value)
	if result != nil {
		return result
	}

	return ctx.SimpleSuccess("操作成功")
}
