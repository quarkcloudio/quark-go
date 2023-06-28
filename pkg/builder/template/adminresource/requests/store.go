package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gobeam/stringy"
	"github.com/gookit/goutil/structs"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type StoreRequest struct{}

// 执行行为
func (p *StoreRequest) Handle(ctx *builder.Context) error {
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").
		Interface()

	dataInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").
		Interface()

	data := map[string]interface{}{}
	ctx.Bind(&data)

	validator := ctx.Template.(interface {
		ValidatorForCreation(ctx *builder.Context, data map[string]interface{}) error
	}).ValidatorForCreation(ctx, data)
	if validator != nil {
		return ctx.JSONError(validator.Error())
	}

	data, err := ctx.Template.(interface {
		BeforeSaving(ctx *builder.Context, data map[string]interface{}) (map[string]interface{}, error)
	}).BeforeSaving(ctx, data)
	if err != nil {
		return ctx.JSONError(err.Error())
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

	structs.SetValues(dataInstance, newData)

	// 获取对象
	model := db.Client.Model(&modelInstance).Create(dataInstance)

	// 因为gorm使用结构体，不更新零值，需要使用map更新零值
	reflectId := reflect.
		ValueOf(dataInstance).
		Elem().
		FieldByName("Id")
	if !reflectId.IsValid() {
		return ctx.JSONError("参数错误!")
	}

	id := int(reflectId.Int())
	db.
		Client.
		Model(&modelInstance).
		Where("id = ?", id).
		Updates(newData)

	return ctx.Template.(interface {
		AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error
	}).AfterSaved(ctx, id, data, model)
}
