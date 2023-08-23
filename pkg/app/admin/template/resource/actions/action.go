package actions

import (
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Action struct {
	Name       string   `json:"name"`       // 设置按钮文字；支持js表达式例如：<%= (status==1 ? '禁用' : '启用') %>，行为在表格行时，可以使用当前行的任意字段值，示例中status即为当前行的“状态”字段
	Reload     string   `json:"reload"`     // 执行成功后刷新的组件
	ApiParams  []string `json:"apiParams"`  // 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	Api        string   `json:"api"`        // 行为接口
	ActionType string   `json:"actionType"` // 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
	SubmitForm string   `json:"submitForm"` // 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
	// 设置按钮的图标组件：
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
	Icon                  string      `json:"icon"`
	Type                  string      `json:"type"`                  // 设置按钮类型，primary | ghost | dashed | link | text | default
	Size                  string      `json:"size"`                  // 设置按钮大小,large | middle | small | default
	WithLoading           bool        `json:"withLoading"`           // 是否具有loading，当action 的作用类型为ajax,submit时有效
	Fields                interface{} `json:"fields"`                // 行为表单字段
	ConfirmTitle          string      `json:"confirmTitle"`          // 确认标题
	ConfirmText           string      `json:"confirmText"`           // 确认文字描述
	ConfirmType           string      `json:"confirmType"`           // 确认类型
	OnlyOnIndex           bool        `json:"onlyOnIndex"`           // 只在列表页展示
	OnlyOnForm            bool        `json:"onlyOnForm"`            // 只在表单页展示
	OnlyOnDetail          bool        `json:"onlyOnDetail"`          // 只在详情页展示
	ShowOnIndex           bool        `json:"showOnIndex"`           // 在列表页展示
	ShowOnIndexTableRow   bool        `json:"showOnIndexTableRow"`   // 在列表页行展示
	ShowOnIndexTableAlert bool        `json:"showOnIndexTableAlert"` // 在列表页弹出层展示
	ShowOnForm            bool        `json:"showOnForm"`            // 在表单页展示
	ShowOnFormExtra       bool        `json:"showOnFormExtra"`       // 在表单页扩展栏展示
	ShowOnDetail          bool        `json:"showOnDetail"`          // 在详情页展示
	ShowOnDetailExtra     bool        `json:"showOnDetailExtra"`     // 在详情页扩展栏展示
}

// 初始化
func (p *Action) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Action) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *Action) Handle(ctx *builder.Context, query *gorm.DB) error {

	return ctx.JSON(200, message.Error("Method not implemented"))
}

// 行为key
func (p *Action) GetUriKey(action interface{}) string {
	uriKey := reflect.TypeOf(action).String()
	uriKeys := strings.Split(uriKey, ".")
	uriKey = stringy.New(uriKeys[1]).KebabCase("?", "").ToLower()

	return uriKey
}

// 获取名称
func (p *Action) GetName() string {
	return p.Name
}

// 执行成功后刷新的组件
func (p *Action) GetReload() string {
	return p.Reload
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) GetApiParams() []string {
	return p.ApiParams
}

// 执行行为的接口
func (p *Action) GetApi() string {
	return p.Api
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) GetActionType() string {
	return p.ActionType
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) GetSubmitForm() string {
	return p.SubmitForm
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) GetType() string {
	return p.Type
}

// 设置按钮大小,large | middle | small | default
func (p *Action) GetSize() string {
	return p.Size
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) GetWithLoading() bool {
	return p.WithLoading
}

// 设置按钮的图标组件
func (p *Action) GetIcon() string {
	return p.Icon
}

// 行为表单字段
func (p *Action) GetFields() interface{} {
	return p.Fields
}

// 确认标题
func (p *Action) GetConfirmTitle() string {
	return p.ConfirmTitle
}

// 确认文字
func (p *Action) GetConfirmText() string {
	return p.ConfirmText
}

// 确认类型
func (p *Action) GetConfirmType() string {
	return p.ConfirmType
}

// 设置名称
func (p *Action) SetName(name string) {
	p.Name = name
}

// 设置执行成功后刷新的组件
func (p *Action) SetReload(componentKey string) {
	p.Reload = componentKey
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) SetApiParams(apiParams []string) {
	p.ApiParams = apiParams
}

// 执行行为的接口
func (p *Action) SetApi(api string) {
	p.Api = api
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) SetActionType(actionType string) {
	p.ActionType = actionType
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) SetSubmitForm(submitForm string) {
	p.SubmitForm = submitForm
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) SetType(buttonType string) {
	p.Type = buttonType
}

