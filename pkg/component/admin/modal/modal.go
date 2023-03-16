package modal

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Modal struct {
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

// 初始化
func (p *Modal) Init() *Modal {
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
func (p *Modal) SetStyle(style map[string]interface{}) *Modal {
	p.Style = style

	return p
}

// 标题
func (p *Modal) SetTitle(title string) *Modal {
	p.Title = title

	return p
}

// Modal body 样式
func (p *Modal) SetBodyStyle(style interface{}) *Modal {
	p.BodyStyle = style

	return p
}

// 容器控件里面的内容
func (p *Modal) SetBody(body interface{}) *Modal {
	p.Body = body

	return p
}

// 垂直居中展示 Modal
func (p *Modal) SetCentered(centered bool) *Modal {
	p.Centered = centered

	return p
}

// 是否显示右上角的关闭按钮
func (p *Modal) SetClosable(closable bool) *Modal {
	p.Closable = closable

	return p
}

// 关闭时销毁 Modal 里的子元素
func (p *Modal) SetDestroyOnClose(destroyOnClose bool) *Modal {
	p.DestroyOnClose = destroyOnClose

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Modal) SetFocusTriggerAfterClose(focusTriggerAfterClose bool) *Modal {
	p.FocusTriggerAfterClose = focusTriggerAfterClose

	return p
}

// 是否支持键盘 esc 关闭
func (p *Modal) SetKeyboard(keyboard bool) *Modal {
	p.Keyboard = keyboard

	return p
}

// 是否展示遮罩
func (p *Modal) SetMask(mask bool) *Modal {
	p.Mask = mask

	return p
}

// 点击蒙层是否允许关闭
func (p *Modal) SetMaskClosable(maskClosable bool) *Modal {
	p.MaskClosable = maskClosable

	return p
}

// 遮罩样式
func (p *Modal) SetMaskStyle(style map[string]interface{}) *Modal {
	p.MaskStyle = style

	return p
}

// 对话框是否可见
func (p *Modal) SetOpen(open bool) *Modal {
	p.Open = open

	return p
}

// 宽度
func (p *Modal) SetWidth(width int) *Modal {
	p.Width = width

	return p
}

// 设置 Modal 的 z-index
func (p *Modal) SetZIndex(zIndex int) *Modal {
	p.ZIndex = zIndex

	return p
}

// 弹窗行为
func (p *Modal) SetActions(actions []interface{}) *Modal {
	p.Actions = actions

	return p
}
