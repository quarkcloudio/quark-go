package requests

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
)

type EditableRequest struct{}

// 执行行为
func (p *EditableRequest) Handle(ctx *quark.Context) error {
	var (
		id    interface{}
		field string
		value interface{}
	)

	// 获取所有Query数据
	data := ctx.AllQuerys()
	if data == nil {
		return ctx.JSONError("参数错误")
	}

	id = data["id"]
	if id == nil {
		return ctx.JSONError("id不能为空")
	}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取模型结构体
	modelInstance := template.GetModel()

	// 创建Gorm对象
	model := db.Client.Model(&modelInstance)

	// 解析数据
	for k, v := range data {
		switch v {
		case "true":
			v = 1
		case "false":
			v = 0
		}

		if k != "id" && k != "_t" {
			field = k
			value = v
		}
	}

	if field == "" || value == nil {
		return ctx.JSONError("参数错误")
	}

	// 表格行内编辑执行完之前回调
	err := template.BeforeEditable(ctx, id, field, value)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 创建表格行内编辑查询
	query := template.BuildEditableQuery(ctx, model)

	// 更新数据
	err = query.Update(field, value).Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 行为执行后回调
	err = template.AfterEditable(ctx, id, field, value)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("操作成功")
}
