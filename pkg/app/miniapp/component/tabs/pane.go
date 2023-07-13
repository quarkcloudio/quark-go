package tabs

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Pane struct {
	component.Element
	Title    string      `json:"title"`
	PaneKey  string      `json:"paneKey"`
	Disabled bool        `json:"disabled"`
	Body     interface{} `json:"body"`
}

// 初始化
func (p *Pane) Init() *Pane {
	p.Component = "tabPane"
	p.SetKey("tabPane", component.DEFAULT_CRYPT)
	return p
}

// Set style.
func (p *Pane) SetStyle(style interface{}) *Pane {
	p.Style = style
	return p
}

// 标题
func (p *Pane) SetTitle(title string) *Pane {
	p.Title = title
	return p
}

// 标签 Key , 匹配的标识符
func (p *Pane) SetPaneKey(paneKey string) *Pane {
	p.PaneKey = paneKey
	return p
}

// 是否禁用标签
func (p *Pane) SetDisabled(disabled bool) *Pane {
	p.Disabled = disabled
	return p
}

// 内容
func (p *Pane) SetBody(body interface{}) *Pane {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Pane) JsonSerialize() *Pane {
	p.Component = "tabPane"

	return p
}
