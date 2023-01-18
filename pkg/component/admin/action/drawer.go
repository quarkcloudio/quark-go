package action

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Drawer struct {
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
	Visible             bool          `json:"visible"`
	Width               int           `json:"width"`
	ZIndex              int           `json:"zIndex"`
	Actions             []interface{} `json:"actions"`
	Placement           string        `json:"placement"`
	Body                interface{}   `json:"body"`
}

// 初始化
func (p *Drawer) Init() *Drawer {
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
func (p *Drawer) SetStyle(style map[string]interface{}) *Drawer {
	p.Style = style

	return p
}

// 标题
func (p *Drawer) SetTitle(title string) *Drawer {
	p.Title = title

	return p
}

// Modal body 样式
func (p *Drawer) SetBodyStyle(style interface{}) *Drawer {
	p.BodyStyle = style

	return p
}

// 容器控件里面的内容
func (p *Drawer) SetBody(body interface{}) *Drawer {
	p.Body = body

	return p
}

// 是否显示右上角的关闭按钮
func (p *Drawer) SetClosable(closable bool) *Drawer {
	p.Closable = closable

	return p
}

// 可用于设置 Drawer 包裹内容部分的样式
func (p *Drawer) SetContentWrapperStyle(style interface{}) *Drawer {
	p.ContentWrapperStyle = style

	return p
}

// 关闭时销毁 Modal 里的子元素
func (p *Drawer) SetDestroyOnClose(destroyOnClose bool) *Drawer {
	p.DestroyOnClose = destroyOnClose

	return p
}

// 用于设置 Drawer 弹出层的样式
func (p *Drawer) SetDrawerStyle(style interface{}) *Drawer {
	p.DrawerStyle = style

	return p
}

// 抽屉页脚部件的样式
func (p *Drawer) SetFooterStyle(style interface{}) *Drawer {
	p.FooterStyle = style

	return p
}

// 高度, 在 placement 为 top 或 bottom 时使用
func (p *Drawer) SetHeight(height int) *Drawer {
	p.Height = height

	return p
}

// 是否支持键盘 esc 关闭
func (p *Drawer) SetKeyboard(keyboard bool) *Drawer {
	p.Keyboard = keyboard

	return p
}

// 是否展示遮罩
func (p *Drawer) SetMask(mask bool) *Drawer {
	p.Mask = mask

	return p
}

// 点击蒙层是否允许关闭
func (p *Drawer) SetMaskClosable(maskClosable bool) *Drawer {
	p.MaskClosable = maskClosable

	return p
}

// 遮罩样式
func (p *Drawer) SetMaskStyle(style interface{}) *Drawer {
	p.MaskStyle = style

	return p
}

// 抽屉的方向,top | right | bottom | left
func (p *Drawer) SetPlacement(placement string) *Drawer {
	p.Placement = placement

	return p
}

// 对话框是否可见
func (p *Drawer) SetVisible(visible bool) *Drawer {
	p.Visible = visible

	return p
}

// 宽度
func (p *Drawer) SetWidth(width int) *Drawer {
	p.Width = width

	return p
}

// 设置 Modal 的 z-index
func (p *Drawer) SetZIndex(zIndex int) *Drawer {
	p.ZIndex = zIndex

	return p
}

// 弹窗行为
func (p *Drawer) SetActions(actions []interface{}) *Drawer {
	p.Actions = actions

	return p
}
