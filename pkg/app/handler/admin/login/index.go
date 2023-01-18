package login

import (
	"bytes"

	"github.com/dchest/captcha"
	"github.com/golang-jwt/jwt/v4"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminlogin"
	"github.com/quarkcms/quark-go/pkg/hash"
	"github.com/quarkcms/quark-go/pkg/msg"
)

type Index struct {
	adminlogin.Template
}

type LoginRequest struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
	Captcha   string `json:"captcha" form:"captcha"`
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	// 登录页面Logo
	p.Logo = false

	// 登录页面标题
	p.Title = "QuarkGo"

	// 登录页面描述
	p.Description = "信息丰富的世界里，唯一稀缺的就是人类的注意力"

	// 登录后跳转地址
	p.Redirect = "/index?api=/api/admin/dashboard/index/index"

	return p
}

// 验证码ID
func (p *Index) CaptchaId(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Success("获取成功", "", map[string]string{
		"captchaId": captcha.NewLen(4),
	})
}

// 生成验证码
func (p *Index) Captcha(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	id := request.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)

	return writer.Bytes()
}

// 登录方法
func (p *Index) Handle(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	loginRequest := &LoginRequest{}
	if err := request.BodyParser(loginRequest); err != nil {
		return msg.Error(err.Error(), "")
	}
	if loginRequest.CaptchaId == "" || loginRequest.Captcha == "" {
		return msg.Error("验证码不能为空", "")
	}

	verifyResult := captcha.VerifyString(loginRequest.CaptchaId, loginRequest.Captcha)
	if !verifyResult {
		return msg.Error("验证码错误", "")
	}
	captcha.Reload(loginRequest.CaptchaId)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return msg.Error("用户名或密码不能为空", "")
	}

	adminInfo, err := (&model.Admin{}).GetInfoByUsername(loginRequest.Username)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 检验账号和密码
	if !hash.Check(adminInfo.Password, loginRequest.Password) {
		return msg.Error("用户名或密码错误", "")
	}

	config := builder.GetConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, (&model.Admin{}).GetClaims(adminInfo))

	// 获取token字符串
	tokenString, err := token.SignedString([]byte(config.AppKey))

	return msg.Success("获取成功", "", map[string]string{
		"token": tokenString,
	})
}

// 退出方法
func (p *Index) Logout(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {

	return msg.Success("退出成功", "", "")
}
