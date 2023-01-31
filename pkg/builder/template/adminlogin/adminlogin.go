package adminlogin

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/component/admin/login"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台登录模板
type Template struct {
	template.AdminTemplate
	Api      string      // 登录接口
	Redirect string      // 登录后跳转地址
	Logo     interface{} // 登录页面Logo
	Title    string      // 标题
	SubTitle string      // 子标题
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
	p.AddRoute("/api/admin/login/:resource/index", "Render")        // 渲染登录页面路由
	p.AddRoute("/api/admin/login/:resource/handle", "Handle")       // 后台登录执行路由
	p.AddRoute("/api/admin/login/:resource/captchaId", "CaptchaId") // 后台登录获取验证码ID路由
	p.AddRoute("/api/admin/login/:resource/captcha/:id", "Captcha") // 后台登录验证码路由
	p.AddRoute("/api/admin/logout/:resource/handle", "Logout")      // 后台退出执行路由

	// 标题
	p.Title = "QuarkGo"

	// 跳转地址
	p.Redirect = "/index?api=/api/admin/dashboard/index/index"

	// 子标题
	p.SubTitle = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	return p
}

// 验证码ID
func (p *Template) CaptchaId(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Error("请实现创建验证码ID方法", "")
}

// 生成验证码
func (p *Template) Captcha(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Error("请实现生成验证码方法", "")
}

// 登录方法
func (p *Template) Handle(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Error("请实现登录方法", "")
}

// 退出方法
func (p *Template) Logout(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Error("请实现退出方法", "")
}

// 组件渲染
func (p *Template) Render(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	// 默认登录接口
	defaultLoginApi := resource.RouteToResourceUrl("/api/admin/login/:resource/handle")

	// 登录接口
	loginApi := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Api").String()
	if loginApi != "" {
		defaultLoginApi = loginApi
	}

	// 登录后跳转地址
	redirect := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Redirect").String()

	// Logo
	logo := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Logo").Interface()

	// 标题
	title := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Title").String()

	// 子标题
	subTitle := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("SubTitle").String()

	// 获取验证码ID链接
	captchaIdUrl := resource.RouteToResourceUrl("/api/admin/login/:resource/captchaId")

	// 验证码链接
	captchaUrl := resource.RouteToResourceUrl("/api/admin/login/:resource/captcha/:id")

	component := (&login.Component{}).
		Init().
		SetApi(defaultLoginApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetSubTitle(subTitle).
		SetCaptchaIdUrl(captchaIdUrl).
		SetCaptchaUrl(captchaUrl).
		JsonSerialize()

	return component
}
