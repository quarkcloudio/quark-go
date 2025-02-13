package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gobeam/stringy"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/types"
)

type UpdateRequest struct{}

// 执行行为
func (p *UpdateRequest) Handle(ctx *quark.Context) error {
	data := map[string]interface{}{}

	// 解析数据
	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	// 验证参数合法性
	if data["id"] == "" {
		return ctx.CJSONError("参数错误")
	}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// 验证数据合法性
	validator := template.ValidatorForUpdate(ctx, data)
	if validator != nil {
		return ctx.CJSONError(validator.Error())
	}

	// 保存前回调
	data, err = template.BeforeSaving(ctx, data)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	// 重组数据
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
			PascalCase("?", "").
			Get()

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
	model := db.Client.Model(modelInstance)

	// 创建更新查询
	query := template.BuildUpdateQuery(ctx, model)

	// 更新数据
	query = query.Updates(newData)

	// 保存后回调
	err = template.AfterSaved(ctx, int(data["id"].(float64)), data, query)

	return template.AfterSavedRedirectTo(ctx, int(data["id"].(float64)), data, err)
}
