package dashboard

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/card"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/descriptions"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/grid"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/pagecontainer"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/statistic"
)

// 后台登录模板
type Template struct {
	quark.Template
	IndexPath string // 路由路径
	Title     string // 页面标题
	SubTitle  string // 页面子标题
	BackIcon  bool   // 页面是否携带返回Icon
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/admin/dashboard/:resource/index" // 路由路径
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render) // 后台仪表盘路由

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 标题
	p.Title = "仪表盘"

	// 页面是否携带返回Icon
	p.BackIcon = false

	return p
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
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

// 页面是否携带返回Icon
func (p *Template) GetBackIcon() bool {
	return p.BackIcon
}

// 内容
func (p *Template) Cards(ctx *quark.Context) []interface{} {
	return nil
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(Dashboarder)

	// 页面容器组件渲染
	return template.PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(Dashboarder)

	// 页面标题
	title := template.GetTitle()

	// 页面子标题
	subTitle := template.GetSubTitle()

	// 页面是否携带返回Icon
	backIcon := template.GetBackIcon()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	if !backIcon {
		header.SetBackIcon(false)
	}

	return (&pagecontainer.Component{}).
		Init().
		SetHeader(header).
		SetBody(body)
}

// 组件渲染
func (p *Template) Render(ctx *quark.Context) error {
	template := ctx.Template.(Dashboarder)

	cards := template.Cards(ctx)
	if cards == nil {
		return ctx.CJSONError("请实现Cards内容")
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
