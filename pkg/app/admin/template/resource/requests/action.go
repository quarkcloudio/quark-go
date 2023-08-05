package requests

import (
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

type ActionRequest struct{}

// 执行行为
func (p *ActionRequest) Handle(ctx *builder.Context) error {
	var result error

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// Gorm对象
	model := db.Client.Model(modelInstance)

	id := ctx.Query("id", "")
	if id != "" {
		if strings.Contains(id.(string), ",") {
			model.Where("id IN ?", strings.Split(id.(string), ","))
		} else {
			model.Where("id = ?", id)
		}
	}

	actions := template.Actions(ctx)
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
						Handle(*builder.Context, *gorm.DB) error
					}).Handle(ctx, model)

					// 执行完后回调
					err := template.AfterAction(ctx, uriKey, model)
					if err != nil {
						return err
					}

					return result
				}
			}
		} else {
			if ctx.Param("uriKey") == uriKey {
				result = v.(interface {
					Handle(*builder.Context, *gorm.DB) error
				}).Handle(ctx, model)

				// 执行完后回调
				err := template.AfterAction(ctx, uriKey, model)
				if err != nil {
					return err
				}

				return result
			}
		}
	}

	return result
}

// 行为表单值
func (p *ActionRequest) Values(ctx *builder.Context) error {
	var data map[string]interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 解析行为
	actions := template.Actions(ctx)
	for _, v := range actions {
		// uri唯一标识
		uriKey := v.(interface{ GetUriKey(interface{}) string }).GetUriKey(v)

		actionType := v.(interface{ GetActionType() string }).GetActionType()
		if actionType == "dropdown" {
			dropdownActions := v.(interface{ GetActions() []interface{} }).GetActions()
			for _, dropdownAction := range dropdownActions {

				// uri唯一标识
				uriKey := dropdownAction.(interface{ GetUriKey(interface{}) string }).GetUriKey(dropdownAction)
				if ctx.Param("uriKey") == uriKey {
					data = dropdownAction.(interface {
						Data(*builder.Context) map[string]interface{}
					}).Data(ctx)
				}
			}
		} else {
			if ctx.Param("uriKey") == uriKey {
				data = v.(interface {
					Data(*builder.Context) map[string]interface{}
				}).Data(ctx)
			}
		}
	}

	return ctx.JSON(200, message.Success("获取成功", "", data))
}
