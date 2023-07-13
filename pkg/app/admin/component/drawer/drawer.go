package drawer

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Title               string        `json:"title"`
	BodyStyle           interface{}   `json:"bodyStyle"`
	Closable            bool          `json:"closable"`
	ContentWrapperStyle interface{}   `json:"contentWrapperStyle"`
	DestroyOnClose      bool          `json:"destroyOnClose"`
	DrawerStyle         interface{}   `json:"drawerStyle"`
	FooterStyle         interface{}   `json:"footerStyle"`
	Height              int           `json:"height"`
	Keyboard            bool          `json:"keyboard"`
	Mask                bool          `json:"mask"`
	MaskClosable        bool          `json:"maskClosable"`
	MaskStyle           interface{}   `json:"maskStyle"`
	Open                bool          `json:"open"`
	Width               int           `json:"width"`
	ZIndex              int           `json:"zIndex"`
	Actions             []interface{} `json:"actions"`
	Placement           string        `json:"placement"`
	Body                interface{}   `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "drawer"
	p.SetKey("drawer", component.DEFAULT_CRYPT)
	p.Closable = true
	p.FooterStyle = map[string]interface{}{
		"textAlign": "right",
	}
	p.Height = 256
	p.Keyboard = true
	p.Mask = true
	p.MaskClosable = true
	p.Placement = "right"
	p.Width = 256
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

// 是否显示右上角的关闭按钮
func (p *Component) SetClosable(closable bool) *Component {
	p.Closable = closable

	return p
}

// 可用于设置 Drawer 包裹内容部分的样式
func (p *Component) SetContentWrapperStyle(style interface{}) *Component {
	p.ContentWrapperStyle = style

	return p
}

// 关闭时销毁 Modal 里的子元素
func (p *Component) SetDestroyOnClose(destroyOnClose bool) *Component {
	p.DestroyOnClose = destroyOnClose

	return p
}

// 用于设置 Drawer 弹出层的样式
func (p *Component) SetDrawerStyle(style interface{}) *Component {
	p.DrawerStyle = style

	return p
}

// 抽屉页脚部件的样式
func (p *Component) SetFooterStyle(style interface{}) *Component {
	p.FooterStyle = style

	return p
}

// 高度, 在 placement 为 top 或 bottom 时使用
func (p *Component) SetHeight(height int) *Component {
	p.Height = height

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
func (p *Component) SetMaskStyle(style interface{}) *Component {
	p.MaskStyle = style

	return p
}

// 抽屉的方向,top | right | bottom | left
func (p *Component) SetPlacement(placement string) *Component {
	p.Placement = placement

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
