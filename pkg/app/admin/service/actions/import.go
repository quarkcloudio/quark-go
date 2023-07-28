package actions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/space"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tpl"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type ImportAction struct {
	actions.Modal
}

// 导入数据
func Import() *ImportAction {
	return &ImportAction{}
}

// 初始化
func (p *ImportAction) Init(ctx *builder.Context) interface{} {

	// 文字
	p.Name = "导入数据"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *ImportAction) GetBody(ctx *builder.Context) interface{} {
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
		(&resource.Field{}).
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
func (p *ImportAction) GetActions(ctx *builder.Context) []interface{} {

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
