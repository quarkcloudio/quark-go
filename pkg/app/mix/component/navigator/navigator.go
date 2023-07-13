package navigator

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Url                  string      `json:"url"`
	OpenType             string      `json:"openType"`
	Delta                int         `json:"delta"`
	AnimationType        string      `json:"animationType"`
	AnimationDuration    int         `json:"animationDuration"`
	HoverClass           string      `json:"hoverClass"`
	HoverStopPropagation string      `json:"hoverStopPropagation"`
	HoverStartTime       int         `json:"hoverStartTime"`
	HoverStayTime        int         `json:"hoverStayTime"`
	Target               string      `json:"target"`
	Body                 interface{} `json:"body"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "navigator"
	p.OpenType = "navigate"
	p.AnimationType = "pop-in"
	p.AnimationDuration = 300
	p.HoverClass = "navigator-hover"
	p.HoverStartTime = 50
	p.HoverStayTime = 600
	p.Target = "self"
	p.SetKey("navigator", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 应用内的跳转链接，值为相对路径或绝对路径，如："../first/first"，"/pages/first/first"，注意不能加 .vue 后缀
func (p *Component) SetUrl(url string) *Component {
	p.Url = url
	return p
}

// 跳转方式
func (p *Component) SetOpenType(openType string) *Component {
	p.OpenType = openType

	return p
}

// 当 open-type 为 'navigateBack' 时有效，表示回退的层数
func (p *Component) SetDelta(delta int) *Component {
	p.Delta = delta

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口的显示/关闭动画效果，详见：窗口动画
func (p *Component) SetAnimationType(animationType string) *Component {
	p.AnimationType = animationType

	return p
}

// 当 open-type 为 navigate、navigateBack 时有效，窗口显示/关闭动画的持续时间。
func (p *Component) SetAnimationDuration(animationDuration int) *Component {
	p.AnimationDuration = animationDuration

	return p
}

// 指定点击时的样式类，当hover-class="none"时，没有点击态效果
func (p *Component) SetHoverClass(hoverClass string) *Component {
	p.HoverClass = hoverClass

	return p
}

// 指定是否阻止本节点的祖先节点出现点击态
func (p *Component) SetHoverStopPropagation(hoverStopPropagation string) *Component {
	p.HoverStopPropagation = hoverStopPropagation

	return p
}

// 是否固定顶部
func (p *Component) SetHoverStartTime(hoverStartTime int) *Component {
	p.HoverStartTime = hoverStartTime

	return p
}

// 是否包含状态栏
func (p *Component) SetHoverStayTime(hoverStayTime int) *Component {
	p.HoverStayTime = hoverStayTime

	return p
}

// 导航栏下是否有阴影
func (p *Component) SetTarget(target string) *Component {
	p.Target = target

	return p
}

// 内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "navigator"

	return p
}
