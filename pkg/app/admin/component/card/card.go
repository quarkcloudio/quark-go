package card

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Title            string      `json:"title"`
	SubTitle         string      `json:"subTitle"`
	Tip              string      `json:"tip"`
	Extra            interface{} `json:"extra"`
	Layout           string      `json:"layout"`
	Loading          bool        `json:"loading"`
	ColSpan          interface{} `json:"colSpan"`
	Gutter           interface{} `json:"gutter"`
	Split            string      `json:"split"`
	Bordered         bool        `json:"bordered"`
	Ghost            bool        `json:"ghost"`
	HeaderBordered   bool        `json:"headerBordered"`
	Collapsible      bool        `json:"collapsible"`
	DefaultCollapsed bool        `json:"defaultCollapsed"`
	Body             interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "card"
	p.Layout = "default"
	p.ColSpan = 24
	p.Gutter = 0
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置标题文字
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 设置二级标题文字
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle

	return p
}

// 标题右侧图标 hover 提示信息
func (p *Component) SetTip(tip string) *Component {
	p.Tip = tip

	return p
}

// 右上角自定义区域
func (p *Component) SetExtra(extra interface{}) *Component {
	p.Extra = extra

	return p
}

// 内容布局，支持垂直居中 default | center
func (p *Component) SetLayout(layout string) *Component {
	p.Layout = layout

	return p
}

// 加载中，支持自定义 loading 样式
func (p *Component) SetLoading(loading bool) *Component {
	p.Loading = loading

	return p
}

// 栅格布局宽度，24 栅格，支持指定宽度 px 或百分比, 支持响应式的对象写法 { xs: 8, sm: 16, md: 24}
func (p *Component) SetColSpan(colSpan interface{}) *Component {
	p.ColSpan = colSpan

	return p
}

// 栅格布局宽度，24 栅格，支持指定宽度 px 或百分比, 支持响应式的对象写法 { xs: 8, sm: 16, md: 24}
func (p *Component) SetGutter(gutter interface{}) *Component {
	p.Gutter = gutter

	return p
}

// 拆分卡片的方向,vertical | horizontal
func (p *Component) SetSplit(split string) *Component {
	p.Split = split

	return p
}

// 是否有边框
func (p *Component) SetBordered(bordered bool) *Component {
	p.Bordered = bordered

	return p
}

// 幽灵模式，即是否取消卡片内容区域的 padding 和 卡片的背景颜色。
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 页头是否有分割线
func (p *Component) SetHeaderBordered(headerBordered bool) *Component {
	p.HeaderBordered = headerBordered

	return p
}

// 页头是否有分割线
func (p *Component) SetCollapsible(collapsible bool) *Component {
	p.Collapsible = collapsible

	return p
}

// 默认折叠, 受控时无效
func (p *Component) SetDefaultCollapsed(defaultCollapsed bool) *Component {
	p.DefaultCollapsed = defaultCollapsed

	return p
}

// 卡牌内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "card"

	return p
}
