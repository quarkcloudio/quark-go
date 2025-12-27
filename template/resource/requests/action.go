package requests

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
	"gorm.io/gorm"
)

type ActionRequest struct{}

// 获取行为
func (p *ActionRequest) GetActions(ctx *quark.Context) []interface{} {
	actions := []interface{}{}
	template := ctx.Template.(types.Resourcer)

	// 资源上的行为
	getActions := template.Actions(ctx)
	if getActions != nil {
		actions = getActions
	}

	// 字段上的行为
	fields := template.Fields(ctx)
	for _, v := range fields {
		component := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Component").
			String()
		if component == "actionField" {
			actionItems := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Items").
				Interface()
			for _, action := range actionItems.([]interface{}) {
				actions = append(actions, action)
			}
		}
	}

	return actions
}

// 执行行为
func (p *ActionRequest) Handle(ctx *quark.Context) error {
	var result error

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// Gorm对象
	model := db.Client.Model(modelInstance)

	// 查询条件
	model = template.BuildActionQuery(ctx, model)

	actions := p.GetActions(ctx)
	for _, v := range actions {
		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.New(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// uri唯一标识
		uriKey := actionInstance.GetUriKey(v)

		// 获取行为类型
		actionType := actionInstance.GetActionType()

		if actionType == "dropdown" {
			dropdownActioner := v.(types.Dropdowner)
			for _, dropdownAction := range dropdownActioner.GetActions() {
				uriKey := dropdownActioner.GetUriKey(dropdownAction)
				if ctx.Param("uriKey") == uriKey {
					// 执行前回调
					err := template.BeforeAction(ctx, uriKey, model)
					if err != nil {
						return ctx.JSONError(err.Error())
					}

					result = dropdownAction.(interface {
						Handle(*quark.Context, *gorm.DB) error
					}).Handle(ctx, model)

					// 执行完回调
					template.AfterAction(ctx, uriKey, model)

					return result
				}
			}
		} else {
			if ctx.Param("uriKey") == uriKey {
				// 执行前回调
				err := template.BeforeAction(ctx, uriKey, model)
				if err != nil {
					return ctx.JSONError(err.Error())
				}

				result = v.(interface {
					Handle(*quark.Context, *gorm.DB) error
				}).Handle(ctx, model)

				// 执行完回调
				template.AfterAction(ctx, uriKey, model)

				return result
			}
		}
	}

	return result
}

// 行为表单值
func (p *ActionRequest) Values(ctx *quark.Context) error {
	var data map[string]interface{}

	// 解析行为
	actions := p.GetActions(ctx)
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.New(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// uri唯一标识
		uriKey := actionInstance.GetUriKey(v)

		// 获取行为类型
		actionType := actionInstance.GetActionType()
		if actionType == "dropdown" {
			dropdownActioner := v.(types.Dropdowner)
			for _, dropdownAction := range dropdownActioner.GetActions() {
				uriKey := dropdownActioner.GetUriKey(dropdownAction)
				if ctx.Param("uriKey") == uriKey {
					data = dropdownAction.(interface {
						Data(*quark.Context) map[string]interface{}
					}).Data(ctx)
				}
			}
		} else {
			if ctx.Param("uriKey") == uriKey {
				data = v.(interface {
					Data(*quark.Context) map[string]interface{}
				}).Data(ctx)
			}
		}
	}

	return ctx.JSONOk("获取成功", data)
}
