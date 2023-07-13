package row

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Title string      `json:"title"`
	Desc  string      `json:"desc"`
	Body  interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "cellGroup"
	p.SetKey("cellGroup", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 自定义 title 标题区域
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 自定义 desc 描述区域
func (p *Component) SetDesc(desc string) *Component {
	p.Desc = desc
	return p
}

// 组件内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "cellGroup"

	return p
}
