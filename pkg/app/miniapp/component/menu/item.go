package menu

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Option struct {
	Text  string      `json:"text,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Item struct {
	component.Element
	Title     string    `json:"title,omitempty"`
	Options   []*Option `json:"options,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Cols      int       `json:"cols,omitempty"`
	Direction string    `json:"direction,omitempty"`
}

// 初始化组件
func NewItem() *Item {
	return (&Item{}).Init()
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "menuItem"
	p.SetKey("menuItem", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Item) SetStyle(style interface{}) *Item {
	p.Style = style

	return p
}

// 菜单项标题
func (p *Item) SetTitle(title string) *Item {
	p.Title = title
	return p
}

// 选项数组
func (p *Item) SetOptions(options []*Option) *Item {
	p.Options = options
	return p
}

// 是否禁用菜单
func (p *Item) SetDisabled(disabled bool) *Item {
	p.Disabled = disabled
	return p
}

// 可以设置一行展示多少列 options
func (p *Item) SetCols(cols int) *Item {
	p.Cols = cols
	return p
}

// 菜单展开方向，可选值为 up
func (p *Item) SetDirection(direction string) *Item {
	p.Direction = direction
	return p
}

// 组件json序列化
func (p *Item) JsonSerialize() *Item {
	p.Component = "menuItem"

	return p
}
