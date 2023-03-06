package mixform

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/mixpage"
	"github.com/quarkcms/quark-go/pkg/component/mix/form"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台登录模板
type Template struct {
	mixpage.Template
	FromStyle string
	Api       string
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
	p.GET("/api/mix/form/:resource/index", "Render")  // 渲染页面路由
	p.Any("/api/mix/form/:resource/submit", "Handle") // 表单提交路由

	// 标题
	p.Title = "QuarkGo"

	return p
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

// 字段
func (p *Template) Fields(ctx *builder.Context) []interface{} {
	return nil
}

// 行为
func (p *Template) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		p.Action("提交", "primary").SetOpenType("submit"),
	}
}

// 表单数据
func (p *Template) Data(ctx *builder.Context) map[string]interface{} {
	return nil
}

// 内容
func (p *Template) Content(ctx *builder.Context) interface{} {

	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)

	data := ctx.Template.(interface {
		Data(ctx *builder.Context) map[string]interface{}
	}).Data(ctx)

	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	// 获取接口地址
	api := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Api").String()

	// 样式
	fromStyle := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("FromStyle").String()

	if api == "" {
		api = "/api/mix/form/" + strings.ToLower(ctx.ResourceName()) + "/submit"
	}

	return p.Form(api, fields).
		SetStyle(fromStyle).
		SetModel(data).
		SetActions(actions)
}

// 执行表单
func (p *Template) Handle(ctx *builder.Context) interface{} {
	return ctx.JSON(200, msg.Error("请自行处理表单逻辑", ""))
}
