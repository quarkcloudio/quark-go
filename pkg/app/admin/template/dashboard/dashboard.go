package dashboard

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/card"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/descriptions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/grid"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/pagecontainer"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/statistic"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

// 后台登录模板
type Template struct {
	builder.Template
	Title    string // 标题
	SubTitle string // 子标题
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	p.TemplateInit(ctx)

	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "仪表盘"

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET("/api/admin/dashboard/:resource/index", p.Render) // 后台仪表盘路由

	return p
}

// 获取页面标题
func (p *Template) GetTitle() string {
	return p.Title
}

// 获取页面子标题
func (p *Template) GetSubTitle() string {
	return p.SubTitle
}

// 内容
func (p *Template) Cards(ctx *builder.Context) []interface{} {
	return nil
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *builder.Context, body interface{}) interface{} {
	template := ctx.Template.(Dashboarder)

	// 页面容器组件渲染
	return template.PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{} {
	template := ctx.Template.(Dashboarder)

	title := template.GetTitle()
	subTitle := template.GetSubTitle()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	return (&pagecontainer.Component{}).
		Init().
		SetHeader(header).
		SetBody(body)
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	template := ctx.Template.(Dashboarder)

	cards := template.Cards(ctx)
	if cards == nil {
		return ctx.JSON(200, message.Error("请实现Cards内容"))
	}

	var cols []interface{}
	var body []interface{}
	var colNum int = 0
	for key, v := range cards {

		// 断言statistic组件类型
		statistic, ok := v.(interface{ Calculate() *statistic.Component })
		item := (&card.Component{}).Init()
		if ok {
			item = item.SetBody(statistic.Calculate())
		} else {
			// 断言descriptions组件类型
			descriptions, ok := v.(interface {
				Calculate() *descriptions.Component
			})
			if ok {
				item = item.SetBody(descriptions.Calculate())
			}
		}

		col := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Col").
			Int()
		colInfo := (&grid.Col{}).Init().SetSpan(int(col)).SetBody(item)
		cols = append(cols, colInfo)
		colNum = colNum + int(col)
		if colNum%24 == 0 {
			row := (&grid.Row{}).Init().SetGutter(8).SetBody(cols)
			if key != 1 {
				row = row.SetStyle(map[string]interface{}{"marginTop": "20px"})
			}
			body = append(body, row)
			cols = nil
		}
	}

	if cols != nil {
		row := (&grid.Row{}).Init().SetGutter(8).SetBody(cols)
		if colNum > 24 {
			row = row.SetStyle(map[string]interface{}{"marginTop": "20px"})
		}
		body = append(body, row)
	}

	component := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, component)
}
