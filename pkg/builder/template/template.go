package template

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/footer"
	"github.com/quarkcms/quark-go/pkg/component/admin/layout"
	"github.com/quarkcms/quark-go/pkg/component/admin/page"
	"github.com/quarkcms/quark-go/pkg/component/admin/pagecontainer"
	"github.com/quarkcms/quark-go/pkg/msg"
	"gorm.io/gorm"
)

// 模板
type AdminTemplate struct {
	DB     *gorm.DB    // DB对象
	Model  interface{} // DB模型结构体
	Routes []*builder.Route
}

// 获取路由
func (p *AdminTemplate) GetRoutes() []*builder.Route {
	return p.Routes
}

// 注册路由
func (p *AdminTemplate) AddRoute(path string, handlerName string) *AdminTemplate {
	getRoute := &builder.Route{
		Path:        path,
		HandlerName: handlerName,
	}
	p.Routes = append(p.Routes, getRoute)

	return p
}

// 页面组件渲染
func (p *AdminTemplate) PageComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{} {

	// Layout组件
	layoutComponent := templateInstance.(interface {
		LayoutComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{}
	}).LayoutComponentRender(request, templateInstance, body)

	return (&page.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"height": "100vh",
		}).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 页面布局组件渲染
func (p *AdminTemplate) LayoutComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{} {
	admin := &model.Admin{}

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(request.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	adminLayout := builder.GetAdminLayoutConfig()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(adminLayout.Copyright).
		SetLinks(adminLayout.Links)

	// 页面容器组件渲染
	pageContainerComponent := templateInstance.(interface {
		PageContainerComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{}
	}).PageContainerComponentRender(request, templateInstance, body)

	return (&layout.Component{}).
		Init().
		SetTitle(adminLayout.Title).
		SetLogo(adminLayout.Logo).
		SetHeaderActions(adminLayout.HeaderActions).
		SetLayout(adminLayout.Layout).
		SetSplitMenus(adminLayout.SplitMenus).
		SetHeaderTheme(adminLayout.HeaderTheme).
		SetContentWidth(adminLayout.ContentWidth).
		SetNavTheme(adminLayout.NavTheme).
		SetPrimaryColor(adminLayout.PrimaryColor).
		SetFixSiderbar(adminLayout.FixSiderbar).
		SetFixedHeader(adminLayout.FixedHeader).
		SetIconfontUrl(adminLayout.IconfontUrl).
		SetLocale(adminLayout.Locale).
		SetSiderWidth(adminLayout.SiderWidth).
		SetMenu(getMenus).
		SetBody(pageContainerComponent).
		SetFooter(footer)
}

// 页面容器组件渲染
func (p *AdminTemplate) PageContainerComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{} {
	value := reflect.ValueOf(templateInstance).Elem()
	title := value.FieldByName("Title").String()
	subTitle := value.FieldByName("SubTitle").String()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	return (&pagecontainer.Component{}).Init().SetHeader(header).SetBody(body)
}

// 默认组件渲染
func (p *AdminTemplate) Render(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Error("请实现组件渲染方法", "")
}
