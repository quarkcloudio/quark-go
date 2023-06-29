package admindashboard

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/admin/card"
	"github.com/quarkcms/quark-go/pkg/component/admin/descriptions"
	"github.com/quarkcms/quark-go/pkg/component/admin/footer"
	"github.com/quarkcms/quark-go/pkg/component/admin/grid"
	"github.com/quarkcms/quark-go/pkg/component/admin/layout"
	"github.com/quarkcms/quark-go/pkg/component/admin/page"
	"github.com/quarkcms/quark-go/pkg/component/admin/pagecontainer"
	"github.com/quarkcms/quark-go/pkg/component/admin/statistic"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台登录模板
type Template struct {
	template.Template
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

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *builder.Context, body interface{}) interface{} {

	// Layout组件
	layoutComponent := ctx.Template.(interface {
		LayoutComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).LayoutComponentRender(ctx, body)

	return (&page.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"height": "100vh",
		}).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 页面布局组件渲染
func (p *Template) LayoutComponentRender(ctx *builder.Context, body interface{}) interface{} {
	admin := &model.Admin{}
	config := ctx.Engine.GetConfig()

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(config.AppKey, ctx.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	adminLayout := ctx.Engine.GetAdminLayout()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(adminLayout.Copyright).
		SetLinks(adminLayout.Links)

	// 页面容器组件渲染
	pageContainerComponent := ctx.Template.(interface {
		PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageContainerComponentRender(ctx, body)

	return (&layout.Component{}).
		Init().
		SetTitle(adminLayout.Title).
		SetLogo(adminLayout.Logo).
		SetActions(adminLayout.Actions).
		SetLayout(adminLayout.Layout).
		SetSplitMenus(adminLayout.SplitMenus).
		SetContentWidth(adminLayout.ContentWidth).
		SetPrimaryColor(adminLayout.PrimaryColor).
		SetFixSiderbar(adminLayout.FixSiderbar).
		SetFixedHeader(adminLayout.FixedHeader).
		SetIconfontUrl(adminLayout.IconfontUrl).
		SetLocale(adminLayout.Locale).
		SetSiderWidth(adminLayout.SiderWidth).
		SetMenu(getMenus).
		SetBody(pageContainerComponent).
		SetFooter(footer)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{} {
	value := reflect.ValueOf(ctx.Template).Elem()
	title := value.FieldByName("Title").String()
	subTitle := value.FieldByName("SubTitle").String()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	return (&pagecontainer.Component{}).Init().SetHeader(header).SetBody(body)
}
