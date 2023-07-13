package footer

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Copyright string                   `json:"copyright"`
	Links     []map[string]interface{} `json:"links"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "footer"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 版权信息
func (p *Component) SetCopyright(copyright string) *Component {
	p.Copyright = copyright
	return p
}

// 版权信息
func (p *Component) SetLinks(links []map[string]interface{}) *Component {
	p.Links = links
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "footer"

	return p
}
