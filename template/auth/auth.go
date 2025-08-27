package auth

import (
	"bytes"
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/divider"
	"github.com/quarkcloudio/quark-go/v4/component/login"
	"github.com/quarkcloudio/quark-go/v4/component/tabs"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	redisclient "github.com/quarkcloudio/quark-go/v4/dal/redis"
	"github.com/redis/go-redis/v9"
)

// 后台登录模板
type Template struct {
	quark.Template
	IndexPath   string      // 登录页面路由
	LoginPath   string      // 登录执行路由
	CaptchaPath string      // 登录验证码路由
	LogoutPath  string      // 退出执行路由
	Api         string      // 登录接口
	Redirect    string      // 登录后跳转地址
	Logo        interface{} // 登录页面Logo
	Title       string      // 标题
	Body        interface{} `json:"body,omitempty"` // 表单内容
}

type CaptchaStore struct {
	RedisClient *redis.Client
	Expiration  time.Duration
}

func (store *CaptchaStore) Set(id string, digits []byte) {
	store.RedisClient.Set(context.Background(), id, string(digits), store.Expiration)
}

func (store *CaptchaStore) Get(id string, clear bool) (digits []byte) {
	bytes, _ := store.RedisClient.Get(context.Background(), id).Bytes()
	return bytes
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/admin/auth/:resource/index"     // 登录页面路由
	p.LoginPath = "/api/admin/auth/:resource/login"     // 登录执行路由
	p.CaptchaPath = "/api/admin/auth/:resource/captcha" // 登录获取验证码ID路由
	p.LogoutPath = "/api/admin/auth/:resource/logout"   // 退出执行路由
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render)    // 登录页面路由
	p.POST(p.LoginPath, p.Login)    // 登录执行路由
	p.GET(p.CaptchaPath, p.Captcha) // 登录验证码路由
	p.GET(p.LogoutPath, p.Logout)   // 退出执行路由

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 登录接口
	p.Api = ctx.RouterPathToUrl("/api/admin/auth/:resource/login")

	// 标题
	p.Title = "QuarkGo"

	// 跳转地址
	p.Redirect = "/layout/index?api=/api/admin/dashboard/index/index"

	// 如果启动了redis缓存，验证码使用redis缓存
	if redisclient.Client != nil {
		captcha.SetCustomStore(&CaptchaStore{
			RedisClient: redisclient.Client,
			Expiration:  time.Second * 1000,
		})
	}

	return p
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
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

// 验证码ID
func (p *Template) CaptchaId(ctx *quark.Context) error {
	return ctx.CJSONOk("获取成功", map[string]string{
		"captchaId": captcha.NewLen(4),
	})
}

// 生成验证码
func (p *Template) Captcha(ctx *quark.Context) error {
	id := ctx.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)
	ctx.Write(writer.Bytes())

	return nil
}

// 字段
func (p *Template) Fields(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

// 登录方法
func (p *Template) Login(ctx *quark.Context) error {
	return ctx.CJSONError("请实现登录方法")
}

// 退出方法
func (p *Template) Logout(ctx *quark.Context) error {
	return ctx.CJSONRedirectTo("退出成功", "/")
}

// 包裹在组件内的创建页字段
func (p *Template) FieldsWithinComponents(ctx *quark.Context) interface{} {

	// 资源实例
	template := ctx.Template.(Auther)

	// 获取字段
	fields := template.Fields(ctx)

	// 解析创建页表单组件内的字段
	items := p.FormFieldsParser(ctx, fields)

	return items
}

// 解析创建页表单组件内的字段
func (p *Template) FormFieldsParser(ctx *quark.Context, fields interface{}) interface{} {
	items := []interface{}{}

	// 解析字段
	if fields, ok := fields.([]interface{}); ok {
		for _, v := range fields {
			hasBody := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Body").
				IsValid()
			if hasBody {

				// 获取内容值
				body := reflect.
					ValueOf(v).
					Elem().
					FieldByName("Body").
					Interface()

				// 解析值
				getFields := p.FormFieldsParser(ctx, body)

				// 更新值
				reflect.
					ValueOf(v).
					Elem().
					FieldByName("Body").
					Set(reflect.ValueOf(getFields))

				items = append(items, v)
			} else {
				component := reflect.
					ValueOf(v).
					Elem().
					FieldByName("Component").
					String()
				if strings.Contains(component, "Field") {

					// 判断是否在创建页面
					if v, ok := v.(interface{ IsShownOnCreation() bool }); ok {
						if v.IsShownOnCreation() {

							// 生成前端验证规则
							v.(interface{ BuildFrontendRules(string) interface{} }).BuildFrontendRules(ctx.Path())

							// 组合数据
							items = append(items, v)
						}
					}
				} else {
					items = append(items, v)
				}
			}
		}
	}

	return items
}

// 组件渲染
func (p *Template) Render(ctx *quark.Context) error {
	var component interface{}

	template := ctx.Template.(Auther)

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

	// 包裹在组件内的字段
	fields := p.FieldsWithinComponents(ctx)

	// 解析tabPane组件
	if _, ok := fields.([]interface{}); ok {
		componentName := reflect.
			ValueOf(fields.([]interface{})[0]).
			Elem().
			FieldByName("Component").
			String()

		if componentName == "tabPane" {
			tabComponent := (&tabs.Component{}).
				Init().
				SetTabPanes(fields).
				SetCentered(true)

			// 组件
			component = (&login.Component{}).
				Init().
				SetApi(loginApi).
				SetRedirect(redirect).
				SetLogo(logo).
				SetTitle(title).
				SetSubTitle(subTitle).
				SetBody(tabComponent)
		} else {
			fields := append([]interface{}{divider.New().SetStyle(map[string]interface{}{"marginTop": "-15px"})}, fields.([]interface{})...)

			// 组件
			component = (&login.Component{}).
				Init().
				SetApi(loginApi).
				SetRedirect(redirect).
				SetLogo(logo).
				SetTitle(title).
				SetSubTitle(subTitle).
				SetBody(fields)
		}
	} else {
		fields := append([]interface{}{divider.New().SetStyle(map[string]interface{}{"marginTop": "-15px"})}, fields.([]interface{})...)

		// 组件
		component = (&login.Component{}).
			Init().
			SetApi(loginApi).
			SetRedirect(redirect).
			SetLogo(logo).
			SetTitle(title).
			SetSubTitle(subTitle).
			SetBody(fields)
	}

	return ctx.JSON(200, component)
}
