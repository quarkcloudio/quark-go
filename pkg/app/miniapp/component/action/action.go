package action

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Label    string `json:"label,omitempty"`
	Type     string `json:"type,omitempty"`
	FormType string `json:"formType,omitempty"`
	Size     string `json:"size,omitempty"`
	Shape    string `json:"shape,omitempty"`
	Color    string `json:"color,omitempty"`
	Plain    bool   `json:"plain,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	Block    bool   `json:"block,omitempty"`
	Loading  bool   `json:"loading,omitempty"`

	OpenType               string   `json:"openType,omitempty"`
	HoverClass             string   `json:"hoverClass,omitempty"`
	HoverStartTime         int      `json:"hoverStartTime,omitempty"`
	HoverStayTime          int      `json:"hoverStayTime,omitempty"`
	AppParameter           string   `json:"appParameter,omitempty"`
	Scope                  string   `json:"scope,omitempty"`
	HoverStopPropagation   bool     `json:"hoverStopPropagation,omitempty"`
	Lang                   string   `json:"lang,omitempty"`
	SessionFrom            string   `json:"sessionFrom,omitempty"`
	SendMessageTitle       string   `json:"sendMessageTitle,omitempty"`
	SendMessagePath        string   `json:"sendMessagePath,omitempty"`
	SendMessageImg         string   `json:"sendMessageImg,omitempty"`
	ShowMessageCard        bool     `json:"showMessageCard,omitempty"`
	PublicId               string   `json:"publicId,omitempty"`
	TemplateId             []string `json:"templateId,omitempty"`
	SubscribeId            string   `json:"subscribeId,omitempty"`
	GroupId                string   `json:"groupId,omitempty"`
	GuildId                string   `json:"guildId,omitempty"`
	ShareType              string   `json:"shareType,omitempty"`
	ShareMode              string   `json:"shareMode,omitempty"`
	AriaLabel              string   `json:"ariaLabel,omitempty"`
	OpenId                 string   `json:"openId,omitempty"`
	ShareMessageFriendInfo string   `json:"shareMessageFriendInfo,omitempty"`
	ShareMessageTitle      string   `json:"shareMessageTitle,omitempty"`
	ShareMessageImg        string   `json:"shareMessageImg,omitempty"`

	ActionType   string      `json:"actionType,omitempty"`
	Href         string      `json:"href,omitempty"`
	Target       string      `json:"target,omitempty"`
	Popup        interface{} `json:"popup,omitempty"`
	Drawer       interface{} `json:"drawer,omitempty"`
	ConfirmTitle string      `json:"confirmTitle,omitempty"`
	ConfirmText  string      `json:"confirmText,omitempty"`
	ConfirmType  string      `json:"confirmType,omitempty"`
	Api          string      `json:"api,omitempty"`
	ApiType      string      `json:"apiType,omitempty"`
	Reload       string      `json:"reload,omitempty"`
	WithLoading  bool        `json:"withLoading,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.Type = "default"
	p.ApiType = "GET"

	p.SetKey("action", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 类型，可选值为 primary info warning danger success default
func (p *Component) SetType(actionType string) *Component {
	p.Type = actionType

	return p
}

// 表单类型，可选值 button submit reset
func (p *Component) SetFormType(formType string) *Component {
	p.FormType = formType

	return p
}

// 尺寸，可选值为 large small mini normal
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 形状，可选值为 square round
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 按钮颜色，支持传入 linear-gradient 渐变色
func (p *Component) SetColor(color string) *Component {
	p.Color = color

	return p
}

// 是否为朴素按钮
func (p *Component) SetPlain(plain bool) *Component {
	p.Plain = plain

	return p
}

// 是否禁用按钮
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 是否为块级元素
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 按钮 loading 状态
func (p *Component) SetLoading(loading bool) *Component {
	p.Loading = loading

	return p
}

// 微信开放能力
func (p *Component) SetOpenType(openType string) *Component {
	p.OpenType = openType

	return p
}

// 指定按下去的样式类。当 hover-class="none" 时，没有点击态效果
func (p *Component) SetHoverClass(hoverClass string) *Component {
	p.HoverClass = hoverClass

	return p
}

// 指定是否阻止本节点的祖先节点出现点击态
func (p *Component) SetHoverStopPropagation(hoverStopPropagation bool) *Component {
	p.HoverStopPropagation = hoverStopPropagation

	return p
}

// 按住后多久出现点击态，单位毫秒
func (p *Component) SetHoverStartTime(hoverStartTime int) *Component {
	p.HoverStartTime = hoverStartTime

	return p
}

// 手指松开后点击态保留时间，单位毫秒
func (p *Component) SetHoverStayTime(hoverStayTime int) *Component {
	p.HoverStayTime = hoverStayTime

	return p
}

// 指定返回用户信息的语言，zh_CN 简体中文，zh_TW 繁体中文，en 英文。生效时机: open-type="getUserInfo"
func (p *Component) SetLang(lang string) *Component {
	p.Lang = lang

	return p
}

// 微信小程序：会话来源 生效时机：open-type="contact"
func (p *Component) SetSessionFrom(sessionFrom string) *Component {
	p.SessionFrom = sessionFrom

	return p
}

// 微信小程序：会话内消息卡片标题 生效时机：open-type="contact"
func (p *Component) SetSendMessageTitle(sendMessageTitle string) *Component {
	p.SendMessageTitle = sendMessageTitle

	return p
}

// 微信小程序：会话内消息卡片点击跳转小程序路径，open-type="contact"时有效
func (p *Component) SetSendMessagePath(sendMessagePath string) *Component {
	p.SendMessagePath = sendMessagePath

	return p
}

// 微信小程序：会话内消息卡片图片，open-type="contact"时有效
func (p *Component) SetSendMessageImg(sendMessageImg string) *Component {
	p.SendMessageImg = sendMessageImg

	return p
}

// 微信小程序、QQ小程序：打开 APP 时，向 APP 传递的参数，open-type=launchApp时有效
func (p *Component) SetAppParameter(appParameter string) *Component {
	p.AppParameter = appParameter

	return p
}

// 支付宝小程序 scope 生效时机：open-type="getAuthorize"
func (p *Component) SetScope(scope string) *Component {
	p.Scope = scope

	return p
}

// 微信小程序：是否显示会话内消息卡片，设置此参数为 true，用户进入客服会话会在右下角显示"可能要发送的小程序"提示，用户点击后可以快速发送小程序消息，open-type="contact"时有效
func (p *Component) SetShowMessageCard(showMessageCard bool) *Component {
	p.ShowMessageCard = showMessageCard

	return p
}

// 生活号 id，必须是当前小程序同主体且已关联的生活号，open-type="lifestyle" 时有效。
func (p *Component) SetPublicId(publicId string) *Component {
	p.PublicId = publicId

	return p
}

// 发送订阅类模板消息所用的模板库标题 ID ，可通过 getTemplateLibraryList 获取当参数类型为 Array 时，可传递 1~3 个模板库标题 ID
func (p *Component) SetTemplateId(templateId []string) *Component {
	p.TemplateId = templateId

	return p
}

// 发送订阅类模板消息时所使用的唯一标识符，内容由开发者自定义，用来标识订阅场景
// 注意：同一用户在同一 subscribe-id 下的多次授权不累积下发权限，只能下发一条。若要订阅多条，需要不同 subscribe-id
func (p *Component) SetSubscribeId(subscribeId string) *Component {
	p.SubscribeId = subscribeId

	return p
}

// QQ小程序：打开群资料卡时，传递的群号
func (p *Component) SetGroupId(groupId string) *Component {
	p.GroupId = groupId

	return p
}

// QQ小程序：打开频道页面时，传递的频道号
func (p *Component) SetGuildId(guildId string) *Component {
	p.GuildId = guildId

	return p
}

// 分享类型集合，请参考下面share-type有效值说明。share-type后续将不再维护，请更新为share-mode
func (p *Component) SetShareType(shareType string) *Component {
	p.ShareType = shareType

	return p
}

// 分享类型集合，请参考下面share-mode有效值说明
func (p *Component) SetShareMode(shareMode string) *Component {
	p.ShareMode = shareMode

	return p
}

// 无障碍访问，（属性）元素的额外描述
func (p *Component) SetAriaLabel(ariaLabel string) *Component {
	p.AriaLabel = ariaLabel

	return p
}

// 添加好友时，对方的 openid
func (p *Component) SetOpenId(openId string) *Component {
	p.OpenId = openId

	return p
}

// 发送对象的 FriendInfo
func (p *Component) SetShareMessageFriendInfo(shareMessageFriendInfo string) *Component {
	p.ShareMessageFriendInfo = shareMessageFriendInfo

	return p
}

// 转发标题，不传则默认使用当前小程序的昵称。 FriendInfo
func (p *Component) SetShareMessageTitle(shareMessageTitle string) *Component {
	p.ShareMessageTitle = shareMessageTitle

	return p
}

// 转发显示图片的链接，可以是网络图片路径（仅 QQ CDN 域名路径）或本地图片文件路径或相对代码包根目录的图片文件路径。显示图片长宽比是 5:4FriendInfo
func (p *Component) SetShareMessageImg(shareMessageImg string) *Component {
	p.ShareMessageImg = shareMessageImg

	return p
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：openType、ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Component) SetActionType(actionType string) *Component {
	p.ActionType = actionType

	return p
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func (p *Component) SetHref(href string) *Component {
	p.Href = href

	return p
}

// 包含switchTab|reLaunch|redirectTo|navigateTo属性；href 存在时生效。
func (p *Component) SetTarget(target string) *Component {
	p.Target = target

	return p
}

// 设置跳转链接
func (p *Component) SetLink(href string, target string) *Component {
	p.SetHref(href)
	p.SetTarget(target)
	p.ActionType = "link"

	return p
}

// 弹出层
func (p *Component) SetPopup(callback interface{}) *Component {
	popup := (&Popup{}).Init()
	getCallback := callback.(func(popup *Popup) interface{})

	p.Popup = getCallback(popup)

	return p
}

// 抽屉
func (p *Component) SetDrawer(callback interface{}) *Component {
	drawer := (&Drawer{}).Init()
	getCallback := callback.(func(drawer *Drawer) interface{})

	p.Drawer = getCallback(drawer)

	return p
}

// 设置行为前的确认操作
func (p *Component) SetWithConfirm(title string, text string, confirmType string) *Component {
	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

// 执行行为的接口链接
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	p.OpenType = "ajax"

	return p
}

// 执行行为的方法，GET/POST
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

// 执行成功后刷新的组件
func (p *Component) SetReload(reload string) *Component {
	p.Reload = reload

	return p
}

// 是否具有loading
func (p *Component) SetWithLoading(loading bool) *Component {
	p.WithLoading = loading

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "action"

	return p
}
