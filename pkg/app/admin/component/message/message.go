package message

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	ClassName string      `json:"className"`
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	Duration  int         `json:"duration"`
	Icon      string      `json:"icon"`
	Style     interface{} `json:"style"`
	Data      interface{} `json:"data"`
	Url       string      `json:"url"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 返回成功，Success("成功") | Success("成功", "/home/index", map[string]interface{}{"title":"标题"})
func Success(message ...interface{}) *Component {
	var (
		content = ""
		url     = ""
		data    interface{}
	)

	if len(message) == 1 {
		content = message[0].(string)
	}
	if len(message) == 2 {
		content = message[0].(string)
		url = message[1].(string)
	}
	if len(message) >= 3 {
		content = message[0].(string)
		url = message[1].(string)
		data = message[2]
	}

	return (&Component{}).
		Init().
		SetType("success").
		SetContent(content).
		SetUrl(url).
		SetData(data)
}

// 返回失败，Error("错误") | Error("操作失败", "/home/index")
func Error(message ...interface{}) *Component {
	var (
		content = ""
		url     = ""
	)

	if len(message) == 1 {
		content = message[0].(string)
	}
	if len(message) == 2 {
		content = message[0].(string)
		url = message[1].(string)
	}

	return (&Component{}).
		Init().
		SetType("error").
		SetContent(content).
		SetUrl(url)
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "message"
	p.Type = "success"
	p.SetKey("message", component.DEFAULT_CRYPT)

	return p
}

// Set ClassName
func (p *Component) SetClassName(className string) *Component {
	p.ClassName = className

	return p
}

// Set Type info | success | error | warning | loading
func (p *Component) SetType(messageType string) *Component {
	p.Type = messageType

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 内容
func (p *Component) SetContent(content interface{}) *Component {
	p.Content = content
	return p
}

// 自动关闭的延时，单位秒。设为 0 时不自动关闭
func (p *Component) SetDuration(duration int) *Component {
	p.Duration = duration
	return p
}

// Set Icon
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = icon

	return p
}

// 设置返回数据
func (p *Component) SetData(data interface{}) *Component {
	p.Data = data
	return p
}

// 设置消息弹出后跳转链接
func (p *Component) SetUrl(url string) *Component {
	p.Url = url
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "message"

	return p
}
