package resource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/card"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tabs"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

// 表单接口
func (p *Template) FormApi(ctx *builder.Context) string {
	return ""
}

// 表单标题
func (p *Template) FormTitle(ctx *builder.Context) string {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取标题
	title := template.GetTitle()

	// 解析标题
	if ctx.IsCreating() {
		return "创建" + title
	} else {
		if ctx.IsEditing() {
			return "编辑" + title
		}
	}

	return title
}

// 渲染表单组件
func (p *Template) FormComponentRender(
	ctx *builder.Context,
	title string,
	extra interface{},
	api string,
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	getFields, ok := fields.([]interface{})

	if ok {
		component := reflect.
			ValueOf(fields.([]interface{})[0]).
			Elem().
			FieldByName("Component").
			String()

		if component == "tabPane" {
			return p.FormWithinTabs(ctx, title, extra, api, getFields, actions, data)
		} else {
			return p.FormWithinCard(ctx, title, extra, api, fields, actions, data)
		}
	} else {
		return p.FormWithinCard(ctx, title, extra, api, fields, actions, data)
	}
}

// 在卡片内的From组件
func (p *Template) FormWithinCard(
	ctx *builder.Context,
	title string,
	extra interface{},
	api string,
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	formComponent := (&form.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"padding": "24px",
		}).
		SetApi(api).
		SetActions(actions).
		SetBody(fields).
		SetInitialValues(data)

	return (&card.Component{}).
		Init().
		SetTitle(title).
		SetHeaderBordered(true).
		SetExtra(extra).
		SetBody(formComponent)
}

// 在标签页内的From组件
func (p *Template) FormWithinTabs(
	ctx *builder.Context,
	title string,
	extra interface{},
	api string,
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	tabsComponent := (&tabs.Component{}).Init().SetTabPanes(fields).SetTabBarExtraContent(extra)

	return (&form.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"backgroundColor": "#fff",
			"paddingBottom":   "20px",
		}).
		SetApi(api).
		SetActions(actions).
		SetBody(tabsComponent).
		SetInitialValues(data)
}

// 保存数据前回调
func (p *Template) BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	return submitData, nil
}

// 保存数据后回调
func (p *Template) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error {

	// 导入操作直接返回
	if ctx.IsImport() {
		return result.Error
	}

	// 返回错误信息
	if result.Error != nil {
		return ctx.JSON(200, message.Error(result.Error.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功！", strings.Replace("/layout/index?api="+IndexPath, ":resource", ctx.Param("resource"), -1)))
}
