package fixednav

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Position struct {
	Top    interface{} `json:"top"`
	Bottom interface{} `json:"bottom"`
}

type Component struct {
	component.Element
	Visible      bool        `json:"visible"`
	NavList      interface{} `json:"navList"`
	ActiveColor  string      `json:"activeColor"`
	ActiveText   string      `json:"activeText"`
	UnActiveText string      `json:"unActiveText"`
	Type         string      `json:"type"`
	Overlay      bool        `json:"overlay"`
	Position     *Position   `json:"position"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "fixedNav"
	p.SetKey("fixedNav", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 是否打开
func (p *Component) SetVisible(visible bool) *Component {
	p.Visible = visible
	return p
}

// 悬浮列表内容数据
func (p *Component) SetNavList(navList interface{}) *Component {
	p.NavList = navList
	return p
}

// 选中按钮文案颜色
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor
	return p
}

// 收起列表按钮文案
func (p *Component) SetActiveText(activeText string) *Component {
	p.ActiveText = activeText
	return p
}

// 展开列表按钮文案
func (p *Component) SetUnActiveText(unActiveText string) *Component {
	p.UnActiveText = unActiveText
	return p
}

// 导航方向,可选值 left right
func (p *Component) SetType(fixednavType string) *Component {
	p.Type = fixednavType
	return p
}

// 展开时是否显示遮罩
func (p *Component) SetOverlay(overlay bool) *Component {
	p.Overlay = overlay
	return p
}

// 展开时是否显示遮罩
func (p *Component) SetPosition(position *Position) *Component {
	p.Position = position
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "fixedNav"

	return p
}
