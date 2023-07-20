package login

import "github.com/quarkcms/quark-go/v2/pkg/builder"

type Loginer interface {

	// 模版接口
	builder.Templater

	// 获取登录接口
	GetApi() string

	// 获取登录成功后跳转地址
	GetRedirect() string

	// 获取登录页面Logo
	GetLogo() interface{}

	// 获取登录页面标题
	GetTitle() string

	// 获取登录页面子标题
	GetSubTitle() string

	// 验证码存储驱动，redis | memory
	CaptchaStore(store string)

	// 验证码ID
	CaptchaId(ctx *builder.Context) error

	// 生成验证码
	Captcha(ctx *builder.Context) error

	// 登录方法
	Handle(ctx *builder.Context) error

	// 退出方法
	Logout(ctx *builder.Context) error

	// 组件渲染
	Render(ctx *builder.Context) error
}
