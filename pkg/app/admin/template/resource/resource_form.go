package resource

import (
	"reflect"
	"strings"

	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/card"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/requests"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/types"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
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

	if _, ok := fields.([]interface{}); ok {
		component := reflect.
			ValueOf(fields.([]interface{})[0]).
			Elem().
			FieldByName("Component").
			String()

		if component == "tabPane" {
			return p.FormWithinTabs(ctx, title, extra, api, fields, actions, data)
		}
	}

	return p.FormWithinCard(ctx, title, extra, api, fields, actions, data)
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

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// Form实例
	formComponent := template.
		GetForm().
		SetStyle(map[string]interface{}{
			"padding": "24px",
		}).
		SetApi(api).
		SetActions(actions).
		SetBody(fields).
		SetInitialValues(data)

	// 返回数据
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

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 返回数据
	return template.
		GetForm().
		SetStyle(map[string]interface{}{
			"backgroundColor": "#fff",
			"paddingBottom":   "20px",
		}).
		SetApi(api).
		SetActions(actions).
		SetBody(tabsComponent).
		SetInitialValues(data)
}

// 表单页面显示前回调
func (p *Template) BeforeFormShowing(ctx *builder.Context) map[string]interface{} {
	return map[string]interface{}{}
}

// 表单执行
func (p *Template) FormHandle(ctx *builder.Context, query *gorm.DB, data map[string]interface{}) (err error) {
	return (&requests.StoreRequest{}).Handle(ctx)
}

// 保存数据前回调
func (p *Template) BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	return submitData, nil
}

// 导入数据后回调
func (p *Template) AfterImported(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) (err error) {
	return err
}

// 保存数据后回调
func (p *Template) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) (err error) {
	return err
}

// 保存数据后跳转回调
func (p *Template) AfterSavedRedirectTo(ctx *builder.Context, err error) error {

	// 返回错误信息
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功！", strings.Replace("/layout/index?api="+IndexPath, ":resource", ctx.Param("resource"), -1)))
}
