package login

import (
	"bytes"
	"time"

	"github.com/dchest/captcha"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/login"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	redisclient "github.com/quarkcms/quark-go/v2/pkg/dal/redis"
)

// 后台登录模板
type Template struct {
	builder.Template
	Api      string      // 登录接口
	Redirect string      // 登录后跳转地址
	Logo     interface{} // 登录页面Logo
	Title    string      // 标题
	SubTitle string      // 子标题
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 登录接口
	p.Api = ctx.RouterPathToUrl("/api/admin/login/:resource/handle")

	// 标题
	p.Title = "QuarkGo"

	// 跳转地址
	p.Redirect = "/layout/index?api=/api/admin/dashboard/index/index"

	// 子标题
	p.SubTitle = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET("/api/admin/login/:resource/index", p.Render)        // 渲染登录页面路由
	p.POST("/api/admin/login/:resource/handle", p.Handle)      // 后台登录执行路由
	p.GET("/api/admin/login/:resource/captchaId", p.CaptchaId) // 后台登录获取验证码ID路由
	p.GET("/api/admin/login/:resource/captcha/:id", p.Captcha) // 后台登录验证码路由
	p.GET("/api/admin/logout/:resource/handle", p.Logout)      // 后台退出执行路由

	return p
}

// 获取Api
func (p *Template) GetApi() string {
	return p.Api
}

// 获取登录成功后跳转地址
func (p *Template) GetRedirect() string {
	return p.Redirect
}

// 获取登录页面Logo
func (p *Template) GetLogo() interface{} {
	return p.Logo
}

// 获取登录页面标题
func (p *Template) GetTitle() string {
	return p.Title
}

// 获取登录页面子标题
func (p *Template) GetSubTitle() string {
	return p.SubTitle
}

// 验证码存储驱动，redis | memory
func (p *Template) CaptchaStore(store string) {
	if store == "redis" {
		captcha.SetCustomStore(&CaptchaStore{
			RedisClient: redisclient.Client,
			Expiration:  time.Second * 1000,
		})
	}
}

// 验证码ID
func (p *Template) CaptchaId(ctx *builder.Context) error {

	return ctx.JSON(200, message.Success("操作成功", "", map[string]string{
		"captchaId": captcha.NewLen(4),
	}))
}

// 生成验证码
func (p *Template) Captcha(ctx *builder.Context) error {
	id := ctx.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)
	ctx.Write(writer.Bytes())

	return nil
}

// 登录方法
func (p *Template) Handle(ctx *builder.Context) error {
	return ctx.JSON(200, message.Error("请实现登录方法"))
}

// 退出方法
func (p *Template) Logout(ctx *builder.Context) error {
	return ctx.JSON(200, message.Error("退出成功"))
}

// 组件渲染
func (p *Template) Render(ctx *builder.Context) error {
	template := ctx.Template.(Loginer)

	// 登录接口
	loginApi := template.GetApi()

	// 登录后跳转地址
	redirect := template.GetRedirect()

	// Logo
	logo := template.GetLogo()

	// 标题
	title := template.GetTitle()

	// 子标题
	subTitle := template.GetSubTitle()

	// 获取验证码ID链接
	captchaIdUrl := ctx.RouterPathToUrl("/api/admin/login/:resource/captchaId")

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/login/:resource/captcha/:id")

	// 组件
	component := (&login.Component{}).
		Init().
		SetApi(loginApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetSubTitle(subTitle).
		SetCaptchaIdUrl(captchaIdUrl).
		SetCaptchaUrl(captchaUrl).
		JsonSerialize()

	return ctx.JSON(200, component)
}
