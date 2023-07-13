package swiper

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	IndicatorDots                bool        `json:"indicatorDots"`
	IndicatorColor               string      `json:"indicatorColor"`
	IndicatorActiveColor         string      `json:"indicatorActiveColor"`
	ActiveClass                  string      `json:"activeClass"`
	ChangingClass                string      `json:"changingClass"`
	Autoplay                     bool        `json:"autoplay"`
	Current                      int         `json:"current"`
	CurrentItemId                string      `json:"currentItemId"`
	Interval                     int         `json:"interval"`
	Duration                     int         `json:"duration"`
	Circular                     bool        `json:"circular"`
	Vertical                     bool        `json:"vertical"`
	PreviousMargin               string      `json:"previousMargin"`
	NextMargin                   string      `json:"nextMargin"`
	Acceleration                 bool        `json:"acceleration"`
	DisableProgrammaticAnimation bool        `json:"disableProgrammaticAnimation"`
	DisplayMultipleItems         int         `json:"displayMultipleItems"`
	SkipHiddenItemLayout         bool        `json:"skipHiddenItemLayout"`
	DisableTouch                 bool        `json:"disableTouch"`
	Touchable                    bool        `json:"touchable"`
	EasingFunction               string      `json:"easingFunction"`
	Items                        interface{} `json:"items"`
	ItemStyle                    interface{} `json:"itemStyle"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "swiper"
	p.SetKey("swiper", component.DEFAULT_CRYPT)
	p.IndicatorColor = "rgba(0, 0, 0, .3)"
	p.IndicatorActiveColor = "#000000"
	p.Interval = 5000
	p.Duration = 500
	p.DisplayMultipleItems = 1
	p.Touchable = true
	p.EasingFunction = "default"

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// Set ItemStyle.
func (p *Component) SetItemStyle(style interface{}) *Component {
	p.ItemStyle = style

	return p
}

// 是否显示面板指示点
func (p *Component) SetIndicatorDots(indicatorDots bool) *Component {
	p.IndicatorDots = indicatorDots

	return p
}

// 指示点颜色
func (p *Component) SetIndicatorColor(indicatorColor string) *Component {
	p.IndicatorColor = indicatorColor
	return p
}

// 当前选中的指示点颜色
func (p *Component) SetIndicatorActiveColor(indicatorActiveColor string) *Component {
	p.IndicatorActiveColor = indicatorActiveColor

	return p
}

// swiper-item 可见时的 class
func (p *Component) SetActiveClass(activeClass string) *Component {
	p.ActiveClass = activeClass

	return p
}

// acceleration 设置为 true 时且处于滑动过程中，中间若干屏处于可见时的class
func (p *Component) SetChangingClass(changingClass string) *Component {
	p.ChangingClass = changingClass

	return p
}

// 是否自动切换
func (p *Component) SetAutoplay(autoplay bool) *Component {
	p.Autoplay = autoplay

	return p
}

// 当前所在滑块的 index
func (p *Component) SetCurrent(current int) *Component {
	p.Current = current

	return p
}

// 当前所在滑块的 item-id ，不能与 current 被同时指定
func (p *Component) SetCurrentItemId(currentItemId string) *Component {
	p.CurrentItemId = currentItemId

	return p
}

// 自动切换时间间隔
func (p *Component) SetInterval(interval int) *Component {
	p.Interval = interval

	return p
}

// 滑动动画时长
func (p *Component) SetDuration(duration int) *Component {
	p.Duration = duration

	return p
}

// 是否采用衔接滑动，即播放到末尾后重新回到开头
func (p *Component) SetCircular(circular bool) *Component {
	p.Circular = circular

	return p
}

// 滑动方向是否为纵向
func (p *Component) SetVertical(vertical bool) *Component {
	p.Vertical = vertical

	return p
}

// 前边距，可用于露出前一项的一小部分，接受 px 和 rpx 值
func (p *Component) SetPreviousMargin(previousMargin string) *Component {
	p.PreviousMargin = previousMargin

	return p
}

// 后边距，可用于露出后一项的一小部分，接受 px 和 rpx 值
func (p *Component) SetNextMargin(nextMargin string) *Component {
	p.NextMargin = nextMargin

	return p
}

// 当开启时，会根据滑动速度，连续滑动多屏
func (p *Component) SetAcceleration(acceleration bool) *Component {
	p.Acceleration = acceleration

	return p
}

// 是否禁用代码变动触发 swiper 切换时使用动画。
func (p *Component) SetDisableProgrammaticAnimation(disableProgrammaticAnimation bool) *Component {
	p.DisableProgrammaticAnimation = disableProgrammaticAnimation

	return p
}

// 同时显示的滑块数量
func (p *Component) SetDisplayMultipleItems(displayMultipleItems int) *Component {
	p.DisplayMultipleItems = displayMultipleItems

	return p
}

// 是否跳过未显示的滑块布局，设为 true 可优化复杂情况下的滑动性能，但会丢失隐藏状态滑块的布局信息
func (p *Component) SetSkipHiddenItemLayout(skipHiddenItemLayout bool) *Component {
	p.SkipHiddenItemLayout = skipHiddenItemLayout

	return p
}

// 是否禁止用户 touch 操作
func (p *Component) SetDisableTouch(disableTouch bool) *Component {
	p.DisableTouch = disableTouch

	return p
}

// 是否监听用户的触摸事件，只在初始化时有效，不能动态变更
func (p *Component) SetTouchable(touchable bool) *Component {
	p.Touchable = touchable

	return p
}

// 指定 swiper 切换缓动动画类型，有效值：default、linear、easeInCubic、easeOutCubic、easeInOutCubic
func (p *Component) SetEasingFunction(easingFunction string) *Component {
	p.EasingFunction = easingFunction

	return p
}

// 轮播图的数据，通过数组长度决定指示点个数
func (p *Component) SetItems(items interface{}) *Component {
	p.Items = items

	return p
}

// 轮播图的数据，通过数组长度决定指示点个数
func (p *Component) SetBody(body interface{}) *Component {
	p.Items = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "swiper"

	return p
}
