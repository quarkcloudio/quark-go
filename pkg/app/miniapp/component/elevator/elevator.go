package elevator

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Height      int         `json:"height"`
	AcceptKey   string      `json:"acceptKey"`
	IndexList   interface{} `json:"indexList"`
	IsSticky    bool        `json:"isSticky"`
	SpaceHeight int         `json:"spaceHeight"`
	TitleHeight int         `json:"titleHeight"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "elevator"
	p.SetKey("elevator", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 电梯区域的高度
func (p *Component) SetHeight(height int) *Component {
	p.Height = height
	return p
}

// 索引 key 值
func (p *Component) SetAcceptKey(acceptKey string) *Component {
	p.AcceptKey = acceptKey
	return p
}

// 索引列表，Array（ item 需包含 id、name 属性, name 支持传入 html 结构）
func (p *Component) SetIndexList(indexList interface{}) *Component {
	p.IndexList = indexList
	return p
}

// 索引是否吸顶
func (p *Component) SetIsSticky(isSticky bool) *Component {
	p.IsSticky = isSticky
	return p
}

// 右侧锚点的上下间距
func (p *Component) SetSpaceHeight(spaceHeight int) *Component {
	p.SpaceHeight = spaceHeight
	return p
}

// 左侧索引的高度
func (p *Component) SetTitleHeight(titleHeight int) *Component {
	p.TitleHeight = titleHeight
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "elevator"

	return p
}
