package adminresource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/card"
	"github.com/quarkcms/quark-go/pkg/component/admin/form"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

// 表单接口
func (p *Template) FormApi(request *builder.Request) string {
	return ""
}

// 表单标题
func (p *Template) FormTitle(request *builder.Request, templateInstance interface{}) string {
	value := reflect.ValueOf(templateInstance).Elem()
	title := value.FieldByName("Title").String()
	if request.IsCreating() {
		return "创建" + title
	} else {
		if request.IsEditing() {
			return "编辑" + title
		}
	}

	return title
}

// 渲染表单组件
func (p *Template) FormComponentRender(
	request *builder.Request,
	templateInstance interface{},
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
			return p.FormWithinTabs(request, templateInstance, title, extra, api, getFields, actions, data)
		} else {
			return p.FormWithinCard(request, templateInstance, title, extra, api, fields, actions, data)
		}
	} else {
		return p.FormWithinCard(request, templateInstance, title, extra, api, fields, actions, data)
	}
}

// 在卡片内的From组件
func (p *Template) FormWithinCard(
	request *builder.Request,
	templateInstance interface{},
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
	request *builder.Request,
	templateInstance interface{},
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
func (p *Template) BeforeSaving(request *builder.Request, submitData map[string]interface{}) (map[string]interface{}, error) {
	return submitData, nil
}

// 保存数据后回调
func (p *Template) AfterSaved(request *builder.Request, model *gorm.DB) interface{} {
	if model.Error != nil {
		return msg.Error(model.Error.Error(), "")
	}

	return msg.Success("操作成功！", strings.Replace("/index?api="+IndexRoute, ":resource", request.Param("resource"), -1), "")
}
