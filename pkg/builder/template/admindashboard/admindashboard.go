package admindashboard

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/admin/card"
	"github.com/quarkcms/quark-go/pkg/component/admin/descriptions"
	"github.com/quarkcms/quark-go/pkg/component/admin/grid"
	"github.com/quarkcms/quark-go/pkg/component/admin/statistic"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台登录模板
type Template struct {
	template.AdminTemplate
	Title    string // 标题
	SubTitle string // 子标题
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
	p.GET("/api/admin/dashboard/:resource/index", "Render") // 后台仪表盘路由

	// 标题
	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Template) Cards(ctx *builder.Context) interface{} {
	return nil
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {
	cards := ctx.Template.(interface {
		Cards(*builder.Context) interface{}
	}).Cards(ctx)
	if cards == nil {
		return ctx.JSON(200, msg.Error("请实现Cards内容", ""))
	}

	var cols []interface{}
	var body []interface{}
	var colNum int = 0
	for key, v := range cards.([]interface{}) {

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
			FieldByName("Col").Int()
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

	component := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, component)
}
