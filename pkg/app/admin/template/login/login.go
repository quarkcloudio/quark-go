package login

import (
	"bytes"
	"reflect"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/login"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	redisclient "github.com/quarkcloudio/quark-go/v2/pkg/dal/redis"
)

// 后台登录模板
type Template struct {
	builder.Template
	Api      string      // 登录接口
	Redirect string      // 登录后跳转地址
	Logo     interface{} // 登录页面Logo
	Title    string      // 标题
	SubTitle string      // 子标题
	Body     interface{} `json:"body,omitempty"` // 表单内容
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

// 字段
func (p *Template) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 登录方法
func (p *Template) Handle(ctx *builder.Context) error {
	return ctx.JSON(200, message.Error("请实现登录方法"))
}

// 退出方法
func (p *Template) Logout(ctx *builder.Context) error {
	return ctx.JSON(200, message.Error("退出成功"))
}

// 包裹在组件内的创建页字段
func (p *Template) FieldsWithinComponents(ctx *builder.Context) interface{} {

	// 资源实例
	template := ctx.Template.(Loginer)

	// 获取字段
	fields := template.Fields(ctx)

	// 解析创建页表单组件内的字段
	items := p.FormFieldsParser(ctx, fields)

	return items
}

// 解析创建页表单组件内的字段
func (p *Template) FormFieldsParser(ctx *builder.Context, fields interface{}) interface{} {
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

// 在标签页内的From组件
func (p *Template) FormWithinTabs(
	ctx *builder.Context,
	title string,
	extra interface{},
	api string,
	fields interface{},
	actions []interface{},
	data map[string]interface{}) interface{} {

	tabsComponent := (&tabs.Component{}).Init().SetTabPanes(fields).SetTabBarExtraContent(extra)

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 返回数据
	return template.
		GetForm().
		SetStyle(map[string]interface{}{
			"backgroundColor": "#fff",
			"paddingBottom":   "20px",
		}).
		SetApi(api).
		SetActions(actions).
		SetBody(tabsComponent).
		SetInitialValues(data)
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

	// 包裹在组件内的字段
	fields := p.FieldsWithinComponents(ctx)

	// 组件
	component := (&login.Component{}).
		Init().
		SetApi(loginApi).
		SetRedirect(redirect).
		SetLogo(logo).
		SetTitle(title).
		SetSubTitle(subTitle).
		SetBody(fields)

	// 解析tabPane组件
	if _, ok := fields.([]interface{}); ok {
		componentName := reflect.
			ValueOf(fields.([]interface{})[0]).
			Elem().
			FieldByName("Component").
			String()

		if componentName == "tabPane" {
			tabComponent := (&tabs.Component{}).Init().SetTabPanes(fields)

			// 组件
			component = (&login.Component{}).
				Init().
				SetApi(loginApi).
				SetRedirect(redirect).
				SetLogo(logo).
				SetTitle(title).
				SetSubTitle(subTitle).
				SetBody(tabComponent)
		}
	}

	return ctx.JSON(200, component)
}
