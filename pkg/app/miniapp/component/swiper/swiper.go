package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Width             int     `json:"width,omitempty"`
	Height            int     `json:"height,omitempty"`
	Direction         string  `json:"direction,omitempty"`
	PaginationVisible bool    `json:"paginationVisible,omitempty"`
	PaginationColor   string  `json:"paginationColor,omitempty"`
	Loop              bool    `json:"loop,omitempty"`
	Duration          int     `json:"duration,omitempty"`
	AutoPlay          int     `json:"autoPlay,omitempty"`
	InitPage          int     `json:"initPage,omitempty"`
	Touchable         bool    `json:"touchable,omitempty"`
	IsPreventDefault  bool    `json:"isPreventDefault,omitempty"`
	IsStopPropagation bool    `json:"isStopPropagation,omitempty"`
	Body              []*Item `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "swiper"
	p.SetKey("swiper", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 轮播卡片的宽度
func (p *Component) SetWidth(width int) *Component {
	p.Width = width
	return p
}

// 轮播卡片的高度
func (p *Component) SetHeight(height int) *Component {
	p.Height = height
	return p
}

// 轮播方向,可选值horizontal,vertical
func (p *Component) SetJustify(direction string) *Component {
	p.Direction = direction
	return p
}

// 分页指示器是否展示
func (p *Component) SetPaginationVisible(paginationVisible bool) *Component {
	p.PaginationVisible = paginationVisible
	return p
}

// 分页指示器选中的颜色
func (p *Component) SetPaginationColor(paginationColor string) *Component {
	p.PaginationColor = paginationColor
	return p
}

// 是否循环轮播
func (p *Component) SetLoop(loop bool) *Component {
	p.Loop = loop
	return p
}

// 动画时长（单位是ms）
func (p *Component) SetDuration(duration int) *Component {
	p.Duration = duration
	return p
}

// 自动轮播时长，0表示不会自动轮播
func (p *Component) SetAutoPlay(autoPlay int) *Component {
	p.AutoPlay = autoPlay
	return p
}

// 初始化索引值
func (p *Component) SetInitPage(initPage int) *Component {
	p.InitPage = initPage
	return p
}

// 是否可触摸滑动
func (p *Component) SetTouchable(touchable bool) *Component {
	p.Touchable = touchable
	return p
}

// 滑动过程中是否禁用默认事件
func (p *Component) SetIsPreventDefault(isPreventDefault bool) *Component {
	p.IsPreventDefault = isPreventDefault
	return p
}

// 滑动过程中是否禁止冒泡
func (p *Component) SetIsStopPropagation(isStopPropagation bool) *Component {
	p.IsStopPropagation = isStopPropagation
	return p
}

// 内容
func (p *Component) SetBody(body []*Item) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "swiper"

	return p
}
