package logins

import (
	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/icon"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/login"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-go/v2/pkg/utils/datetime"
	"github.com/quarkcloudio/quark-go/v2/pkg/utils/hash"
	"gorm.io/gorm"
)

type Index struct {
	login.Template
}

type Captcha struct {
	Id    string `json:"id" form:"id"`
	Value string `json:"value" form:"value"`
}

type LoginRequest struct {
	Username string   `json:"username" form:"username"`
	Password string   `json:"password" form:"password"`
	Captcha  *Captcha `json:"captcha" form:"captcha"`
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {

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
func (p *Index) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	// 获取验证码ID链接
	captchaIdUrl := ctx.RouterPathToUrl("/api/admin/login/index/captchaId")

	// 验证码链接
	captchaUrl := ctx.RouterPathToUrl("/api/admin/login/index/captcha/:id")

	return []interface{}{
		field.Text("username").
			SetRules([]*rule.Rule{
				rule.Required(true, "请输入用户名"),
			}).
			SetPlaceholder("用户名").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-user")),

		field.Password("password").
			SetRules([]*rule.Rule{
				rule.Required(true, "请输入密码"),
			}).
			SetPlaceholder("密码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-lock")),

		field.ImageCaptcha("captcha").
			SetCaptchaIdUrl(captchaIdUrl).
			SetCaptchaUrl(captchaUrl).
			SetRules([]*rule.Rule{
				rule.Required(true, "请输入验证码"),
			}).
			SetPlaceholder("验证码").
			SetWidth("100%").
			SetSize("large").
			SetPrefix(icon.New().SetType("icon-safetycertificate")),
	}
}

// 登录方法
func (p *Index) Handle(ctx *builder.Context) error {
	loginRequest := &LoginRequest{}
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

	adminInfo, err := (&model.Admin{}).GetInfoByUsername(loginRequest.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(200, message.Error("用户不存在"))
		}
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 检验账号和密码
	if !hash.Check(adminInfo.Password, loginRequest.Password) {
		return ctx.JSON(200, message.Error("用户名或密码错误"))
	}

	// 更新登录信息
	(&model.Admin{}).UpdateLastLogin(adminInfo.Id, ctx.ClientIP(), datetime.TimeNow())

	// 获取token字符串
	tokenString, err := ctx.JwtToken((&model.Admin{}).GetClaims(adminInfo))
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("登录成功", "", map[string]string{
		"token": tokenString,
	}))
}
