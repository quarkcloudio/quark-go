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

	// 清空路由
	p.Routes = nil

	// 注册路由
	p.AddRoute("/api/admin/dashboard/:resource/index", "Render") // 后台仪表盘路由

	// 标题
	p.Title = "仪表盘"

	return p
}

// 内容
func (p *Template) Cards(request *builder.Request) interface{} {
	return nil
}

// 组件渲染
func (p *Template) Render(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	cards := templateInstance.(interface {
		Cards(*builder.Request) interface{}
	}).Cards(request)
	if cards == nil {
		return msg.Error("请实现Cards内容", "")
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

	return templateInstance.(interface {
		PageComponentRender(request *builder.Request, templateInstance interface{}, body interface{}) interface{}
	}).PageComponentRender(request, templateInstance, body)
}
