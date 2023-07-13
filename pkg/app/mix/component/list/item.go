package list

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type ListItem struct {
	component.Element
	Title         string      `json:"title"`
	Note          string      `json:"note"`
	Ellipsis      int         `json:"ellipsis"`
	Thumb         string      `json:"thumb"`
	ThumbSize     string      `json:"thumbSize"`
	ShowBadge     bool        `json:"showBadge"`
	BadgeText     string      `json:"badgeText"`
	BadgeType     string      `json:"badgeType"`
	BadgeStyle    interface{} `json:"badgeStyle"`
	RightText     string      `json:"rightText"`
	Disabled      bool        `json:"disabled"`
	ShowArrow     bool        `json:"showArrow"`
	Link          string      `json:"link"`
	To            string      `json:"to"`
	Clickable     bool        `json:"clickable"`
	ShowSwitch    bool        `json:"showSwitch"`
	SwitchChecked bool        `json:"switchChecked"`
	ShowExtraIcon bool        `json:"showExtraIcon"`
	ExtraIcon     interface{} `json:"extraIcon"`
	Direction     string      `json:"direction"`
}

// 初始化
func (p *ListItem) Init() *ListItem {
	p.Component = "listItem"
	p.Ellipsis = 0

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *ListItem) SetStyle(style interface{}) *ListItem {
	p.Style = style

	return p
}

// 标题
func (p *ListItem) SetTitle(title string) *ListItem {
	p.Title = title

	return p
}

// 描述
func (p *ListItem) SetNote(note string) *ListItem {
	p.Note = note

	return p
}

// title 是否溢出隐藏，可选值，0:默认; 1:显示一行; 2:显示两行;
func (p *ListItem) SetEllipsis(ellipsis int) *ListItem {
	p.Ellipsis = ellipsis

	return p
}

// 左侧缩略图，若thumb有值，则不会显示扩展图标
func (p *ListItem) SetThumb(thumb string) *ListItem {
	p.Thumb = thumb

	return p
}

// 略缩图尺寸，可选值，lg:大图; medium:一般; sm:小图;
func (p *ListItem) SetThumbSize(thumbSize string) *ListItem {
	p.ThumbSize = thumbSize

	return p
}

// 是否显示数字角标
func (p *ListItem) SetShowBadge(showBadge bool) *ListItem {
	p.ShowBadge = showBadge

	return p
}

// 数字角标内容
func (p *ListItem) SetBadgeText(badgeText string) *ListItem {
	p.BadgeText = badgeText

	return p
}

// 数字角标类型，参考uni-icons
func (p *ListItem) SetBadgeType(badgeType string) *ListItem {
	p.BadgeType = badgeType

	return p
}

// 数字角标样式，使用uni-badge的custom-style参数
func (p *ListItem) SetBadgeStyle(badgeStyle interface{}) *ListItem {
	p.BadgeStyle = badgeStyle

	return p
}

// 右侧文字内容
func (p *ListItem) SetRightText(rightText string) *ListItem {
	p.RightText = rightText

	return p
}

// 是否禁用
func (p *ListItem) SetDisabled(disabled bool) *ListItem {
	p.Disabled = disabled

	return p
}

// 是否显示箭头图标
func (p *ListItem) SetShowArrow(showArrow bool) *ListItem {
	p.ShowArrow = showArrow

	return p
}

// 新页面跳转方式，可选值见下表
func (p *ListItem) SetLink(link string) *ListItem {
	p.Link = link

	return p
}

// 新页面跳转地址，如填写此属性，click 会返回页面是否跳转成功
func (p *ListItem) SetTo(to string) *ListItem {
	p.To = to

	return p
}

// 是否开启点击反馈
func (p *ListItem) SetClickable(clickable bool) *ListItem {
	p.Clickable = clickable

	return p
}

// 是否显示Switch
func (p *ListItem) SetShowSwitch(showSwitch bool) *ListItem {
	p.ShowSwitch = showSwitch

	return p
}

// Switch是否被选中
func (p *ListItem) SetSwitchChecked(switchChecked bool) *ListItem {
	p.SwitchChecked = switchChecked

	return p
}

// 左侧是否显示扩展图标
func (p *ListItem) SetShowExtraIcon(showExtraIcon bool) *ListItem {
	p.ShowExtraIcon = showExtraIcon

	return p
}

// 扩展图标参数，格式为 {color: '#4cd964',size: '22',type: 'spinner'}，参考 uni-icons
func (p *ListItem) SetExtraIcon(extraIcon interface{}) *ListItem {
	p.ExtraIcon = extraIcon

	return p
}

// 排版方向，可选值，row:水平排列; column:垂直排列; 3个插槽是水平排还是垂直排，也受此属性控制
func (p *ListItem) SetDirection(direction string) *ListItem {
	p.Direction = direction

	return p
}

// 组件json序列化
func (p *ListItem) JsonSerialize() *ListItem {
	p.Component = "listItem"

	return p
}
