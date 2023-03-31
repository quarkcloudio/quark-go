package adminfrontpage

import (
	"github.com/quarkcms/quark-go/pkg/component/admin/page"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/admin/footer"
	"github.com/quarkcms/quark-go/pkg/component/admin/layout"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台登录模板
type Template struct {
	template.AdminTemplate
}

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 注册路由映射
	p.GET("/api/admin/frontpage/:resource/:component", "Render") // 自定义模板路由

	return p
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {
	data := map[string]interface{}{
		"component": ctx.Param("component"),
	}

	admin := &model.Admin{}
	config := ctx.Engine.GetConfig()

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(config.AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, msg.Error(err.Error(), ""))
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return ctx.JSON(200, msg.Error(err.Error(), ""))
	}

	adminLayout := ctx.Engine.GetAdminLayout()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(adminLayout.Copyright).
		SetLinks(adminLayout.Links)

	layoutComponent := (&layout.Component{}).
		Init().
		SetTitle(adminLayout.Title).
		SetLogo(adminLayout.Logo).
		SetActions(adminLayout.Actions).
		SetLayout(adminLayout.Layout).
		SetSplitMenus(adminLayout.SplitMenus).
		SetContentWidth(adminLayout.ContentWidth).
		SetPrimaryColor(adminLayout.PrimaryColor).
		SetFixSiderbar(adminLayout.FixSiderbar).
		SetFixedHeader(adminLayout.FixedHeader).
		SetIconfontUrl(adminLayout.IconfontUrl).
		SetLocale(adminLayout.Locale).
		SetSiderWidth(adminLayout.SiderWidth).
		SetMenu(getMenus).
		SetBody(data).
		SetFooter(footer)

	component := (&page.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"height": "100vh",
		}).
		SetBody(layoutComponent).
		JsonSerialize()

	return ctx.JSON(200, component)
}
