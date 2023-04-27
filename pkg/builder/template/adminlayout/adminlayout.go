package adminlayout

import (
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/admin/footer"
	"github.com/quarkcms/quark-go/pkg/component/admin/layout"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	template.Template
}

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 注册路由映射
	p.GET("/api/admin/layout/:resource/index", p.Render) // 获取布局配置

	return p
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	adminLayout := ctx.Engine.GetAdminLayout()

	admin := &model.Admin{}
	config := ctx.Engine.GetConfig()

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(config.AppKey, ctx.Token())
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(adminLayout.Copyright).
		SetLinks(adminLayout.Links)

	component := (&layout.Component{}).
		Init().
		SetTitle(adminLayout.Title).
		SetLogo(adminLayout.Logo).
		SetMenu(getMenus).
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
		SetFooter(footer)

	return ctx.JSON(200, component)
}
