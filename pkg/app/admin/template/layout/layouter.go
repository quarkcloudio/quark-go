package layout

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Layouter interface {

	// 模版接口
	builder.Templater

	// 获取 layout 的左上角 的 title
	GetTitle() string

	// 获取 layout 的左上角 的 logo
	GetLogo() interface{}

	// 获取 layout 的头部行为
	GetActions() interface{}

	// 获取 layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	GetLayout() string

	// 获取 layout 的菜单模式为mix时，是否自动分割菜单
	GetSplitMenus() bool

	// 获取 layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	GetContentWidth() string

	// 获取主题色,"#1890ff"
	GetPrimaryColor() string

	// 获取是否固定 header 到顶部
	GetFixedHeader() bool

	// 获取是否固定导航
	GetFixSiderbar() bool

	// 获取使用 IconFont 的图标配置
	GetIconfontUrl() string

	// 获取当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	GetLocale() string

	// 侧边菜单宽度
	GetSiderWidth() int

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	GetCopyright() string

	// 友情链接
	GetLinks() []map[string]interface{}

	// 获取当前登录管理员菜单
	GetMenus(ctx *builder.Context) (list interface{}, err error)

	// 组件渲染
	Render(ctx *builder.Context) error
}
