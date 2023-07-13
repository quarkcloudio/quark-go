package tabs

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Color        string  `json:"color"`
	Background   string  `json:"background"`
	Direction    string  `json:"direction"`
	Type         string  `json:"type"`
	Swipeable    bool    `json:"swipeable"`
	TitleScroll  bool    `json:"titleScroll"`
	Ellipsis     bool    `json:"ellipsis"`
	AnimatedTime int     `json:"animatedTime"`
	TitleGutter  int     `json:"titleGutter"`
	Size         string  `json:"size"`
	AutoHeight   bool    `json:"autoHeight"`
	Name         string  `json:"name"`
	Panes        []*Pane `json:"panes"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Pane
func NewPane() *Pane {
	return (&Pane{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "tabs"
	p.SetKey("tabs", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 标签选中色
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 标签栏背景颜色
func (p *Component) SetBackground(background string) *Component {
	p.Background = background

	return p
}

// 使用横纵方向 可选值 horizontal、vertical
func (p *Component) SetDirection(direction string) *Component {
	p.Direction = direction

	return p
}

// 选中底部展示样式 可选值 line、smile
func (p *Component) SetType(tabType string) *Component {
	p.Type = tabType

	return p
}

// 固定在底部时，是否在标签位置生成一个等高的占位元素
func (p *Component) SetSwipeable(swipeable bool) *Component {
	p.Swipeable = swipeable

	return p
}

// 标签栏是否可以滚动
func (p *Component) SetTitleScroll(titleScroll bool) *Component {
	p.TitleScroll = titleScroll

	return p
}

// 是否省略过长的标题文字
func (p *Component) SetEllipsis(ellipsis bool) *Component {
	p.Ellipsis = ellipsis

	return p
}

// 切换动画时长,单位 ms 0 代表无动画(小程序场景数据过大建议设置0，解决切换卡顿问题)
func (p *Component) SetAnimatedTime(animatedTime int) *Component {
	p.AnimatedTime = animatedTime

	return p
}

// 标签间隙
func (p *Component) SetTitleGutter(titleGutter int) *Component {
	p.TitleGutter = titleGutter

	return p
}

// 标签栏字体尺寸大小 可选值 large normal small
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 自动高度。设置为 true 时，nut-tabs 和 nut-tabs__content 会随着当前 nut-tab-pane 的高度而发生变化。
func (p *Component) SetAutoHeight(autoHeight bool) *Component {
	p.AutoHeight = autoHeight

	return p
}

// 在taro环境下，必须设置name以开启标题栏自动滚动功能。
func (p *Component) SetName(name string) *Component {
	p.Name = name

	return p
}

// 选项卡内容
func (p *Component) SetPanes(panes []*Pane) *Component {
	p.Panes = panes

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "tabs"

	return p
}
