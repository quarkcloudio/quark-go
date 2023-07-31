package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type UpdateRequest struct{}

// 执行行为
func (p *UpdateRequest) Handle(ctx *builder.Context) error {
	data := map[string]interface{}{}
	ctx.Bind(&data)

	// 验证参数合法性
	if data["id"] == "" {
		return ctx.JSON(200, message.Error("参数错误"))
	}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// 验证数据合法性
	validator := template.ValidatorForUpdate(ctx, data)
	if validator != nil {
		return ctx.JSON(200, message.Error(validator.Error()))
	}

	// 保存前回调
	data, err := template.BeforeSaving(ctx, data)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
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
	model := db.Client.
		Model(modelInstance).
		Where("id = ?", data["id"]).
		Updates(newData)

	return template.AfterSaved(ctx, int(data["id"].(float64)), data, model)
}
