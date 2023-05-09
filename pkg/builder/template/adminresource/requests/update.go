package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type UpdateRequest struct{}

// 执行行为
func (p *UpdateRequest) Handle(ctx *builder.Context) error {
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").
		Interface()

	data := map[string]interface{}{}
	ctx.Bind(&data)
	if data["id"] == "" {
		return ctx.JSONError("参数错误！")
	}

	data, err := ctx.Template.(interface {
		BeforeSaving(ctx *builder.Context, data map[string]interface{}) (map[string]interface{}, error)
	}).BeforeSaving(ctx, data)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	validator := ctx.Template.(interface {
		ValidatorForUpdate(ctx *builder.Context, data map[string]interface{}) error
	}).ValidatorForUpdate(ctx, data)
	if validator != nil {
		return ctx.JSONError(validator.Error())
	}

	newData := map[string]interface{}{}
	for k, v := range data {
		nv := v

		// 将数组、map数据转换为字符串存储
		if gv, ok := v.([]interface{}); ok {
			nv, _ = json.Marshal(gv)
		}
		if gv, ok := v.([]map[string]interface{}); ok {
			nv, _ = json.Marshal(gv)
		}
		if gv, ok := v.(map[string]interface{}); ok {
			nv, _ = json.Marshal(gv)
		}

		camelCaseName := stringy.
			New(k).
			CamelCase("?", "")

		fieldIsValid := reflect.
			ValueOf(modelInstance).
			Elem().
			FieldByName(camelCaseName).
			IsValid()
		if fieldIsValid {
			newData[k] = nv
		}
	}

	// 获取对象
	model := db.
		Client.
		Model(&modelInstance).
		Where("id = ?", data["id"]).
		Updates(newData)

	return ctx.Template.(interface {
		AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error
	}).AfterSaved(ctx, int(data["id"].(float64)), data, model)
}
