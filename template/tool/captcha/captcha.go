package captcha

import (
	"bytes"
	"reflect"
	"time"

	"github.com/dchest/captcha"
	"github.com/quarkcloudio/quark-go/v3"
	redisclient "github.com/quarkcloudio/quark-go/v3/dal/redis"
)

// 文件上传
type Template struct {
	quark.Template
	GetIdPath        string // 登录获取验证码ID路由
	GetImagePath     string // 登录验证码路由
	ToolGetIdPath    string // 登录获取验证码ID路由
	ToolGetImagePath string // 登录验证码路由
	Length           int    // 验证码长度
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.GetIdPath = "/api/captcha/:resource/getId"                    // 登录获取验证码ID路由
	p.GetImagePath = "/api/captcha/:resource/getImage/:id"          // 登录验证码路由
	p.ToolGetIdPath = "/api/tool/captcha/:resource/getId"           // 登录获取验证码ID路由
	p.ToolGetImagePath = "/api/tool/captcha/:resource/getImage/:id" // 登录验证码路由

	return p
}

// 初始化路由映射
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.GetIdPath, p.CaptchaId)      // 登录获取验证码ID路由
	p.GET(p.GetImagePath, p.Captcha)     // 登录验证码路由
	p.GET(p.ToolGetIdPath, p.CaptchaId)  // 登录获取验证码ID路由
	p.GET(p.ToolGetImagePath, p.Captcha) // 登录验证码路由

	return p
}

// 初始化模板
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

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

// 验证码ID
func (p *Template) CaptchaId(ctx *quark.Context) error {
	length := int(reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Length").Int())

	return ctx.JSONOk("获取成功", map[string]string{
		"captchaId": captcha.NewLen(length),
	})
}

// 生成验证码
func (p *Template) Captcha(ctx *quark.Context) error {
	id := ctx.Param("id")
	writer := bytes.Buffer{}
	captcha.WriteImage(&writer, id, 110, 38)
	ctx.Write(writer.Bytes())
	return nil
}
