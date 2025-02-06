package logins

import (
	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto/request"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/icon"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/login"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type Index struct {
	login.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// 登录页面Logo
	p.Logo = false

	// 登录页面标题
	p.Title = "QuarkGo"

	// 登录页面子标题
	p.SubTitle = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	// 登录后跳转地址
	p.Redirect = "/layout/index?api=/api/admin/dashboard/index/index"

	return p
}

// 字段
func (p *Index) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 获取验证码ID链接
	captchaIdUrl := ctx.RouterPathToUrl("/api/admin/login/index/captchaId")

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/login/index/captcha/:id")

	return []interface{}{
		field.Text("username").
			SetRules([]rule.Rule{
				rule.Required("请输入用户名"),
			}).
			SetPlaceholder("用户名").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-user")),

		field.Password("password").
			SetRules([]rule.Rule{
				rule.Required("请输入密码"),
			}).
			SetPlaceholder("密码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-lock")),

		field.ImageCaptcha("captcha").
			SetCaptchaIdUrl(captchaIdUrl).
			SetCaptchaUrl(captchaUrl).
			SetRules([]rule.Rule{
				rule.Required("请输入验证码"),
			}).
			SetPlaceholder("验证码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-safetycertificate")),
	}
}

// 登录方法
func (p *Index) Handle(ctx *quark.Context) error {
	loginRequest := &request.LoginReq{}
	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if loginRequest.Captcha.Id == "" || loginRequest.Captcha.Value == "" {
		return ctx.JSON(200, message.Error("验证码不能为空"))
	}

	verifyResult := captcha.VerifyString(loginRequest.Captcha.Id, loginRequest.Captcha.Value)
	if !verifyResult {
		return ctx.JSON(200, message.Error("验证码错误"))
	}
	captcha.Reload(loginRequest.Captcha.Id)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return ctx.JSON(200, message.Error("用户名或密码不能为空"))
	}
	token, err := service.NewAuthService(ctx).AdminLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return ctx.JSON(200, message.Error("用户名或密码错误"))
	}

	return ctx.JSON(200, message.Success("登录成功", "", map[string]string{
		"token": token,
	}))
}
