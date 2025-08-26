package login

import (
	"github.com/quarkcloudio/quark-go/v4"
)

type Loginer interface {

	// 模版接口
	quark.Templater

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

	// 验证码ID
	CaptchaId(ctx *quark.Context) error

	// 生成验证码
	Captcha(ctx *quark.Context) error

	// 字段
	Fields(ctx *quark.Context) []interface{}

	// 登录方法
	Handle(ctx *quark.Context) error

	// 退出方法
	Logout(ctx *quark.Context) error

	// 包裹在组件内的创建页字段
	FieldsWithinComponents(ctx *quark.Context) interface{}

	// 组件渲染
	Render(ctx *quark.Context) error
}
