package tabs

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type TabPane struct {
	component.Element
	Title string      `json:"title"`
	Body  interface{} `json:"body"`
}

// 初始化
func (p *TabPane) Init() *TabPane {
	p.Component = "tabPane"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *TabPane) SetStyle(style map[string]interface{}) *TabPane {
	p.Style = style

	return p
}

// 标签标题
func (p *TabPane) SetTitle(title string) *TabPane {
	p.Title = title

	return p
}

// 内容
func (p *TabPane) SetBody(body interface{}) *TabPane {
	p.Body = body

	return p
}
