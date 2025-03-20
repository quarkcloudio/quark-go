package page

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/action"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/col"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/grid"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/image"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/navbar"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/page"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/row"
	"github.com/quarkcloudio/quark-go/v3/template/miniapp/component/swiper"
)

// MiniApp模板
type Template struct {
	quark.Template
	IndexPath string // 渲染页面路由
	Title     string
	Style     string
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/miniapp/page/:resource/index" // 渲染页面路由
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render) // 渲染页面路由

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "QuarkGo"

	return p
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
	return p
}

// 头部导航
func (p *Template) Navbar(ctx *quark.Context, navbar *navbar.Component) interface{} {
	return nil
}

// 轮播图
func (p *Template) Banners(ctx *quark.Context) []*image.Component {
	return nil
}

// 内容
func (p *Template) Content(ctx *quark.Context) interface{} {
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
func (p *Template) Render(ctx *quark.Context) error {
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
		Navbar(ctx *quark.Context, navbar *navbar.Component) interface{}
	}).Navbar(ctx, navbar.New())

	// 样式
	style := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Style").
		String()

	// 轮播图
	banners := ctx.Template.(interface {
		Banners(ctx *quark.Context) []*image.Component
	}).Banners(ctx)
	if len(banners) > 0 {
		swiperItems := []*swiper.Item{}
		for _, banner := range banners {
			swiperItems = append(swiperItems, p.SwiperItem(banner))
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
		Content(ctx *quark.Context) interface{}
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
