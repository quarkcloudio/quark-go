package login

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Component struct {
	component.Element
	Component    string                   `json:"component"`
	Api          string                   `json:"api"`
	Redirect     string                   `json:"redirect"`
	Logo         interface{}              `json:"logo"`
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	CaptchaIdUrl string                   `json:"captchaIdUrl"`
	CaptchaUrl   string                   `json:"captchaUrl"`
	Copyright    string                   `json:"copyright"`
	Links        []map[string]interface{} `json:"links"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "login"

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

// 描述
func (p *Component) SetDescription(description string) *Component {
	p.Description = description
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

// 页脚版权信息
func (p *Component) SetCopyright(copyright string) *Component {
	p.Copyright = copyright
	return p
}

// 页脚友情链接
func (p *Component) SetLinks(links []map[string]interface{}) *Component {
	p.Links = links
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "login"

	return p
}
