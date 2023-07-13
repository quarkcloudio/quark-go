package space

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Align     string      `json:"align"`
	Direction string      `json:"direction"`
	Size      string      `json:"size"`
	Split     string      `json:"split"`
	Wrap      bool        `json:"wrap"`
	Body      interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "space"
	p.Size = "small"

	p.SetKey("space", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

/**
 * 对齐方式
 *
 * @param  string  align
 * @return p
 */
func (p *Component) SetAlign(align string) *Component {
	p.Align = align

	return p
}

/**
 * 间距方向
 *
 * @param  string  direction
 * @return p
 */
func (p *Component) SetDirection(direction string) *Component {
	p.Direction = direction

	return p
}

/**
 * 间距大小
 *
 * @param  string  size
 * @return p
 */
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

/**
 * 拆分卡片的方向,vertical | horizontal
 *
 * @param  string  split
 * @return p
 */
func (p *Component) SetSplit(split string) *Component {

	p.Split = split

	return p
}

/**
 * 是否自动换行，仅在 horizontal 时有效
 *
 * @param  bool  wrap
 * @return p
 */
func (p *Component) SetWrap(wrap bool) *Component {
	p.Wrap = wrap

	return p
}

/**
 * 容器控件里面的内容
 *
 * @param  string|array  body
 * @return p
 */
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}
