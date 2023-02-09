package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

type EditableRequest struct{}

// 执行行为
func (p *EditableRequest) Handle(ctx *builder.Context) interface{} {
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	data := ctx.AllQuerys()
	for k, v := range data {
		if v == "true" {
			v = 1
		}
		if v == "false" {
			v = 0
		}
		data[k] = v
	}

	err := model.Where("id = ?", data["id"]).Updates(data).Error
	if err != nil {
		msg.Error(err.Error(), "")
	}

	return msg.Success("操作成功", "", "")
}
