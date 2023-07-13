package pagecontainer

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Content            interface{}         `json:"content"`
	ExtraContent       interface{}         `json:"extraContent"`
	TabList            []map[string]string `json:"tabList"`
	TabActiveKey       string              `json:"tabActiveKey"`
	TabBarExtraContent interface{}         `json:"tabBarExtraContent"`
	Header             interface{}         `json:"header"`
	Ghost              bool                `json:"ghost"`
	FixedHeader        bool                `json:"fixedHeader"`
	AffixProps         interface{}         `json:"affixProps"`
	Footer             interface{}         `json:"footer"`
	Body               interface{}         `json:"body"`
	WaterMarkProps     interface{}         `json:"waterMarkProps"`
	TabProps           interface{}         `json:"tabProps"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取PageHeader
func NewPageHeader() *PageHeader {

	return (&PageHeader{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "pageContainer"

	p.SetKey("pageContainer", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 内容区
func (p *Component) SetContent(content interface{}) *Component {
	p.Content = content
	return p
}

// 额外内容区，位于 content 的右侧
func (p *Component) SetExtraContent(extraContent interface{}) *Component {
	p.ExtraContent = extraContent
	return p
}

// tab 标题列表
func (p *Component) SetTabList(tabList []map[string]string) *Component {
	p.TabList = tabList
	return p
}

// 当前高亮的 tab 项
func (p *Component) SetTabActiveKey(tabActiveKey string) *Component {
	p.TabActiveKey = tabActiveKey
	return p
}

// tab bar 上额外的元素
func (p *Component) SetTabBarExtraContent(tabBarExtraContent interface{}) *Component {
	p.TabBarExtraContent = tabBarExtraContent
	return p
}

// PageHeader 的所有属性
func (p *Component) SetHeader(header interface{}) *Component {
	p.Header = header
	return p
}

// 配置头部区域的背景颜色为透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost
	return p
}

// 固定 pageHeader 的内容到顶部，如果页面内容较少，最好不要使用，会有严重的遮挡问题
func (p *Component) SetFixedHeader(fixedHeader bool) *Component {
	p.FixedHeader = fixedHeader
	return p
}

// 固钉的配置，与 antd 完全相同
func (p *Component) SetAffixProps(affixProps interface{}) *Component {
	p.AffixProps = affixProps
	return p
}

// 悬浮在底部的操作栏，传入一个数组，会自动加空格
func (p *Component) SetFooter(footer interface{}) *Component {
	p.Footer = footer
	return p
}

// 容器控件里面的内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body
	return p
}

// 配置水印，Layout 会透传给 PageContainer，但是以 PageContainer 的配置优先
func (p *Component) SetWaterMarkProps(waterMarkProps interface{}) *Component {
	p.WaterMarkProps = waterMarkProps
	return p
}

// 配置水印，Layout 会透传给 PageContainer，但是以 PageContainer 的配置优先
func (p *Component) SetTabProps(tabProps interface{}) *Component {
	p.TabProps = tabProps
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "pageContainer"

	return p
}
