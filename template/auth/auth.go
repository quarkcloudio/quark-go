package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"reflect"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/auth"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	redisclient "github.com/quarkcloudio/quark-go/v4/dal/redis"
	"github.com/quarkcloudio/quark-go/v4/dto/response"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/redis/go-redis/v9"
)

// 后台登录模板
type Template struct {
	quark.Template
	IndexPath      string      // 登录页面路由
	CaptchaPath    string      // 登录验证码路由
	LoginPath      string      // 登录执行路由
	LogoutPath     string      // 退出执行路由
	UserInfoPath   string      // 获取用户信息路由
	UserRoutesPath string      // 获取用户路由
	LoginApi       string      // 登录接口
	UserInfoApi    string      // 获取用户信息接口
	UserRoutesApi  string      // 获取用户路由接口
	Redirect       string      // 登录后跳转地址
	Logo           interface{} // 登录页面Logo
	Title          string      // 标题
	Body           interface{} `json:"body,omitempty"` // 表单内容
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
	p.IndexPath = "/api/admin/auth/:resource/index"           // 登录组件路由
	p.CaptchaPath = "/api/admin/auth/:resource/captcha"       // 登录获取验证码ID路由
	p.LoginPath = "/api/admin/auth/:resource/login"           // 登录执行路由
	p.LogoutPath = "/api/admin/auth/:resource/logout"         // 退出执行路由
	p.UserInfoPath = "/api/admin/auth/:resource/userInfo"     // 获取用户信息路由
	p.UserRoutesPath = "/api/admin/auth/:resource/userRoutes" // 获取用户路由
	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.Render)          // 登录页面路由
	p.GET(p.CaptchaPath, p.Captcha)       // 登录验证码路由
	p.POST(p.LoginPath, p.Login)          // 登录执行路由
	p.GET(p.LogoutPath, p.Logout)         // 退出执行路由
	p.GET(p.UserInfoPath, p.UserInfo)     // 获取用户信息路由
	p.GET(p.UserRoutesPath, p.UserRoutes) // 获取用户路由

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 登录接口
	p.LoginApi = ctx.RouterPathToUrl(p.LoginPath)

	// 获取用户信息接口
	p.UserInfoApi = ctx.RouterPathToUrl(p.UserInfoPath)

	// 获取用户路由接口
	p.UserRoutesApi = ctx.RouterPathToUrl(p.UserRoutesPath)

	// 标题
	p.Title = "QuarkGo"

	// 跳转地址
	p.Redirect = "/engine?api=/api/admin/dashboard/index/index"

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
func (p *Template) GetLoginApi() string {
	return p.LoginApi
}

// 获取获取用户信息接口
func (p *Template) GetUserInfoApi() string {
	return p.UserInfoApi
}

// 获取获取用户路由接口
func (p *Template) GetUserRoutesApi() string {
	return p.UserRoutesApi
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

// 生成验证码
func (p *Template) Captcha(ctx *quark.Context) error {
	uuid := captcha.NewLen(4)
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, uuid, 110, 38)

	return ctx.JSONOk("请求成功", map[string]interface{}{
		"captchaEnabled": true,
		"img":            base64.StdEncoding.EncodeToString(writer.Bytes()),
		"uuid":           uuid,
	})
}

// 字段
func (p *Template) Fields(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

// 登录方法
func (p *Template) Login(ctx *quark.Context) error {
	return ctx.JSONError("请实现登录方法")
}

// 获取用户信息
func (p *Template) UserInfo(ctx *quark.Context) error {
	userInfo, err := service.NewAuthService(ctx).GetAdmin()
	if err != nil {
		return ctx.JSONError("获取用户信息失败")
	}

	return ctx.JSONOk("请求成功", userInfo)
}

// 获取用户路由
func (p *Template) UserRoutes(ctx *quark.Context) error {
	routeType := ctx.Query("type")

	authRoutes, err := service.NewAuthService(ctx).GetUserRoutes(routeType.(string))
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("请求成功", response.UserRoutesResp{
		Routes: authRoutes,
		Home:   "home",
	})
}

// 退出方法
func (p *Template) Logout(ctx *quark.Context) error {
	return ctx.JSONRedirectTo("退出成功", "/")
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
	loginApi := template.GetLoginApi()

	// 获取用户信息接口
	userInfoApi := template.GetUserInfoApi()

	// 获取用户路由接口
	UserRoutesApi := template.GetUserRoutesApi()

	// 登录后跳转地址
	redirect := template.GetRedirect()

	// Logo
	logo := template.GetLogo()

	// 标题
	title := template.GetTitle()

	// 包裹在组件内的字段
	fields := p.FieldsWithinComponents(ctx)

	// 组件
	component = auth.New().
		SetLoginApi(loginApi).
		SetUserInfoApi(userInfoApi).
		SetUserRoutesApi(UserRoutesApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetBody(fields)

	return ctx.JSONOk("ok", component)
}
