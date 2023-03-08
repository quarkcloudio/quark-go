package actions

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/actions"
	"github.com/quarkcms/quark-go/pkg/component/admin/action"
	"github.com/quarkcms/quark-go/pkg/component/admin/form"
	"github.com/quarkcms/quark-go/pkg/component/admin/space"
	"github.com/quarkcms/quark-go/pkg/component/admin/tpl"
)

type Import struct {
	actions.Modal
}

// 初始化
func (p *Import) Init() *Import {
	// 初始化父结构
	p.ParentInit()

	// 文字
	p.Name = "导入数据"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *Import) GetBody(ctx *builder.Context) interface{} {
	api := "/api/admin/" + ctx.Param("resource") + "/import"
	getTpl := (&tpl.Component{}).
		Init().
		SetBody("模板文件: <a href='/api/admin/" + ctx.Param("resource") + "/import/template?token=" + ctx.Token() + "' target='_blank'>下载模板</a>").
		SetStyle(map[string]interface{}{
			"marginLeft": "50px",
		})

	fields := []interface{}{
		(&space.Component{}).
			Init().
			SetBody(getTpl).
			SetDirection("vertical").
			SetSize("middle").
			SetStyle(map[string]interface{}{
				"marginBottom": "20px",
			}),
		(&builder.AdminField{}).
			File("fileId", "导入文件").
			SetLimitNum(1).
			SetLimitType([]string{
				"application/vnd.ms-excel",
				"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			}).
			SetHelp("请上传xls格式的文件"),
	}

	return (&form.Component{}).
		Init().
		SetKey("importModalForm", false).
		SetApi(api).
		SetBody(fields).
		SetLabelCol(map[string]interface{}{
			"span": 6,
		}).
		SetWrapperCol(map[string]interface{}{
			"span": 18,
		})
}

// 弹窗行为
func (p *Import) GetActions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&action.Component{}).
			Init().
			SetLabel("取消").
			SetActionType("cancel"),

		(&action.Component{}).
			Init().
			SetLabel("提交").
			SetWithLoading(true).
			SetReload("table").
			SetActionType("submit").
			SetType("primary", false).
			SetSubmitForm("importModalForm"),
	}
}
