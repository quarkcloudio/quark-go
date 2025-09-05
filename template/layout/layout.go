package layout

import (
	"time"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/action"
	"github.com/quarkcloudio/quark-go/v4/component/footer"
	"github.com/quarkcloudio/quark-go/v4/component/layout"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
)

// 后台登录模板
type Template struct {
	quark.Template
	IndexPath    string                   // 路由路径
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
	RightMenus   []interface{}            // 右上角菜单
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/admin/layout/:resource/index" // 路由路径
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render) // 获取布局配置

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

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

	// 网站版权 time.Now().Format("2006") + " QuarkCloud"
	p.Copyright = time.Now().Format("2006") + " QuarkCloud"

	// 友情链接
	p.Links = []map[string]interface{}{
		{
			"key":   "1",
			"title": "QuarkGo",
			"href":  "https://github.com/quarkcloudio/quark-go",
		},
		{
			"key":   "2",
			"title": "QuarkCloud",
			"href":  "http://quarkcloud.io",
		},
		{
			"key":   "3",
			"title": "Github",
			"href":  "https://github.com/quarkcloudio",
		},
	}

	// 右上角菜单
	p.RightMenus = []interface{}{
		action.
			New().
			SetLabel("个人设置").
			SetActionType("link").
			SetType("link", false).
			SetIcon("setting").
			SetStyle(map[string]interface{}{
				"color": "rgb(0 0 0 / 88%)",
			}).
			SetHref("#/layout/index?api=/api/admin/account/form").
			SetSize("small"),

		action.
			New().
			SetLabel("退出登录").
			SetActionType("ajax").
			SetType("link", false).
			SetIcon("logout").
			SetStyle(map[string]interface{}{
				"color": "rgb(0 0 0 / 88%)",
			}).
			SetApi("/api/admin/logout/index/handle").
			SetSize("small"),
	}

	return p
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
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

// 右上角菜单
func (p *Template) GetRightMenus() []interface{} {
	return p.RightMenus
}

// 组件渲染
func (p *Template) Render(ctx *quark.Context) error {
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

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	copyright := template.GetCopyright()

	// 友情链接
	links := template.GetLinks()

	// 右上角菜单
	rightMenus := template.GetRightMenus()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(copyright).
		SetLinks(links)

	component := (&layout.Component{}).
		Init().
		SetTitle(title).
		SetLogo(logo).
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
		SetRightMenus(rightMenus).
		SetFooter(footer)

	return ctx.JSON(200, component)
}
