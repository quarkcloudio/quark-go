package login

import (
	"bytes"
	"context"
	"reflect"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v8"
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

type Store struct {
	RedisClient *redis.Client
	Expiration  time.Duration
}

func (store *Store) Set(id string, digits []byte) {
	store.RedisClient.Set(context.Background(), id, string(digits), store.Expiration)
}

func (store *Store) Get(id string, clear bool) (digits []byte) {
	bytes, _ := store.RedisClient.Get(context.Background(), id).Bytes()
	return bytes
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
	p.GET("/api/admin/login/:resource/index", p.Render)        // 渲染登录页面路由
	p.POST("/api/admin/login/:resource/handle", p.Handle)      // 后台登录执行路由
	p.GET("/api/admin/login/:resource/captchaId", p.CaptchaId) // 后台登录获取验证码ID路由
	p.GET("/api/admin/login/:resource/captcha/:id", p.Captcha) // 后台登录验证码路由
	p.GET("/api/admin/logout/:resource/handle", p.Logout)      // 后台退出执行路由

	// 标题
	p.Title = "QuarkGo"

	// 跳转地址
	p.Redirect = "/layout/index?api=/api/admin/dashboard/index/index"

	// 子标题
	p.SubTitle = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	return p
}

// 验证码存储驱动，redis | memory
func (p *Template) CaptchaStore(store string) {

	if store == "redis" {
		captcha.SetCustomStore(&Store{
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

	// 模板实例
	templateInstance := ctx.Template
	if templateInstance == nil {
		return ctx.JSON(200, message.Error("模板实例获取失败"))
	}

	// 默认登录接口
	defaultLoginApi := ctx.RouterPathToUrl("/api/admin/login/:resource/handle")

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
	captchaIdUrl := ctx.RouterPathToUrl("/api/admin/login/:resource/captchaId")

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/login/:resource/captcha/:id")

	// 组件
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

	return ctx.JSON(200, component)
}
