package requests

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type ActionRequest struct{}

// 执行行为
func (p *ActionRequest) Handle(ctx *builder.Context) interface{} {
	var result interface{}
	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	id := ctx.Query("id", "")
	if id != "" {
		if strings.Contains(id.(string), ",") {
			model.Where("id IN ?", strings.Split(id.(string), ","))
		} else {
			model.Where("id = ?", id)
		}
	}

	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

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

				if ctx.Param("uriKey") == uriKey {
					result = dropdownAction.(interface {
						Handle(*builder.Context, *gorm.DB) interface{}
					}).Handle(ctx, model)

					return result
				}
			}
		} else {
			if ctx.Param("uriKey") == uriKey {
				result = v.(interface {
					Handle(*builder.Context, *gorm.DB) interface{}
				}).Handle(ctx, model)

				return result
			}
		}
	}

	return result
}
