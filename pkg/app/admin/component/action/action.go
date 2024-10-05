package action

import (
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/component"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/drawer"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/modal"
)

type Component struct {
	component.Element
	Label        string      `json:"label"`
	Block        bool        `json:"block"`
	Danger       bool        `json:"danger"`
	Disabled     bool        `json:"disabled"`
	Ghost        bool        `json:"ghost"`
	Icon         string      `json:"icon"`
	Shape        string      `json:"shape"`
	Size         string      `json:"size"`
	Type         string      `json:"type"`
	ActionType   string      `json:"actionType"`
	SubmitForm   any         `json:"submitForm"`
	Href         string      `json:"href"`
	Target       string      `json:"target"`
	Modal        interface{} `json:"modal"`
	Drawer       interface{} `json:"drawer"`
	ConfirmTitle string      `json:"confirmTitle"`
	ConfirmText  string      `json:"confirmText"`
	ConfirmType  string      `json:"confirmType"`
	Api          string      `json:"api"`
	Reload       string      `json:"reload"`
	WithLoading  bool        `json:"withLoading"`
}

// 初始化组件
func New() *Component {

	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.Size = "default"
	p.Type = "default"

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

// 将按钮宽度调整为其父宽度的选项
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 设置危险按钮
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
// "icon-database", "icon-sever", "icon-mobile", "icon-tablet", "icon-redenvelope",
// "icon-book", "icon-filedone", "icon-reconciliation", "icon-file-exception",
// "icon-filesync", "icon-filesearch", "icon-solution", "icon-fileprotect",
// "icon-file-add", "icon-file-excel", "icon-file-exclamation", "icon-file-pdf",
// "icon-file-image", "icon-file-markdown", "icon-file-unknown", "icon-file-ppt",
// "icon-file-word", "icon-file", "icon-file-zip", "icon-file-text", "icon-file-copy",
// "icon-snippets", "icon-audit", "icon-diff", "icon-Batchfolding", "icon-securityscan",
// "icon-propertysafety", "icon-insurance", "icon-alert", "icon-delete", "icon-hourglass",
// "icon-bulb", "icon-experiment", "icon-bell", "icon-trophy", "icon-rest", "icon-USB",
// "icon-skin", "icon-home", "icon-bank", "icon-filter", "icon-funnelplot", "icon-like",
// "icon-unlike", "icon-unlock", "icon-lock", "icon-customerservice", "icon-flag",
// "icon-moneycollect", "icon-medicinebox", "icon-shop", "icon-rocket", "icon-shopping",
// "icon-folder", "icon-folder-open", "icon-folder-add", "icon-deploymentunit",
// "icon-accountbook", "icon-contacts", "icon-carryout", "icon-calendar-check",
// "icon-calendar", "icon-scan", "icon-select", "icon-boxplot", "icon-build", "icon-sliders",
// "icon-laptop", "icon-barcode", "icon-camera", "icon-cluster", "icon-gateway", "icon-car",
// "icon-printer", "icon-read", "icon-cloud-server", "icon-cloud-upload", "icon-cloud",
// "icon-cloud-download", "icon-cloud-sync", "icon-video", "icon-notification", "icon-sound",
// "icon-radarchart", "icon-qrcode", "icon-fund", "icon-image", "icon-mail", "icon-table",
// "icon-idcard", "icon-creditcard", "icon-heart", "icon-block", "icon-error", "icon-star",
// "icon-gold", "icon-heatmap", "icon-wifi", "icon-attachment", "icon-edit", "icon-key",
// "icon-api", "icon-disconnect", "icon-highlight", "icon-monitor", "icon-link", "icon-man",
// "icon-percentage", "icon-pushpin", "icon-phone", "icon-shake", "icon-tag", "icon-wrench",
// "icon-tags", "icon-scissor", "icon-mr", "icon-share", "icon-branches", "icon-fork", "icon-shrink",
// "icon-arrawsalt", "icon-verticalright", "icon-verticalleft", "icon-right", "icon-left",
// "icon-up", "icon-down", "icon-fullscreen", "icon-fullscreen-exit", "icon-doubleleft",
// "icon-doubleright", "icon-arrowright", "icon-arrowup", "icon-arrowleft", "icon-arrowdown",
// "icon-upload", "icon-colum-height", "icon-vertical-align-botto", "icon-vertical-align-middl",
// "icon-totop", "icon-vertical-align-top", "icon-download", "icon-sort-descending",
// "icon-sort-ascending", "icon-fall", "icon-swap", "icon-stock", "icon-rise", "icon-indent",
// "icon-outdent", "icon-menu", "icon-unorderedlist", "icon-orderedlist", "icon-align-right",
// "icon-align-center", "icon-align-left", "icon-pic-center", "icon-pic-right", "icon-pic-left",
// "icon-bold", "icon-font-colors", "icon-exclaimination", "icon-font-size", "icon-check-circle",
// "icon-infomation", "icon-CI", "icon-line-height", "icon-Dollar", "icon-strikethrough", "icon-compass",
// "icon-underline", "icon-close-circle", "icon-number", "icon-frown", "icon-italic", "icon-info-circle",
// "icon-code", "icon-left-circle", "icon-column-width", "icon-down-circle", "icon-check", "icon-EURO",
// "icon-ellipsis", "icon-copyright", "icon-dash", "icon-minus-circle", "icon-close", "icon-meh",
// "icon-enter", "icon-plus-circle", "icon-line", "icon-play-circle", "icon-minus", "icon-question-circle",
// "icon-question", "icon-Pound", "icon-rollback", "icon-right-circle", "icon-small-dash", "icon-smile",
// "icon-pause", "icon-trademark", "icon-bg-colors", "icon-time-circle", "icon-crown", "icon-timeout",
// "icon-drag", "icon-earth", "icon-desktop", "icon-YUAN", "icon-gift", "icon-up-circle", "icon-stop",
// "icon-warning-circle", "icon-fire", "icon-sync", "icon-thunderbolt", "icon-transaction",
// "icon-alipay", "icon-undo", "icon-taobao", "icon-redo", "icon-wechat-fill", "icon-reload",
// "icon-comment", "icon-reloadtime", "icon-login", "icon-message", "icon-clear", "icon-dashboard",
// "icon-issuesclose", "icon-poweroff", "icon-logout", "icon-piechart", "icon-setting",
// "icon-eye", "icon-location", "icon-edit-square", "icon-export", "icon-save", "icon-Import",
// "icon-appstore", "icon-close-square", "icon-down-square", "icon-layout", "icon-left-square",
// "icon-play-square", "icon-control", "icon-codelibrary", "icon-detail", "icon-minus-square",
// "icon-plus-square", "icon-right-square", "icon-project", "icon-wallet", "icon-up-square",
// "icon-calculator", "icon-interation", "icon-check-square", "icon-border", "icon-border-outer",
// "icon-border-top", "icon-border-bottom", "icon-border-left", "icon-border-right", "icon-border-inner",
// "icon-border-verticle", "icon-border-horizontal", "icon-radius-bottomleft", "icon-radius-bottomright",
// "icon-radius-upleft", "icon-radius-upright", "icon-radius-setting", "icon-adduser", "icon-deleteteam",
// "icon-deleteuser", "icon-addteam", "icon-user", "icon-team", "icon-areachart", "icon-linechart",
// "icon-barchart", "icon-pointmap", "icon-container", "icon-atom", "icon-zanwutupian", "icon-safetycertificate",
// "icon-password", "icon-article", "icon-page", "icon-plugin", "icon-admin", "icon-banner"
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = "icon-" + icon

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Component) SetActionType(actionType string) *Component {
	p.ActionType = actionType

	return p
}

// 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Component) SetSubmitForm(formKey string) *Component {
	p.SubmitForm = formKey

	return p
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func (p *Component) SetHref(href string) *Component {
	p.Href = href

	return p
}

// 相当于 a 链接的 target 属性，href 存在时生效
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

// 弹窗
func (p *Component) SetModal(callback interface{}) *Component {
	component := (&modal.Component{}).Init()
	getCallback := callback.(func(modal *modal.Component) interface{})

	p.Modal = getCallback(component)

	return p
}

// 抽屉
func (p *Component) SetDrawer(callback interface{}) *Component {
	component := (&drawer.Component{}).Init()
	getCallback := callback.(func(drawer *drawer.Component) interface{})

	p.Drawer = getCallback(component)

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
	p.ActionType = "ajax"

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
