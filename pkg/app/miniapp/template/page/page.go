package page

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/col"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/grid"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/image"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/navbar"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/page"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/row"
	"github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/swiper"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

// MiniApp模板
type Template struct {
	builder.Template
	Title string
	Style string
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "QuarkGo"

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET("/api/miniapp/page/:resource/index", p.Render) // 渲染页面路由

	return p
}

// 头部导航
func (p *Template) Navbar(ctx *builder.Context, navbar *navbar.Component) interface{} {
	return nil
}

// 轮播图
func (p *Template) Banners(ctx *builder.Context) []*image.Component {
	return nil
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {
	return nil
}

// 行为：label按钮文字；buttonType类型，可选值为 primary info warning danger success default
func (p *Template) Action(label string, buttonType string) *action.Component {
	return action.
		New().
		SetLabel(label).
		SetType(buttonType)
}

// 图片
func (p *Template) Image(src string) *image.Component {
	return image.
		New().
		SetSrc(src)
}

// 行
func (p *Template) Row(body []*col.Component) *row.Component {
	return row.
		New().
		SetBody(body)
}

// 列
func (p *Template) Col(span int, body interface{}) *col.Component {
	return col.
		New().
		SetSpan(span).
		SetBody(body)
}

// 宫格
func (p *Template) Grid(columnNum int, body []*grid.Item) *grid.Component {
	return grid.
		New().
		SetColumnNum(columnNum).
		SetBody(body)
}

// 宫格项
func (p *Template) GridItem(body interface{}) *grid.Item {
	return grid.
		NewItem().
		SetBody(body)
}

// 轮播
func (p *Template) Swiper(body []*swiper.Item) *swiper.Component {
	return swiper.
		New().
		SetBody(body)
}

// 轮播项
func (p *Template) SwiperItem(body interface{}) *swiper.Item {
	return swiper.
		NewItem().
		SetBody(body)
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	var (
		components []interface{}
	)

	// 标题
	title := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Title").
		String()

	// 导航
	navbar := ctx.Template.(interface {
		Navbar(ctx *builder.Context, navbar *navbar.Component) interface{}
	}).Navbar(ctx, navbar.New())

	// 样式
	style := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Style").
		String()

	// 轮播图
	banners := ctx.Template.(interface {
		Banners(ctx *builder.Context) []*image.Component
	}).Banners(ctx)
	if len(banners) > 0 {
		swiperItems := []*swiper.Item{}
		for _, banner := range banners {
			swiperItems = append(swiperItems, p.SwiperItem(banner.SetStyle("width:100%;height:200px;")))
		}
		components = append(components,
			p.
				Swiper(swiperItems).
				SetPaginationVisible(true).
				SetPaginationColor("#426543").
				SetAutoPlay(3000),
		)
	}

	// 内容
	content := ctx.Template.(interface {
		Content(ctx *builder.Context) interface{}
	}).Content(ctx)
	components = append(components, content)

	// 组件
	component := (&page.Component{}).
		Init().
		SetTitle(title).
		SetNavbar(navbar).
		SetStyle(style).
		SetContent(components).
		JsonSerialize()

	return ctx.JSON(200, component)
}
