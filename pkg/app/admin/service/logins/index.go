package logins

import (
	"time"

	"github.com/dchest/captcha"
	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/login"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/utils/hash"
	"gorm.io/gorm"
)

type Index struct {
	login.Template
}

type LoginRequest struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
	Captcha   string `json:"captcha" form:"captcha"`
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

// 登录方法
func (p *Index) Handle(ctx *builder.Context) error {
	loginRequest := &LoginRequest{}
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if loginRequest.CaptchaId == "" || loginRequest.Captcha == "" {
		return ctx.JSON(200, message.Error("验证码不能为空"))
	}

	verifyResult := captcha.VerifyString(loginRequest.CaptchaId, loginRequest.Captcha)
	if !verifyResult {
		return ctx.JSON(200, message.Error("验证码错误"))
	}
	captcha.Reload(loginRequest.CaptchaId)

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

	config := ctx.Engine.GetConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, (&model.Admin{}).GetClaims(adminInfo))

	// 更新登录信息
	(&model.Admin{}).UpdateLastLogin(adminInfo.Id, ctx.ClientIP(), time.Now())

	// 获取token字符串
	tokenString, err := token.SignedString([]byte(config.AppKey))
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("登录成功", "", map[string]string{
		"token": tokenString,
	}))
}
