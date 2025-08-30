package auths

import (
	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/component/icon"
	"github.com/quarkcloudio/quark-go/v4/dto/request"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/auth"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
)

type Index struct {
	auth.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// 登录页面Logo
	p.Logo = false

	// 登录页面标题
	p.Title = "QuarkGo"

	// 登录后跳转地址
	p.Redirect = "/engine?api=/api/admin/dashboard/index/index"

	return p
}

// 字段
func (p *Index) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/auth/index/captcha")

	return []interface{}{
		field.Text("username").
			SetRules([]rule.Rule{
				rule.Required("请输入用户名"),
			}).
			SetPlaceholder("请输入用户名").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("ant-design:user-outlined")),

		field.Password("password").
			SetRules([]rule.Rule{
				rule.Required("请输入密码"),
			}).
			SetPlaceholder("请输入密码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("ant-design:lock-outlined")),

		field.ImageCaptcha("captcha").
			SetCaptchaUrl(captchaUrl).
			SetRules([]rule.Rule{
				rule.Required("请输入验证码"),
			}).
			SetPlaceholder("请输入验证码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("ant-design:safety-certificate-outlined")),
	}
}

// 登录方法
func (p *Index) Login(ctx *quark.Context) error {
	loginRequest := &request.LoginReq{}
	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSONError(err.Error())
	}
	if loginRequest.Captcha.Uuid == "" || loginRequest.Captcha.Value == "" {
		return ctx.JSONError("验证码不能为空")
	}

	verifyResult := captcha.VerifyString(loginRequest.Captcha.Uuid, loginRequest.Captcha.Value)
	if !verifyResult {
		return ctx.JSONError("验证码错误")
	}
	captcha.Reload(loginRequest.Captcha.Uuid)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return ctx.JSONError("用户名或密码不能为空")
	}
	token, err := service.NewAuthService(ctx).AdminLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return ctx.JSONError("用户名或密码错误")
	}

	return ctx.JSONOk("登录成功", map[string]string{
		"token": token,
	})
}
