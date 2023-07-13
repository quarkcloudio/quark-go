package menu

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	ActiveColor         string      `json:"activeColor,omitempty"`
	CloseOnClickOverlay bool        `json:"closeOnClickOverlay,omitempty"`
	ScrollFixed         interface{} `json:"scrollFixed,omitempty"`
	LockScroll          bool        `json:"lockScroll,omitempty"`
	Body                []*Item     `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "menu"
	p.SetKey("menu", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 选项的选中态图标颜色
func (p *Component) SetActiveColor(activeColor string) *Component {
	p.ActiveColor = activeColor
	return p
}

// 是否在点击遮罩层后关闭菜单
func (p *Component) SetCloseOnClickOverlay(closeOnClickOverlay bool) *Component {
	p.CloseOnClickOverlay = closeOnClickOverlay
	return p
}

// 背景是否锁定
func (p *Component) SetScrollFixed(scrollFixed string) *Component {
	p.ScrollFixed = scrollFixed
	return p
}

// 分页指示器是否展示
func (p *Component) SetLockScroll(lockScroll bool) *Component {
	p.LockScroll = lockScroll
	return p
}

// 内容
func (p *Component) SetBody(body []*Item) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "menu"

	return p
}
