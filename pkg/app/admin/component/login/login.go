package login

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type ActivityConfig struct {
	Title    string                 `json:"title,omitempty"`
	SubTitle string                 `json:"subTitle,omitempty"`
	Action   interface{}            `json:"action,omitempty"`
	Style    map[string]interface{} `json:"style,omitempty"`
}

type Component struct {
	component.Element
	Component          string          `json:"component"`
	Api                string          `json:"api,omitempty"`
	Redirect           string          `json:"redirect,omitempty"`
	Logo               interface{}     `json:"logo,omitempty"`
	Title              string          `json:"title,omitempty"`
	SubTitle           string          `json:"subTitle,omitempty"`
	BackgroundImageUrl string          `json:"backgroundImageUrl,omitempty"`
	CaptchaIdUrl       string          `json:"captchaIdUrl,omitempty"`
	CaptchaUrl         string          `json:"captchaUrl,omitempty"`
	LoginType          []string        `json:"loginType"`
	ActivityConfig     *ActivityConfig `json:"activityConfig,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "login"
	p.LoginType = []string{"account"}
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 登录接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	return p
}

// 登录后跳转地址
func (p *Component) SetRedirect(redirect string) *Component {
	p.Redirect = redirect
	return p
}

// Logo
func (p *Component) SetLogo(logo interface{}) *Component {
	p.Logo = logo
	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 子标题
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle
	return p
}

// 整个区域的背景图片配置，手机端不会展示
func (p *Component) SeBackgroundImageUrl(backgroundImageUrl string) *Component {
	p.BackgroundImageUrl = backgroundImageUrl
	return p
}

// 获取验证码ID链接
func (p *Component) SetCaptchaIdUrl(captchaIdUrl string) *Component {
	p.CaptchaIdUrl = captchaIdUrl
	return p
}

// 验证码链接
func (p *Component) SetCaptchaUrl(captchaUrl string) *Component {
	p.CaptchaUrl = captchaUrl
	return p
}

// 登录类型，{"account","phone"}
func (p *Component) SetLoginType(loginType []string) *Component {
	p.LoginType = loginType
	return p
}

// 活动的配置，包含 title，subTitle，action，分别代表标题，次标题和行动按钮，也可配置 style 来控制区域的样式
func (p *Component) SetActivityConfig(activityConfig *ActivityConfig) *Component {
	p.ActivityConfig = activityConfig
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "login"

	return p
}
