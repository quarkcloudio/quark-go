package modal

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Title                  string        `json:"title"`
	BodyStyle              interface{}   `json:"bodyStyle"`
	Centered               bool          `json:"centered"`
	Closable               bool          `json:"closable"`
	DestroyOnClose         bool          `json:"destroyOnClose"`
	FocusTriggerAfterClose bool          `json:"focusTriggerAfterClose"`
	Keyboard               bool          `json:"keyboard"`
	Mask                   bool          `json:"mask"`
	MaskClosable           bool          `json:"maskClosable"`
	MaskStyle              interface{}   `json:"maskStyle"`
	Open                   bool          `json:"open"`
	Width                  int           `json:"width"`
	ZIndex                 int           `json:"zIndex"`
	Actions                []interface{} `json:"actions"`
	Body                   interface{}   `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "modal"
	p.SetKey("modal", component.DEFAULT_CRYPT)
	p.Closable = true
	p.Keyboard = true
	p.Mask = true
	p.MaskClosable = true
	p.Width = 520
	p.ZIndex = 1000

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// Modal body 样式
func (p *Component) SetBodyStyle(style interface{}) *Component {
	p.BodyStyle = style

	return p
}

// 容器控件里面的内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 垂直居中展示 Modal
func (p *Component) SetCentered(centered bool) *Component {
	p.Centered = centered

	return p
}

// 是否显示右上角的关闭按钮
func (p *Component) SetClosable(closable bool) *Component {
	p.Closable = closable

	return p
}

// 关闭时销毁 Modal 里的子元素
func (p *Component) SetDestroyOnClose(destroyOnClose bool) *Component {
	p.DestroyOnClose = destroyOnClose

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetFocusTriggerAfterClose(focusTriggerAfterClose bool) *Component {
	p.FocusTriggerAfterClose = focusTriggerAfterClose

	return p
}

// 是否支持键盘 esc 关闭
func (p *Component) SetKeyboard(keyboard bool) *Component {
	p.Keyboard = keyboard

	return p
}

// 是否展示遮罩
func (p *Component) SetMask(mask bool) *Component {
	p.Mask = mask

	return p
}

// 点击蒙层是否允许关闭
func (p *Component) SetMaskClosable(maskClosable bool) *Component {
	p.MaskClosable = maskClosable

	return p
}

// 遮罩样式
func (p *Component) SetMaskStyle(style map[string]interface{}) *Component {
	p.MaskStyle = style

	return p
}

// 对话框是否可见
func (p *Component) SetOpen(open bool) *Component {
	p.Open = open

	return p
}

// 宽度
func (p *Component) SetWidth(width int) *Component {
	p.Width = width

	return p
}

// 设置 Modal 的 z-index
func (p *Component) SetZIndex(zIndex int) *Component {
	p.ZIndex = zIndex

	return p
}

// 弹窗行为
func (p *Component) SetActions(actions []interface{}) *Component {
	p.Actions = actions

	return p
}
