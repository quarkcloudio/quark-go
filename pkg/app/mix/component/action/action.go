package action

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Label                string      `json:"label"`
	Size                 string      `json:"size"`
	Type                 string      `json:"type"`
	Plain                bool        `json:"plain"`
	Disabled             bool        `json:"disabled"`
	Loading              bool        `json:"loading"`
	FormType             string      `json:"formType"`
	OpenType             string      `json:"openType"`
	HoverClass           string      `json:"hoverClass"`
	HoverStartTime       int         `json:"hoverStartTime"`
	HoverStayTime        int         `json:"hoverStayTime"`
	AppParameter         string      `json:"appParameter"`
	HoverStopPropagation bool        `json:"hoverStopPropagation"`
	Lang                 string      `json:"lang"`
	SessionFrom          string      `json:"sessionFrom"`
	SendMessageTitle     string      `json:"sendMessageTitle"`
	SendMessagePath      string      `json:"sendMessagePath"`
	SendMessageImg       string      `json:"sendMessageImg"`
	ShowMessageCard      bool        `json:"showMessageCard"`
	GroupId              string      `json:"groupId"`
	GuildId              string      `json:"guildId"`
	PublicId             string      `json:"publicId"`
	Popup                interface{} `json:"popup"`
	Drawer               interface{} `json:"drawer"`
	ConfirmTitle         string      `json:"confirmTitle"`
	ConfirmText          string      `json:"confirmText"`
	ConfirmType          string      `json:"confirmType"`
	Api                  string      `json:"api"`
	ApiType              string      `json:"apiType"`
	Reload               string      `json:"reload"`
	WithLoading          bool        `json:"withLoading"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.Size = "default"
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

// 按钮的大小
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 按钮的样式类型 primary | default | warn
func (p *Component) SetType(actionType string) *Component {
	p.Type = actionType

	return p
}

// 按钮是否镂空，背景色透明
func (p *Component) SetPlain(plain bool) *Component {
	p.Plain = plain

	return p
}

// 是否禁用
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 名称前是否带 loading 图标
func (p *Component) SetLoading(loading bool) *Component {
	p.Loading = loading

	return p
}

// 用于 <form> 组件，点击分别会触发 <form> 组件的 submit/reset 事件
func (p *Component) SetFormType(formType string) *Component {
	p.FormType = formType

	return p
}

// 开放能力
// feedback	打开“意见反馈”页面，用户可提交反馈内容并上传日志	App、微信小程序、QQ小程序
// share	触发用户转发	微信小程序、百度小程序、支付宝小程序、字节跳动小程序、飞书小程序、QQ小程序、快手小程序、京东小程序、360小程序
// getUserInfo	获取用户信息，可以从@getuserinfo回调中获取到用户信息	微信小程序、百度小程序、QQ小程序、快手小程序、京东小程序、360小程序
// contact	打开客服会话，如果用户在会话中点击消息卡片后返回应用，可以从 @contact 回调中获得具体信息	微信小程序、百度小程序、快手小程序
// getPhoneNumber	获取用户手机号，可以从@getphonenumber回调中获取到用户信息	微信小程序、百度小程序、字节跳动小程序、支付宝小程序、快手小程序、京东小程序。App平台另见一键登陆(opens new window)
// launchApp	小程序中打开APP，可以通过app-parameter属性设定向APP传的参数	微信小程序 (opens new window)、QQ小程序 (opens new window)、快手小程序、京东小程序
// openSetting	打开授权设置页	微信小程序、QQ小程序、百度小程序、快手小程序、京东小程序、360小程序
// chooseAvatar	获取用户头像，可以从@chooseavatar回调中获取到头像信息	微信小程序2.21.2版本+
// uploadDouyinVideo	发布抖音视频	字节小程序2.65.0版本+
// getAuthorize	支持小程序授权	支付宝小程序
// lifestyle	关注生活号	支付宝小程序
// contactShare	分享到通讯录好友	支付宝小程序基础库1.11.0版本+
// openGroupProfile	呼起QQ群资料卡页面，可以通过group-id属性设定需要打开的群资料卡的群号，同时manifest.json中必须配置groupIdList	QQ小程序基础库1.4.7版本+
// openGuildProfile	呼起频道页面，可以通过guild-id属性设定需要打开的频道ID	QQ小程序基础库1.46.8版本+
// openPublicProfile	打开公众号资料卡，可以通过public-id属性设定需要打开的公众号资料卡的号码，同时manifest.json中必须配置publicIdList	QQ小程序基础库1.12.0版本+
// shareMessageToFriend	在自定义开放数据域组件中,向指定好友发起分享据	QQ小程序基础库1.17.0版本+
// addFriend	添加好友， 对方需要通过该小程序进行授权，允许被加好友后才能调用成功用户授权	QQ小程序
// addColorSign	添加彩签，点击后添加状态有用户提示，无回调	QQ小程序基础库1.10.0版本+
// addGroupApp	添加群应用（只有管理员或群主有权操作），添加后给button绑定@addgroupapp事件接收回调数据	QQ小程序基础库1.16.0版本+
// addToFavorites	收藏当前页面，点击按钮后会触发Page.onAddToFavorites方法	QQ小程序基础库1.19.0版本+
// chooseAddress	选择用户收货地址，可以从@chooseaddress回调中获取到用户选择的地址信息	百度小程序3.160.3版本+
// chooseInvoiceTitle	选择用户发票抬头，可以从@chooseinvoicetitle回调中获取到用户选择发票抬头信息	百度小程序3.160.3版本+
// login	登录，可以从@login回调中确认是否登录成功	百度小程序3.230.1版本+
// subscribe	订阅类模板消息，需要用户授权才可发送	百度小程序
// favorite	触发用户收藏	快手小程序
// watchLater	触发用户稍后再看	快手小程序
// openProfile	触发打开用户主页	快手小程序
// ajax
// submit
// popup
// drawer
// link
func (p *Component) SetOpenType(openType string) *Component {
	p.OpenType = openType

	return p
}

//指定按钮按下去的样式类。当 hover-class="none" 时，没有点击态效果
func (p *Component) SetHoverClass(hoverClass string) *Component {
	p.HoverClass = hoverClass

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

// 微信小程序、QQ小程序：打开 APP 时，向 APP 传递的参数，open-type=launchApp时有效
func (p *Component) SetAppParameter(appParameter string) *Component {
	p.AppParameter = appParameter

	return p
}

// 微信小程序：指定是否阻止本节点的祖先节点出现点击态
func (p *Component) SetHoverStopPropagation(hoverStopPropagation bool) *Component {
	p.HoverStopPropagation = hoverStopPropagation

	return p
}

// 微信小程序：指定返回用户信息的语言，zh_CN 简体中文，zh_TW 繁体中文，en 英文。
func (p *Component) SetLang(lang string) *Component {
	p.Lang = lang

	return p
}

// 微信小程序：会话来源，open-type="contact"时有效
func (p *Component) SetSessionFrom(sessionFrom string) *Component {
	p.SessionFrom = sessionFrom

	return p
}

// 微信小程序：会话内消息卡片标题，open-type="contact"时有效
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

// 微信小程序：是否显示会话内消息卡片，设置此参数为 true，用户进入客服会话会在右下角显示"可能要发送的小程序"提示，用户点击后可以快速发送小程序消息，open-type="contact"时有效
func (p *Component) SetShowMessageCard(showMessageCard bool) *Component {
	p.ShowMessageCard = showMessageCard

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

// QQ小程序：打开公众号资料卡时，传递的号码
func (p *Component) SetPublicId(publicId string) *Component {
	p.PublicId = publicId

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

//  执行行为的接口链接
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	p.OpenType = "ajax"

	return p
}

//  执行行为的方法，GET/POST
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

//  执行成功后刷新的组件
func (p *Component) SetReload(reload string) *Component {
	p.Reload = reload

	return p
}

//  是否具有loading
func (p *Component) SetWithLoading(loading bool) *Component {
	p.WithLoading = loading

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "action"

	return p
}
