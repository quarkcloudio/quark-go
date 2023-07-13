package collapse

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type CollapseItem struct {
	component.Element
	Title         string      `json:"title"`
	Thumb         string      `json:"thumb"`
	Disabled      bool        `json:"disabled"`
	Open          bool        `json:"open"`
	ShowAnimation bool        `json:"showAnimation"`
	Border        bool        `json:"border"`
	TitleBorder   string      `json:"titleBorder"`
	ShowArrow     bool        `json:"showArrow"`
	Body          interface{} `json:"body"`
}

// 初始化
func (p *CollapseItem) Init() *CollapseItem {
	p.Component = "collapseItem"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *CollapseItem) SetStyle(style interface{}) *CollapseItem {
	p.Style = style

	return p
}

// 标题文字
func (p *CollapseItem) SetTitle(title string) *CollapseItem {
	p.Title = title

	return p
}

// 标题左侧缩略图
func (p *CollapseItem) SetThumb(thumb string) *CollapseItem {
	p.Thumb = thumb

	return p
}

// 是否禁用
func (p *CollapseItem) SetDisabled(disabled bool) *CollapseItem {
	p.Disabled = disabled

	return p
}

// 是否展开面板
func (p *CollapseItem) SetOpen(open bool) *CollapseItem {
	p.Open = open

	return p
}

// 开启动画
func (p *CollapseItem) SetShowAnimation(showAnimation bool) *CollapseItem {
	p.ShowAnimation = showAnimation

	return p
}

// 折叠面板内容分隔线
func (p *CollapseItem) SetBorder(border bool) *CollapseItem {
	p.Border = border

	return p
}

// 折叠面板标题分隔线可选值见下方 TitleBorder Params
func (p *CollapseItem) SetTitleBorder(titleBorder string) *CollapseItem {
	p.TitleBorder = titleBorder

	return p
}

// 是否显示右侧箭头
func (p *CollapseItem) SetShowArrow(showArrow bool) *CollapseItem {
	p.ShowArrow = showArrow

	return p
}

// 内容
func (p *CollapseItem) SetBody(body interface{}) *CollapseItem {
	p.Body = body

	return p
}

// 组件json序列化
func (p *CollapseItem) JsonSerialize() *CollapseItem {
	p.Component = "collapseItem"

	return p
}
