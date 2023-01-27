package adminresource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type ActionRequest struct{}

// 执行行为
func (p *ActionRequest) Handle(request *builder.Request, templateInstance interface{}) interface{} {
	var result interface{}
	modelInstance := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	id := request.Query("id", "")
	if id != "" {
		if strings.Contains(id.(string), ",") {
			model.Where("id IN ?", strings.Split(id.(string), ","))
		} else {
			model.Where("id = ?", id)
		}
	}

	actions := templateInstance.(interface {
		Actions(request *builder.Request) []interface{}
	}).Actions(request)

	for _, v := range actions {
		// uri唯一标识
		uriKey := v.(interface {
			GetUriKey(interface{}) string
		}).GetUriKey(v)

		actionType := v.(interface{ GetActionType() string }).GetActionType()
		if actionType == "dropdown" {
			dropdownActions := v.(interface{ GetActions() []interface{} }).GetActions()
			for _, dropdownAction := range dropdownActions {
				// uri唯一标识
				uriKey := dropdownAction.(interface {
					GetUriKey(interface{}) string
				}).GetUriKey(dropdownAction)

				if request.Param("uriKey") == uriKey {
					result = dropdownAction.(interface {
						Handle(*builder.Request, *gorm.DB) interface{}
					}).Handle(request, model)
				}
			}
		} else {
			if request.Param("uriKey") == uriKey {
				result = v.(interface {
					Handle(*builder.Request, *gorm.DB) interface{}
				}).Handle(request, model)
			}
		}
	}

	return result
}
