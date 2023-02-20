package adminresource

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

type StoreRequest struct{}

// 执行行为
func (p *StoreRequest) Handle(ctx *builder.Context) interface{} {
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	// 获取字段
	fields := ctx.Template.(interface {
		CreationFields(ctx *builder.Context) interface{}
	}).CreationFields(ctx)

	data := map[string]interface{}{}
	json.Unmarshal(ctx.Body(), &data)
	data, err := ctx.Template.(interface {
		BeforeSaving(ctx *builder.Context, data map[string]interface{}) (map[string]interface{}, error)
	}).BeforeSaving(ctx, data)
	if err != nil {
		return ctx.JSON(200, msg.Error(err.Error(), ""))
	}

	validator := ctx.Template.(interface {
		ValidatorForCreation(ctx *builder.Context, data map[string]interface{}) error
	}).ValidatorForCreation(ctx, data)
	if validator != nil {
		return ctx.JSON(200, msg.Error(validator.Error(), ""))
	}
	zeroValues := map[string]interface{}{}
	for _, v := range fields.([]interface{}) {
		name := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Name").String()
		formValue := data[name]
		if getValue, ok := formValue.([]interface{}); ok {
			formValue, _ = json.Marshal(getValue)
		}
		if getValue, ok := formValue.([]map[string]interface{}); ok {
			formValue, _ = json.Marshal(getValue)
		}
		if getValue, ok := formValue.(map[string]interface{}); ok {
			formValue, _ = json.Marshal(getValue)
		}
		if name != "" && formValue != nil {
			fieldName := stringy.New(name).CamelCase("?", "")
			reflectFieldName := reflect.
				ValueOf(modelInstance).
				Elem().
				FieldByName(fieldName)

			if reflectFieldName.IsValid() {
				var reflectValue reflect.Value
				switch reflectFieldName.Type().String() {
				case "int":
					if value, ok := formValue.(bool); ok {
						if value {
							reflectValue = reflect.ValueOf(1)
						} else {
							reflectValue = reflect.ValueOf(0)
						}
					}
					if value, ok := formValue.(float64); ok {
						reflectValue = reflect.ValueOf(int(value))
					}
					if reflectValue.IsZero() {
						zeroValues[fieldName] = 0
					}
				case "float64":
					if value, ok := formValue.(float64); ok {
						reflectValue = reflect.ValueOf(float64(value))
					}
					if reflectValue.IsZero() {
						zeroValues[fieldName] = 0
					}
				case "float32":
					if value, ok := formValue.(float64); ok {
						reflectValue = reflect.ValueOf(float32(value))
					}
					if reflectValue.IsZero() {
						zeroValues[fieldName] = 0
					}
				case "time.Time":
					getTime, _ := time.ParseInLocation("2006-01-02 15:04:05", formValue.(string), time.Local)
					reflectValue = reflect.ValueOf(getTime)
				default:
					reflectValue = reflect.ValueOf(formValue)
					if reflect.ValueOf(formValue).Type().String() == "[]uint8" {
						reflectValue = reflect.ValueOf(string(formValue.([]uint8)))
					}
					if reflectValue.IsZero() {
						zeroValues[fieldName] = nil
					}
				}

				if reflectValue.IsValid() {
					if reflectFieldName.Type().String() != reflectValue.Type().String() {
						return ctx.JSON(200, msg.Error("结构体类型与传参类型不一致！", ""))
					}

					reflectFieldName.Set(reflectValue)
				}
			}
		}
	}

	// 获取对象
	getModel := model.Create(modelInstance)

	// 因为gorm使用结构体，不更新零值，需要使用map更新零值
	if len(zeroValues) > 0 {
		reflectId := reflect.
			ValueOf(modelInstance).
			Elem().
			FieldByName("Id")
		if reflectId.IsValid() {
			db.Client.Model(&modelInstance).Where("id = ?", reflectId).Updates(zeroValues)
		}
	}

	return ctx.Template.(interface {
		AfterSaved(ctx *builder.Context, model *gorm.DB) interface{}
	}).AfterSaved(ctx, getModel)
}
