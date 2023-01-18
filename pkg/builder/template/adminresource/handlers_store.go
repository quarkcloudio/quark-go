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
func (p *StoreRequest) Handle(request *builder.Request, templateInstance interface{}) interface{} {
	modelInstance := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	// 获取字段
	fields := templateInstance.(interface {
		CreationFields(request *builder.Request, templateInstance interface{}) interface{}
	}).CreationFields(request, templateInstance)

	data := map[string]interface{}{}
	json.Unmarshal(request.Body(), &data)
	data, err := templateInstance.(interface {
		BeforeSaving(request *builder.Request, data map[string]interface{}) (map[string]interface{}, error)
	}).BeforeSaving(request, data)
	if err != nil {
		msg.Error(err.Error(), "")
	}

	validator := templateInstance.(interface {
		ValidatorForCreation(request *builder.Request, templateInstance interface{}, data map[string]interface{}) error
	}).ValidatorForCreation(request, templateInstance, data)
	if validator != nil {
		msg.Error(validator.Error(), "")
	}

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
			var reflectValue reflect.Value
			switch reflectFieldName.Type().String() {
			case "int":
				if value, ok := formValue.(bool); ok {
					if value == true {
						reflectValue = reflect.ValueOf(1)
					} else {
						reflectValue = reflect.ValueOf(0)
					}
				}
				if value, ok := formValue.(float64); ok {
					reflectValue = reflect.ValueOf(int(value))
				}
			case "float64":
				if value, ok := formValue.(float64); ok {
					reflectValue = reflect.ValueOf(float64(value))
				}
			case "float32":
				if value, ok := formValue.(float64); ok {
					reflectValue = reflect.ValueOf(float32(value))
				}
			case "time.Time":
				getTime, _ := time.ParseInLocation("2006-01-02 15:04:05", formValue.(string), time.Local)
				reflectValue = reflect.ValueOf(getTime)
			default:
				reflectValue = reflect.ValueOf(formValue)
				if reflect.ValueOf(formValue).Type().String() == "[]uint8" {
					reflectValue = reflect.ValueOf(string(formValue.([]uint8)))
				}
			}

			if reflectFieldName.Type().String() != reflectValue.Type().String() {
				return msg.Error("结构体类型与传参类型不一致！", "")
			}

			reflectFieldName.Set(reflectValue)
		}
	}

	// 获取对象
	getModel := model.Create(modelInstance)

	return templateInstance.(interface {
		AfterSaved(request *builder.Request, model *gorm.DB) interface{}
	}).AfterSaved(request, getModel)
}
