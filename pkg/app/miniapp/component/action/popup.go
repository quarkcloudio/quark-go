package action

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Popup struct {
	component.Element
	Animation           bool        `json:"animation"`
	Type                string      `json:"type"`
	IsMaskClick         bool        `json:"isMaskClick"`
	MaskBackgroundColor string      `json:"maskBackgroundColor"`
	BackgroundColor     string      `json:"backgroundColor"`
	SafeArea            bool        `json:"safeArea"`
	MessageType         string      `json:"messageType"`
	MessageDuration     int         `json:"messageDuration"`
	DialogType          string      `json:"dialogType"`
	DialogMode          string      `json:"dialogMode"`
	DialogTitle         string      `json:"dialogTitle"`
	DialogConfirmText   string      `json:"dialogConfirmText"`
	DialogCancelText    string      `json:"dialogCancelText"`
	DialogValue         interface{} `json:"dialogValue"`
	DialogPlaceholder   string      `json:"dialogPlaceholder"`
	DialogBeforeClose   bool        `json:"dialogBeforeClose"`
	ShareTitle          string      `json:"shareTitle"`
	ShareBeforeClose    bool        `json:"ShareBeforeClose"`
	Body                interface{} `json:"body"`
}

// 初始化
func (p *Popup) Init() *Popup {
	p.Component = "popup"
	p.SetKey("popup", component.DEFAULT_CRYPT)
	p.Animation = true
	p.Type = "center"
	p.IsMaskClick = true
	p.MaskBackgroundColor = "rgba(0,0,0,0.4)"
	p.BackgroundColor = "#fff"
	p.SafeArea = true
	p.MessageType = "success"
	p.MessageDuration = 3000
	p.DialogType = "success"
	p.DialogMode = "base"

	return p
}

// Set style.
func (p *Popup) SetStyle(style map[string]interface{}) *Popup {
	p.Style = style

	return p
}

// 是否开启动画
func (p *Popup) SetAnimation(animation bool) *Popup {
	p.Animation = animation

	return p
}

// 弹出方式，top:顶部弹出,center:居中弹出,bottom:底部弹出,left:左侧弹出,right:右侧弹出,message:消息提示,dialog:对话框,share:底部弹出分享示例
func (p *Popup) SetType(popupType string) *Popup {
	p.Type = popupType

	return p
}

// 容器控件里面的内容
func (p *Popup) SetBody(body interface{}) *Popup {
	p.Body = body

	return p
}

// 蒙版点击是否关闭弹窗
func (p *Popup) SetIsMaskClick(isMaskClick bool) *Popup {
	p.IsMaskClick = isMaskClick

	return p
}

// 蒙版颜色，建议使用 rgba 颜色值,rgba(0,0,0,0.4)
func (p *Popup) SetMaskBackgroundColor(maskBackgroundColor string) *Popup {
	p.MaskBackgroundColor = maskBackgroundColor

	return p
}

// 主窗口背景色
func (p *Popup) SetBackgroundColor(backgroundColor string) *Popup {
	p.BackgroundColor = backgroundColor

	return p
}

// 是否适配底部安全区
func (p *Popup) SetSafeArea(safeArea bool) *Popup {
	p.SafeArea = safeArea

	return p
}

// 消息提示主题,success成功,warn警告,error失败,info消息
func (p *Popup) SetMessageType(messageType string) *Popup {
	p.MessageType = messageType

	return p
}

// 消息显示时间，超过显示时间组件自动关闭，设置为0 将不会关闭，需手动调用 close 方法关闭
func (p *Popup) SetMessageDuration(messageDuration int) *Popup {
	p.MessageDuration = messageDuration

	return p
}

// 对话框标题主题，可选值： success/warn/info/error
func (p *Popup) SetDialogType(dialogType string) *Popup {
	p.DialogType = dialogType

	return p
}

// 对话框模式，可选值：base（提示对话框）/input（可输入对话框）
func (p *Popup) SetDialogMode(dialogMode string) *Popup {
	p.DialogMode = dialogMode

	return p
}

// 对话框标题
func (p *Popup) SetDialogTitle(dialogTitle string) *Popup {
	p.DialogTitle = dialogTitle

	return p
}

// 定义确定按钮文本
func (p *Popup) SetDialogConfirmText(dialogConfirmText string) *Popup {
	p.DialogConfirmText = dialogConfirmText

	return p
}

// 定义取消按钮文本
func (p *Popup) SetDialogCancelText(dialogCancelText string) *Popup {
	p.DialogCancelText = dialogCancelText

	return p
}

// 输入框默认值，input模式下生效
func (p *Popup) SetDialogValue(dialogValue interface{}) *Popup {
	p.DialogValue = dialogValue

	return p
}

// 输入框提示文字，input模式下生效
func (p *Popup) SetDialogPlaceholder(dialogPlaceholder string) *Popup {
	p.DialogPlaceholder = dialogPlaceholder

	return p
}

// 是否拦截按钮事件，如为true，则不会关闭对话框，关闭需要手动执行 uni-popup 的 close 方法
func (p *Popup) SetDialogBeforeClose(dialogBeforeClose bool) *Popup {
	p.DialogBeforeClose = dialogBeforeClose

	return p
}

// 分享弹窗标题
func (p *Popup) SetShareTitle(shareTitle string) *Popup {
	p.ShareTitle = shareTitle

	return p
}

// 是否拦截按钮事件，如为true，则不会关闭对话框，关闭需要手动执行 uni-popup 的 close 方法
func (p *Popup) SetShareBeforeClose(shareBeforeClose bool) *Popup {
	p.ShareBeforeClose = shareBeforeClose

	return p
}
