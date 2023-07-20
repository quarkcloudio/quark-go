package layout

import (
	"time"

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
	Title        string                   // layout 的左上角 的 title
	Logo         interface{}              // layout 的左上角 的 logo
	Actions      interface{}              // layout 的头部行为
	Layout       string                   // layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	SplitMenus   bool                     // layout 的菜单模式为mix时，是否自动分割菜单
	ContentWidth string                   // layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	PrimaryColor string                   // 主题色,"#1890ff"
	FixedHeader  bool                     // 是否固定 header 到顶部
	FixSiderbar  bool                     // 是否固定导航
	IconfontUrl  string                   // 使用 IconFont 的图标配置
	Locale       string                   // 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	SiderWidth   int                      // 侧边菜单宽度
	Copyright    string                   // 网站版权 time.Now().Format("2006") + " QuarkGo"
	Links        []map[string]interface{} // 友情链接
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// layout 的左上角 的 title
	p.Title = "QuarkGo"

	// layout 的左上角 的 logo
	p.Logo = false

	// layout 的头部行为
	p.Actions = nil

	// layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	p.Layout = "mix"

	// layout 的菜单模式为mix时，是否自动分割菜单
	p.SplitMenus = false

	// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	p.ContentWidth = "Fluid"

	// 主题色,"#1890ff"
	p.PrimaryColor = "#1890ff"

	// 是否固定 header 到顶部
	p.FixedHeader = true

	// 是否固定导航
	p.FixSiderbar = true

	// 使用 IconFont 的图标配置
	p.IconfontUrl = "//at.alicdn.com/t/font_1615691_3pgkh5uyob.js"

	// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	p.Locale = "zh-CN"

	// 侧边菜单宽度
	p.SiderWidth = 208

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	p.Copyright = time.Now().Format("2006") + " QuarkGo"

	// 友情链接
	p.Links = []map[string]interface{}{
		{
			"key":   "1",
			"title": "Quark",
			"href":  "http://www.quarkcms.com/",
		},
		{
			"key":   "2",
			"title": "爱小圈",
			"href":  "http://www.ixiaoquan.com",
		},
		{
			"key":   "3",
			"title": "Github",
			"href":  "https://github.com/quarkcms",
		},
	}

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET("/api/admin/layout/:resource/index", p.Render) // 获取布局配置

	return p
}

// 获取 layout 的左上角 的 title
func (p *Template) GetTitle() string {
	return p.Title
}

// 获取 layout 的左上角 的 logo
func (p *Template) GetLogo() interface{} {
	return p.Logo
}

// 获取 layout 的头部行为
func (p *Template) GetActions() interface{} {
	return p.Actions
}

// 获取 layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
func (p *Template) GetLayout() string {
	return p.Layout
}

// 获取 layout 的菜单模式为mix时，是否自动分割菜单
func (p *Template) GetSplitMenus() bool {
	return p.SplitMenus
}

// 获取 layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
func (p *Template) GetContentWidth() string {
	return p.ContentWidth
}

// 获取主题色,"#1890ff"
func (p *Template) GetPrimaryColor() string {
	return p.PrimaryColor
}

// 获取是否固定 header 到顶部
func (p *Template) GetFixedHeader() bool {
	return p.FixedHeader
}

// 获取是否固定导航
func (p *Template) GetFixSiderbar() bool {
	return p.FixSiderbar
}

// 获取使用 IconFont 的图标配置
func (p *Template) GetIconfontUrl() string {
	return p.IconfontUrl
}

// 获取当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
func (p *Template) GetLocale() string {
	return p.Locale
}

// 侧边菜单宽度
func (p *Template) GetSiderWidth() int {
	return p.SiderWidth
}

// 网站版权 time.Now().Format("2006") + " QuarkGo"
func (p *Template) GetCopyright() string {
	return p.Copyright
}

// 友情链接
func (p *Template) GetLinks() []map[string]interface{} {
	return p.Links
}

// 获取当前登录用户菜单
func (p *Template) GetMenus(ctx *builder.Context) (list interface{}, err error) {
	config := ctx.Engine.GetConfig()
	admin := &model.Admin{}

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(config.AppKey, ctx.Token())
	if err != nil {
		return nil, err
	}

	// 获取管理员菜单
	return admin.GetMenuListById(adminInfo.Id)
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	template := ctx.Template.(Layouter)

	// 获取 layout 的左上角 的 title
	title := template.GetTitle()

	// 获取 layout 的左上角 的 logo
	logo := template.GetLogo()

	// 获取 layout 的头部行为
	actions := template.GetActions()

	// 获取 layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	layoutMode := template.GetLayout()

	// 获取 layout 的菜单模式为mix时，是否自动分割菜单
	splitMenus := template.GetSplitMenus()

	// 获取 layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	contentWidth := template.GetContentWidth()

	// 获取主题色,"#1890ff"
	primaryColor := template.GetPrimaryColor()

	// 获取是否固定导航
	fixSiderbar := template.GetFixSiderbar()

	// 获取是否固定 header 到顶部
	fixedHeader := template.GetFixedHeader()

	// 获取使用 IconFont 的图标配置
	iconfontUrl := template.GetIconfontUrl()

	// 获取当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	locale := template.GetLocale()

	// 侧边菜单宽度
	siderWidth := template.GetSiderWidth()

	// 获取管理员菜单
	getMenus, err := template.GetMenus(ctx)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	copyright := template.GetCopyright()

	// 友情链接
	links := template.GetLinks()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(copyright).
		SetLinks(links)

	component := (&layout.Component{}).
		Init().
		SetTitle(title).
		SetLogo(logo).
		SetMenu(getMenus).
		SetActions(actions).
		SetLayout(layoutMode).
		SetSplitMenus(splitMenus).
		SetContentWidth(contentWidth).
		SetPrimaryColor(primaryColor).
		SetFixSiderbar(fixSiderbar).
		SetFixedHeader(fixedHeader).
		SetIconfontUrl(iconfontUrl).
		SetLocale(locale).
		SetSiderWidth(siderWidth).
		SetFooter(footer)

	return ctx.JSON(200, component)
}