// 设置按钮大小,large | middle | small | default
func (p *Action) SetSize(size string) {
	p.Size = size
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) SetWithLoading(loading bool) {
	p.WithLoading = loading
}

// 设置按钮的图标组件
func (p *Action) SetIcon(icon string) {
	p.Icon = icon
}

// 行为表单字段
func (p *Action) SetFields(fields interface{}) {
	p.Fields = fields
}

// 确认标题
func (p *Action) SetConfirmTitle(confirmTitle string) {
	p.ConfirmTitle = confirmTitle
}

// 确认文字
func (p *Action) SetConfirmText(confirmText string) {
	p.ConfirmText = confirmText
}

// 确认类型
func (p *Action) SetConfirmType(confirmType string) {
	p.ConfirmType = confirmType
}

// 设置行为前的确认操作
func (p *Action) WithConfirm(title string, text string, confirmType string) {

	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType
}

// 只在列表页展示
func (p *Action) SetOnlyOnIndex(value bool) {
	p.OnlyOnIndex = value
	p.ShowOnIndex = value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了列表页外展示
func (p *Action) SetExceptOnIndex() {
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
	p.ShowOnIndex = false
}

// 只在表单页展示
func (p *Action) SetOnlyOnForm(value bool) {
	p.ShowOnForm = value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表单页外展示
func (p *Action) SetExceptOnForm() {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = false
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 只在表单页右上角自定义区域展示
func (p *Action) SetOnlyOnFormExtra(value bool) {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表单页右上角自定义区域外展示
func (p *Action) SetExceptOnFormExtra() {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = false
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 只在详情页展示
func (p *Action) SetOnlyOnDetail(value bool) {
	p.OnlyOnDetail = value
	p.ShowOnDetail = value
	p.ShowOnIndex = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetailExtra = !value
}

// 除了详情页外展示
func (p *Action) SetExceptOnDetail() {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetailExtra = true
}

// 只在详情页右上角自定义区域展示
func (p *Action) SetOnlyOnDetailExtra(value bool) {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = value
}

// 除了详情页右上角自定义区域外展示
func (p *Action) SetExceptOnDetailExtra() {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = false
}

// 在表格行内展示
func (p *Action) SetOnlyOnIndexTableRow(value bool) {
	p.ShowOnIndexTableRow = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表格行内外展示
func (p *Action) SetExceptOnIndexTableRow() {
	p.ShowOnIndexTableRow = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 在表格多选弹出层展示
func (p *Action) SetOnlyOnIndexTableAlert(value bool) {
	p.ShowOnIndexTableAlert = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表格多选弹出层外展示
func (p *Action) SetExceptOnIndexTableAlert() {
	p.ShowOnIndexTableAlert = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 在列表页展示
func (p *Action) SetShowOnIndex() {
	p.ShowOnIndex = true
}

// 在表单页展示
func (p *Action) SetShowOnForm() {
	p.ShowOnForm = true
}

// 在表单页右上角自定义区域展示
func (p *Action) SetShowOnFormExtra() {
	p.ShowOnFormExtra = true
}

// 在详情页展示
func (p *Action) SetShowOnDetail() {
	p.ShowOnDetail = true
}

// 在详情页右上角自定义区域展示
func (p *Action) SetShowOnDetailExtra() {
	p.ShowOnDetailExtra = true
}

// 在表格行内展示
func (p *Action) SetShowOnIndexTableRow() {
	p.ShowOnIndexTableRow = true
}

// 在多选弹出层展示
func (p *Action) SetShowOnIndexTableAlert() {
	p.ShowOnIndexTableAlert = true
}

// 判断是否在列表页展示
func (p *Action) ShownOnIndex() bool {
	if p.OnlyOnIndex {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnIndex
}

// 判断是否在表单页展示
func (p *Action) ShownOnForm() bool {
	if p.OnlyOnForm {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnIndex {
		return false
	}

	return p.ShowOnForm
}

// 判断是否在详情页展示
func (p *Action) ShownOnDetail() bool {
	if p.OnlyOnDetail {
		return true
	}

	if p.OnlyOnIndex {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnDetail
}

// 判断是否在表格行内展示
func (p *Action) ShownOnIndexTableRow() bool {
	return p.ShowOnIndexTableRow
}

// 判断是否在多选弹出层展示
func (p *Action) ShownOnIndexTableAlert() bool {
	return p.ShowOnIndexTableAlert
}

// 判断是否在表单页右上角自定义区域展示
func (p *Action) ShownOnFormExtra() bool {
	return p.ShowOnFormExtra
}

// 判断是否在详情页右上角自定义区域展示
func (p *Action) ShownOnDetailExtra() bool {
	return p.ShowOnDetailExtra
}
