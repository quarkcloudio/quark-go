package adminresource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/card"
	"github.com/quarkcms/quark-go/pkg/component/admin/form"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
	"gorm.io/gorm"
)

// 表单接口
func (p *Template) FormApi(ctx *builder.Context) string {
	return ""
}

// 表单标题
func (p *Template) FormTitle(ctx *builder.Context) string {
	value := reflect.ValueOf(ctx.Template).Elem()
	title := value.FieldByName("Title").String()
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
			FieldByName("Component").String()

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
func (p *Template) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) interface{} {
	if result.Error != nil {
		return ctx.JSONError(result.Error.Error())
	}

	return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+IndexPath, ":resource", ctx.Param("resource"), -1))
}
