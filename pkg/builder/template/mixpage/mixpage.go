package mixpage

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/mix/action"
	"github.com/quarkcms/quark-go/pkg/component/mix/form"
	"github.com/quarkcms/quark-go/pkg/component/mix/grid"
	"github.com/quarkcms/quark-go/pkg/component/mix/group"
	"github.com/quarkcms/quark-go/pkg/component/mix/icon"
	"github.com/quarkcms/quark-go/pkg/component/mix/image"
	"github.com/quarkcms/quark-go/pkg/component/mix/layout"
	"github.com/quarkcms/quark-go/pkg/component/mix/link"
	"github.com/quarkcms/quark-go/pkg/component/mix/list"
	"github.com/quarkcms/quark-go/pkg/component/mix/navbar"
	"github.com/quarkcms/quark-go/pkg/component/mix/navigator"
	"github.com/quarkcms/quark-go/pkg/component/mix/page"
	"github.com/quarkcms/quark-go/pkg/component/mix/searchbar"
	"github.com/quarkcms/quark-go/pkg/component/mix/section"
	"github.com/quarkcms/quark-go/pkg/component/mix/segmentedcontrol"
	"github.com/quarkcms/quark-go/pkg/component/mix/swiper"
	"github.com/quarkcms/quark-go/pkg/component/mix/video"
	"github.com/quarkcms/quark-go/pkg/component/mix/view"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	template.MixTemplate
	Title string
	Style string
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
	p.GET("/api/mix/page/:resource/index", "Render") // 渲染页面路由

	// 标题
	p.Title = "QuarkGo"

	return p
}

// 头部导航
func (p *Template) NavBar(ctx *builder.Context, navbar *navbar.Component) interface{} {
	return nil
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {
	return nil
}

// 底部导航
func (p *Template) TabBar(ctx *builder.Context) interface{} {
	return nil
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {

	// 标题
	title := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Title").String()

	// 导航
	navbarInstance := (&navbar.Component{}).Init()
	navBar := ctx.Template.(interface {
		NavBar(ctx *builder.Context, navbar *navbar.Component) interface{}
	}).NavBar(ctx, navbarInstance)

	// 底部菜单
	tabBar := ctx.Template.(interface {
		TabBar(ctx *builder.Context) interface{}
	}).TabBar(ctx)

	// 样式
	style := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Style").String()

	// 内容
	content := ctx.Template.(interface {
		Content(ctx *builder.Context) interface{}
	}).Content(ctx)

	// 组件
	component := (&page.Component{}).
		Init().
		SetTitle(title).
		SetNavBar(navBar).
		SetTabBar(tabBar).
		SetStyle(style).
		SetContent(content).
		JsonSerialize()

	return ctx.JSON(200, component)
}

// View
func (p *Template) View(body interface{}) *view.Component {
	return (&view.Component{}).Init().SetBody(body)
}

// 布局-行
func (p *Template) Row(cols []interface{}) *layout.Row {
	return (&layout.Row{}).Init().SetBody(cols)
}

// 布局-列
func (p *Template) Col(span int, body interface{}) *layout.Col {
	return (&layout.Col{}).Init().SetSpan(span).SetBody(body)
}

// 列表组件
func (p *Template) List(items []interface{}) *list.Component {
	return (&list.Component{}).Init().SetBody(items)
}

// 列表子组件
func (p *Template) ListItem(title string, url string) *list.ListItem {
	return (&list.ListItem{}).Init().SetTitle(title).SetTo(url)
}

// 标题栏组件
func (p *Template) Section(title string, body interface{}) *section.Component {
	return (&section.Component{}).Init().SetTitle(title).SetBody(body)
}

// 分组组件
func (p *Template) Group(title string, body interface{}) *group.Component {
	return (&group.Component{}).Init().SetTitle(title).SetBody(body)
}

// 宫格组件
func (p *Template) Grid(column int, body []interface{}) *grid.Component {
	return (&grid.Component{}).Init().SetColumn(column).SetBody(body)
}

// 宫格子组件
func (p *Template) GridItem(body interface{}) *grid.GridItem {
	return (&grid.GridItem{}).Init().SetBody(body)
}

// 页面跳转
func (p *Template) Navigator(content interface{}, url string) *navigator.Component {
	return (&navigator.Component{}).Init().SetBody(content).SetUrl(url)
}

// 图片
func (p *Template) Image(src string) *image.Component {
	return (&image.Component{}).
		Init().
		SetSrc(src)
}

// Icon
func (p *Template) Icon(iconType string) *icon.Component {
	return (&icon.Component{}).Init().SetType(iconType)
}

// 视频
func (p *Template) Video(src string) *video.Component {
	return (&video.Component{}).Init().SetSrc(src)
}

// Link
func (p *Template) Link(href string, body interface{}) *link.Component {
	return (&link.Component{}).Init().SetHref(href).SetBody(body)
}

// 分段器
func (p *Template) SegmentedControl(titles []interface{}, items []interface{}) *segmentedcontrol.Component {
	return (&segmentedcontrol.Component{}).Init().SetTitles(titles).SetItems(items)
}

// 轮播图
func (p *Template) Swiper(items []interface{}) *swiper.Component {
	return (&swiper.Component{}).
		Init().
		SetAutoplay(true).
		SetIndicatorDots(true).
		SetItems(items)
}

// 轮播图子组件
func (p *Template) SwiperItem(body interface{}) *swiper.SwiperItem {
	return (&swiper.SwiperItem{}).SetBody(body)
}

// 表单
func (p *Template) Form(api string, items []interface{}) *form.Component {
	return (&form.Component{}).
		Init().
		SetApi(api).
		SetBody(items)
}

// 表单项
func (p *Template) Field() *form.Field {
	return (&form.Field{})
}

// 表单项
func (p *Template) FormItem() *form.Field {
	return (&form.Field{})
}

// 行为：label按钮文字，actionType按钮的样式类型 primary | default | warn
func (p *Template) Action(label string, actionType string) *action.Component {
	return (&action.Component{}).
		Init().
		SetLabel(label).
		SetType(actionType)
}

// 搜索栏
func (p *Template) SearchBar() *searchbar.Component {
	return (&searchbar.Component{}).Init()
}
