package requests

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type EditableRequest struct{}

// 执行行为
func (p *EditableRequest) Handle(ctx *builder.Context) error {
	var (
		id    interface{}
		field string
		value interface{}
	)

	// 获取所有Query数据
	data := ctx.AllQuerys()
	if data == nil {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	id = data["id"]
	if id == nil {
		return ctx.JSON(200, message.Error("id不能为空！"))
	}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取模型结构体
	modelInstance := template.GetModel()

	// 创建Gorm对象
	model := db.Client.Model(&modelInstance)

	// 解析数据
	for k, v := range data {
		if v == "true" {
			v = 1
		} else if v == "false" {
			v = 0
		}

		if k != "id" {
			field = k
			value = v
		}
	}

	if field == "" {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	if value == nil {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	// 更新数据
	err := model.Where("id = ?", id).Update(field, value).Error
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 行为执行后回调
	result := template.AfterEditable(ctx, id, field, value)
	if result != nil {
		return result
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
