package layout

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/footer"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/layout"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	builder.Template
	Title        string            `json:"title,omitempty"`
	Logo         interface{}       `json:"logo,omitempty"`
	Loading      bool              `json:"loading,omitempty"`
	ContentStyle map[string]string `json:"contentStyle,omitempty"`
	Actions      interface{}       `json:"actions,omitempty"`
	Layout       string            `json:"layout,omitempty"`
	SplitMenus   bool              `json:"splitMenus,omitempty"`
	ContentWidth string            `json:"contentWidth,omitempty"`
	PrimaryColor string            `json:"primaryColor,omitempty"`
	FixedHeader  bool              `json:"fixedHeader,omitempty"`
	FixSiderbar  bool              `json:"fixSiderbar,omitempty"`
	IconfontUrl  string            `json:"iconfontUrl,omitempty"`
	Locale       string            `json:"locale,omitempty"`
	SiderWidth   int               `json:"siderWidth,omitempty"`
	Menu         interface{}       `json:"menu,omitempty"`
	Footer       interface{}       `json:"footer,omitempty"`
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
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
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
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
