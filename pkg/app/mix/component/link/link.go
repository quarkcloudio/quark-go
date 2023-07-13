package link

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Href          string      `json:"href"`
	Text          string      `json:"text"`
	Download      string      `json:"download"`
	ShowUnderLine bool        `json:"showUnderLine"`
	CopyTips      string      `json:"copyTips"`
	Color         string      `json:"color"`
	FontSize      string      `json:"fontSize"`
	Body          interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "link"
	p.SetShowUnderLine(true)
	p.SetCopyTips("已自动复制网址，请在手机浏览器里粘贴该网址")
	p.SetColor("#999999")
	p.SetFontSize("14")
	p.SetKey("link", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 应用内的跳转链接，值为相对路径或绝对路径，如："../first/first"，"/pages/first/first"，注意不能加 .vue 后缀
func (p *Component) SetHref(href string) *Component {
	p.Href = href
	return p
}

// 跳转方式
func (p *Component) SetText(text string) *Component {
	p.Text = text

	return p
}

// 当 open-type 为 'navigateBack' 时有效，表示回退的层数
func (p *Component) SetDownload(download string) *Component {
	p.Download = download

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口的显示/关闭动画效果，详见：窗口动画
func (p *Component) SetShowUnderLine(showUnderLine bool) *Component {
	p.ShowUnderLine = showUnderLine

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口显示/关闭动画的持续时间。
func (p *Component) SetCopyTips(copyTips string) *Component {
	p.CopyTips = copyTips

	return p
}

// 指定点击时的样式类，当hover-class="none"时，没有点击态效果
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 指定是否阻止本节点的祖先节点出现点击态
func (p *Component) SetFontSize(fontSize string) *Component {
	p.FontSize = fontSize

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "link"

	return p
}
