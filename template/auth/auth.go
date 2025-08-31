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
	data := map[string]interface{}{
		"id":       "0",
		"username": "quark",
		"roles": []string{
			"R_SUPER",
		},
		"buttons": []string{
			"B_CODE1",
			"B_CODE2",
			"B_CODE3",
		},
	}

	return ctx.JSONOk("请求成功", data)
}

// 获取用户路由
func (p *Template) UserRoutes(ctx *quark.Context) error {
	data := map[string]interface{}{
		"routes": []map[string]interface{}{
			{
				"name":      "exception",
				"path":      "/exception",
				"component": "layout.base",
				"meta": map[string]interface{}{
					"title":   "exception",
					"i18nKey": "route.exception",
					"icon":    "ant-design:exception-outlined",
					"order":   7,
				},
				"children": []map[string]interface{}{
					{
						"name":      "exception_403",
						"path":      "/exception/403",
						"component": "view.403",
						"meta": map[string]interface{}{
							"title":   "exception_403",
							"i18nKey": "route.exception_403",
							"icon":    "ic:baseline-block",
						},
					},
					{
						"name":      "exception_404",
						"path":      "/exception/404",
						"component": "view.404",
						"meta": map[string]interface{}{
							"title":   "exception_404",
							"i18nKey": "route.exception_404",
							"icon":    "ic:baseline-web-asset-off",
						},
					},
					{
						"name":      "exception_500",
						"path":      "/exception/500",
						"component": "view.500",
						"meta": map[string]interface{}{
							"title":   "exception_500",
							"i18nKey": "route.exception_500",
							"icon":    "ic:baseline-wifi-off",
						},
					},
				},
			},
			{
				"name":      "about",
				"path":      "/about",
				"component": "layout.base$view.about",
				"meta": map[string]interface{}{
					"title":   "about",
					"i18nKey": "route.about",
					"icon":    "fluent:book-information-24-regular",
					"order":   10,
				},
			},
			{
				"name":      "function",
				"path":      "/function",
				"component": "layout.base",
				"meta": map[string]interface{}{
					"title":   "function",
					"i18nKey": "route.function",
					"icon":    "icon-park-outline:all-application",
					"order":   6,
				},
				"children": []map[string]interface{}{
					{
						"name":     "function_hide-child",
						"path":     "/function/hide-child",
						"redirect": "/function/hide-child/one",
						"meta": map[string]interface{}{
							"title":   "function_hide-child",
							"i18nKey": "route.function_hide-child",
							"icon":    "material-symbols:filter-list-off",
							"order":   2,
						},
						"children": []map[string]interface{}{
							{
								"name":      "function_hide-child_one",
								"path":      "/function/hide-child/one",
								"component": "view.function_hide-child_one",
								"meta": map[string]interface{}{
									"title":      "function_hide-child_one",
									"i18nKey":    "route.function_hide-child_one",
									"icon":       "material-symbols:filter-list-off",
									"hideInMenu": true,
									"activeMenu": "function_hide-child",
								},
							},
							{
								"name":      "function_hide-child_three",
								"path":      "/function/hide-child/three",
								"component": "view.function_hide-child_three",
								"meta": map[string]interface{}{
									"title":      "function_hide-child_three",
									"i18nKey":    "route.function_hide-child_three",
									"hideInMenu": true,
									"activeMenu": "function_hide-child",
								},
							},
							{
								"name":      "function_hide-child_two",
								"path":      "/function/hide-child/two",
								"component": "view.function_hide-child_two",
								"meta": map[string]interface{}{
									"title":      "function_hide-child_two",
									"i18nKey":    "route.function_hide-child_two",
									"hideInMenu": true,
									"activeMenu": "function_hide-child",
								},
							},
						},
					},
					{
						"name":      "function_multi-tab",
						"path":      "/function/multi-tab",
						"component": "view.function_multi-tab",
						"meta": map[string]interface{}{
							"title":      "function_multi-tab",
							"i18nKey":    "route.function_multi-tab",
							"icon":       "ic:round-tab",
							"multiTab":   true,
							"hideInMenu": true,
							"activeMenu": "function_tab",
						},
					},
					{
						"name":      "function_request",
						"path":      "/function/request",
						"component": "view.function_request",
						"meta": map[string]interface{}{
							"title":   "function_request",
							"i18nKey": "route.function_request",
							"icon":    "carbon:network-overlay",
							"order":   3,
						},
					},
					{
						"name":      "function_super-page",
						"path":      "/function/super-page",
						"component": "view.function_super-page",
						"meta": map[string]interface{}{
							"title":   "function_super-page",
							"i18nKey": "route.function_super-page",
							"icon":    "ic:round-supervisor-account",
							"order":   5,
						},
					},
					{
						"name":      "function_tab",
						"path":      "/function/tab",
						"component": "view.function_tab",
						"meta": map[string]interface{}{
							"title":   "function_tab",
							"i18nKey": "route.function_tab",
							"icon":    "ic:round-tab",
							"order":   1,
						},
					},
					{
						"name":      "function_toggle-auth",
						"path":      "/function/toggle-auth",
						"component": "view.function_toggle-auth",
						"meta": map[string]interface{}{
							"title":   "function_toggle-auth",
							"i18nKey": "route.function_toggle-auth",
							"icon":    "ic:round-construction",
							"order":   4,
						},
					},
				},
			},
			{
				"name":      "home",
				"path":      "/home",
				"component": "layout.base$view.home",
				"meta": map[string]interface{}{
					"title":   "home",
					"i18nKey": "route.home",
					"icon":    "mdi:monitor-dashboard",
					"order":   1,
				},
			},
			{
				"name":      "manage",
				"path":      "/manage",
				"component": "layout.base",
				"meta": map[string]interface{}{
					"title":   "manage",
					"i18nKey": "route.manage",
					"icon":    "carbon:cloud-service-management",
					"order":   9,
				},
				"children": []map[string]interface{}{
					{
						"name":      "manage_menu",
						"path":      "/manage/menu",
						"component": "view.manage_menu",
						"meta": map[string]interface{}{
							"title":     "manage_menu",
							"i18nKey":   "route.manage_menu",
							"icon":      "material-symbols:route",
							"order":     3,
							"keepAlive": true,
						},
					},
					{
						"name":      "manage_role",
						"path":      "/manage/role",
						"component": "view.manage_role",
						"meta": map[string]interface{}{
							"title":   "manage_role",
							"i18nKey": "route.manage_role",
							"icon":    "carbon:user-role",
							"order":   2,
						},
					},
					{
						"name":      "manage_user",
						"path":      "/manage/user",
						"component": "view.manage_user",
						"meta": map[string]interface{}{
							"title":   "manage_user",
							"i18nKey": "route.manage_user",
							"icon":    "ic:round-manage-accounts",
							"order":   1,
						},
					},
					{
						"name":      "manage_user-detail",
						"path":      "/manage/user-detail/:id",
						"component": "view.manage_user-detail",
						"props":     true,
						"meta": map[string]interface{}{
							"title":      "manage_user-detail",
							"i18nKey":    "route.manage_user-detail",
							"hideInMenu": true,
							"activeMenu": "manage_user",
						},
					},
				},
			},
			{
				"name":      "multi-menu",
				"path":      "/multi-menu",
				"component": "layout.base",
				"meta": map[string]interface{}{
					"title":   "multi-menu",
					"i18nKey": "route.multi-menu",
					"order":   8,
				},
				"children": []map[string]interface{}{
					{
						"name": "multi-menu_first",
						"path": "/multi-menu/first",
						"meta": map[string]interface{}{
							"title":   "multi-menu_first",
							"i18nKey": "route.multi-menu_first",
							"order":   1,
						},
						"children": []map[string]interface{}{
							{
								"name":      "multi-menu_first_child",
								"path":      "/multi-menu/first/child",
								"component": "view.multi-menu_first_child",
								"meta": map[string]interface{}{
									"title":   "multi-menu_first_child",
									"i18nKey": "route.multi-menu_first_child",
								},
							},
						},
					},
					{
						"name": "multi-menu_second",
						"path": "/multi-menu/second",
						"meta": map[string]interface{}{
							"title":   "multi-menu_second",
							"i18nKey": "route.multi-menu_second",
							"order":   2,
						},
						"children": []map[string]interface{}{
							{
								"name": "multi-menu_second_child",
								"path": "/multi-menu/second/child",
								"meta": map[string]interface{}{
									"title":   "multi-menu_second_child",
									"i18nKey": "route.multi-menu_second_child",
								},
								"children": []map[string]interface{}{
									{
										"name":      "multi-menu_second_child_home",
										"path":      "/multi-menu/second/child/home",
										"component": "view.multi-menu_second_child_home",
										"meta": map[string]interface{}{
											"title":   "multi-menu_second_child_home",
											"i18nKey": "route.multi-menu_second_child_home",
										},
									},
								},
							},
						},
					},
				},
			},
			{
				"name":      "user-center",
				"path":      "/user-center",
				"component": "layout.base$view.user-center",
				"meta": map[string]interface{}{
					"title":      "user-center",
					"i18nKey":    "route.user-center",
					"hideInMenu": true,
				},
			},
		},
		"home": "home",
	}

	return ctx.JSONOk("请求成功", data)
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
